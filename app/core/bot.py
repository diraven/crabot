"""Core bot module."""
import re
from typing import List

import discord
import sentry_sdk
from discord.ext import commands

from core.cog import Cog
from core.models import Guild
from settings import settings
from .context import Context
from .message import Message

if settings.SENTRY_DSN:
    sentry_sdk.init(settings.SENTRY_DSN)


async def get_prefix(bot: commands.Bot, message: discord.Message) -> List[str]:
    """Return server prefix based on message."""
    return commands.when_mentioned_or(
        settings.DISCORD_DEFAULT_PREFIX,
    )(bot, message)


class Bot(commands.Bot):
    """Bot class."""

    def __init__(self, command_prefix=None, **options) -> None:
        """Make bot instance."""
        if not command_prefix:
            command_prefix = get_prefix

        super().__init__(command_prefix=command_prefix, **options)
        self.add_cog(Cog(self))  # load core cog

    async def get_context(
            self,
            msg: discord.Message,
            *,
            cls=Context,
    ) -> Context:
        """Return command invocation context."""
        ctx: Context = await super().get_context(msg, cls=Context)

        # If command not found - try to find it using alias.
        if ctx.command is None and msg.guild:
            guild, _ = await Guild.get_or_create(ctx.guild.id)
            for alias in guild.get('aliases', []):
                if alias['src'] == ctx.invoked_with:
                    msg.content = re.sub(
                        '^{}{}'.format(ctx.prefix, alias['src']),
                        '{}{}'.format(ctx.prefix, alias['dst']),
                        msg.content,
                    )
                ctx = await super().get_context(
                    msg,
                    cls=Context,
                )

        return ctx

    # noinspection PyBroadException
    async def on_command_error(self, ctx: Context, exception) -> None:
        """Global command errors handler."""
        if isinstance(exception, commands.errors.CommandInvokeError) and \
                isinstance(exception.original, discord.errors.Forbidden):
            await ctx.post(
                Message.danger(
                    f' {exception.original.response.url}.',
                    title=f'Oops... Unable to {exception.original.response.method}: {exception.original.response.reason}',
                ),
            )
            return

        # Ignore other checks failures.
        if isinstance(
                exception, commands.CheckFailure,
        ) or isinstance(
            exception, commands.CommandNotFound,
        ):
            return

        # Process missing command arguments error by responding to the user.
        if isinstance(exception, commands.MissingRequiredArgument):
            await ctx.post(
                Message.danger(
                    'Incorrect command usage, missing argument. '
                    'Did you forget to add something?',
                ),
            )
            return

        if settings.SENTRY_DSN:
            # Send issue to sentry if configured.
            sentry_sdk.capture_exception(exception)
        else:
            # Raise error otherwise.
            return await super().on_command_error(ctx, exception)
