import asyncio
import re
from typing import List, Optional, Tuple

import editdistance
from discord import Role, Forbidden, Reaction, User
from discord.ext import commands
from django.utils.text import slugify as django_slugify

from mydiscord.cogbase import CogBase
from mydiscord.context import Context
from mydiscord.message import Message
from mydiscord.models import Guild
from publicroles.models import PublicRole


def slugify(text: str) -> str:
    text = re.sub(r'[\s]', '', text).strip().lower()
    return django_slugify(text)


def sync_roles(ctx: Context) -> None:
    """
    Synchronizes locally stored list of roles with the real server roles.
    """
    # Get a list of stored public roles discord_ids.
    public_roles_ids = PublicRole.objects.filter(
        guild__discord_id=ctx.guild.id,
    ).values_list('discord_id', flat=True)

    # Get a list of real roles of the server.
    real_roles = {}
    for discord_role in ctx.guild.roles:
        real_roles[discord_role.id] = discord_role

    # Delete all roles we have stored locally that do not exist on the server
    # any more.
    for discord_id in public_roles_ids:
        if discord_id not in real_roles.keys():
            PublicRole.objects.get(discord_id=discord_id).delete()


def format_list(items: List[Role], offset: int = 0) -> Tuple[str, int]:
    """
    Converts discord roles list to the string representation. Limited to 1000
    symbols.
    """
    if items:
        names = []
        length = 0
        # For eash item after an offset.
        for item in items[offset:]:
            # Calculate resulting string length.
            length += len(str(item))
            # If list is too large - stop adding further.
            if length > 1000:
                break
            # Add item to the list.
            names.append(item.name)
            # Adjust offset.
            offset += 1
        # Prepare result.
        result = ", ".join(['`{}`'.format(name) for name in names])
        # Offset already contains count of names, so we should not add it here.
        if offset < len(items):
            # Post a message notifying of leftover items.
            result += "... Click button below to see {} more.".format(
                len(items) - offset
            )
        return result, offset
    else:
        return "Nothing found.", offset


async def post_list(
        ctx: Context,
        items: List,
        offset: int = 0,
        title: str = None
):
    result, offset = format_list(items, offset)
    message = await ctx.post(Message(result, title))

    if offset < len(items):
        # Emoji to be used as button.
        emoji = '▶'

        # Add emoji.
        await ctx.react(emoji, message)

        def check(reaction: Reaction, user: User):
            return user == ctx.author and str(reaction.emoji) == emoji

        try:
            await ctx.bot.wait_for(
                'reaction_add',
                timeout=20.0,
                check=check,
            )
            await post_list(ctx, items, offset, title)
        except asyncio.TimeoutError:
            await message.remove_reaction(emoji, ctx.me)


async def fuzzy_search_public_roles(
        ctx: Context, roles: List[Role], arg: str
) -> List[Role]:
    public_roles_ids = PublicRole.objects.filter(
        guild__discord_id=ctx.guild.id,
    ).values_list('discord_id', flat=True)

    # Now search guild roles list for the term provided.
    result = []

    # Strict search.
    for role in roles:
        if role.id in public_roles_ids and arg == role.name:
            result.append(role)
    if result:
        return result

    # Partial match search.
    for role in roles:
        if role.id in public_roles_ids and slugify(arg) in slugify(role.name):
            result.append(role)
    if result:
        return result

    # Fuzzy search.
    for role in roles:
        if role.id in public_roles_ids and \
                editdistance.eval(
                    slugify(role.name),
                    slugify(arg),
                ) < 2:
            result.append(role)
    return result


async def fuzzy_search_public_role(
        ctx: Context, roles: List[Role], arg: str
) -> Optional[Role]:
    result = await fuzzy_search_public_roles(ctx, roles, arg)

    # Set symbols limit to optimize search results.
    if 0 < len(arg) < 3:
        await ctx.post(
            Message.danger('Search must contain at least 3 symbols.')
        )
        return None

    # If no roles found.
    if len(result) == 0:
        await ctx.post(Message.danger('Role not found.'))

    # If one role found.
    if len(result) == 1:
        return result[0]

    # If multiple roles found.
    if len(result) > 1:
        await post_list(ctx, result, title='Multiple roles found')

    return None


class Cog(CogBase):
    @commands.group(invoke_without_command=True)
    async def publicroles(
            self,
            ctx: Context,
            *args
    ) -> None:
        """
        Searches available public roles.
        """
        sync_roles(ctx)
        arg = ' '.join(args)

        # Set symbols limit to optimize search results.
        if 0 < len(arg) < 3:
            await ctx.post(
                Message.danger('Search must contain at least 3 symbols.')
            )
            return None

        roles = await fuzzy_search_public_roles(ctx, ctx.guild.roles, arg)

        if len(roles) > 0:
            await post_list(ctx, roles, title='Public Roles')
        else:
            await ctx.post(Message.danger("No roles found."))

    @publicroles.command()
    async def my(
            self,
            ctx: Context
    ) -> None:
        """
        Shows public roles you have.
        """
        sync_roles(ctx)

        # Get a list of stored public roles discord_ids.
        public_roles_ids = PublicRole.objects.filter(
            guild__discord_id=ctx.guild.id,
        ).values_list('discord_id', flat=True)

        roles = []
        for role in ctx.author.roles:
            if role.id in public_roles_ids:
                roles.append(role)

        await post_list(ctx, roles, title='Your Public Roles')

    @publicroles.command()
    async def join(
            self,
            ctx: Context,
            *args
    ) -> None:
        """
        Gives you a public role.
        """
        sync_roles(ctx)
        arg = ' '.join(args)

        # Search roles.
        role = await fuzzy_search_public_role(ctx, ctx.guild.roles, arg)

        if role:
            try:
                await ctx.author.add_roles(role)
                await ctx.success()
            except Forbidden as e:
                await ctx.post(Message.danger(str(e)))
            return

    @publicroles.command()
    async def leave(
            self,
            ctx: Context,
            *args
    ) -> None:
        """
        Removes you from a public role.
        """
        sync_roles(ctx)
        arg = ' '.join(args)

        # Search roles.
        role = await fuzzy_search_public_role(ctx, ctx.author.roles, arg)

        if role:
            try:
                await ctx.author.remove_roles(role)
                await ctx.success()
            except Forbidden as e:
                await ctx.post(Message.danger(str(e)))
            return

    @publicroles.command()
    async def who(
            self,
            ctx: Context,
            *args
    ) -> None:
        """
        Shows a list of people who has the public role.
        """
        sync_roles(ctx)
        arg = ' '.join(args)

        # Search roles.
        role = await fuzzy_search_public_role(ctx, ctx.guild.roles, arg)

        if role:
            try:  # Try to get the role from our public roles list.
                PublicRole.objects.get(discord_id=role.id)
                members = []
                for member in ctx.guild.members:
                    for member_role in member.roles:
                        if member_role.id == role.id:
                            members.append(member)

                await post_list(ctx, members,
                                title='People with "{}" public role'.format(
                                    role.name)
                                )
                return

            except PublicRole.DoesNotExist:
                await ctx.post(Message.danger(
                    "The '{}' role is not public.".format(arg)
                ))
                return

    @publicroles.command()
    @commands.check(lambda ctx: ctx.author.guild_permissions.manage_roles)
    async def register(
            self,
            ctx: Context,
            *args
    ) -> None:
        """
        Registers existing role as a public role.
        """
        sync_roles(ctx)
        arg = ' '.join(args)

        # Try to get current guild.
        guild = Guild.objects.get(discord_id=ctx.guild.id)

        # Argument must be an exact name of the role to be made public.
        for role in ctx.guild.roles:
            if role.name == arg:
                PublicRole.objects.get_or_create(
                    discord_id=role.id,
                    guild=guild,
                )
                await ctx.success()
                return

        await ctx.post(
            Message.danger('There is no `{}` role.'.format(arg))
        )

    @publicroles.command()
    @commands.check(lambda ctx: ctx.author.guild_permissions.manage_roles)
    async def unregister(
            self,
            ctx: Context,
            *args
    ) -> None:
        """
        Unregisters role as a public role.
        """
        sync_roles(ctx)
        arg = ' '.join(args)

        # Argument must be an exact name of the role that's public.
        # Find out the role discord_id first.
        for role in ctx.guild.roles:
            if role.name == arg:
                try:  # Try to delete the role.
                    PublicRole.objects.get(discord_id=role.id).delete()
                    await ctx.success()
                    return
                except PublicRole.DoesNotExist:
                    await ctx.post(
                        Message.danger(
                            'Role `{}` is not public.'.format(arg)
                        )
                    )
                    return

        await ctx.post(
            Message.danger('There is no `{}` role.'.format(arg))
        )

    @publicroles.command()
    @commands.check(lambda ctx: ctx.author.guild_permissions.manage_roles)
    async def create(
            self,
            ctx: Context,
            *args
    ) -> None:
        """
        Creates new role and makes it public.
        """
        sync_roles(ctx)
        arg = ' '.join(args)

        # Get current guild.
        guild = Guild.objects.get(discord_id=ctx.guild.id)

        # There should not be a server role with such name.
        for role in ctx.guild.roles:
            if slugify(role.name) == slugify(arg):
                await ctx.post(Message.danger(
                    'Role `{}` already exists. Try `register` '
                    'instead of `create`.'.format(arg)
                ))
                return

        # Now it's safe to create a new role.
        try:
            r = await ctx.guild.create_role(
                name=arg,
                mentionable=True,
            )
        except Forbidden as e:
            await ctx.post(Message.danger(str(e)))
            return

        # Save the role locally as public.
        PublicRole(guild=guild, discord_id=r.id).save()

        await ctx.success()