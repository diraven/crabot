"""Custom bot commands context."""
import discord
from discord.ext import commands

from core.emoji import EMOJI_UNICODE

DISCONTINUED = "*WARNING:* this functionality is deprecated. Contact server owner for details."

class Context(commands.Context):
    """Bot command context."""

    async def post_info(self, text: str):
        """Send info message."""
        return await self.send(
            embed=discord.Embed(
                description=text,
                title=":information_source:",
                color=discord.Color.blue(),
                footer=DISCONTINUED,
            )
        )

    async def post_warning(self, text: str):
        """Send warning message."""
        return await self.send(
            embed=discord.Embed(
                description=text,
                title=":warning:",
                color=discord.Color.orange(),
                footer=DISCONTINUED,
            )
        )

    async def post_error(self, text: str):
        """Send error message."""
        return await self.send(
            embed=discord.Embed(
                description=text,
                title=":no_entry:",
                color=discord.Color.red(),
                footer=DISCONTINUED,
            )
        )

    async def react_ok(self) -> discord.Reaction:
        """Add ok hand reaction to the message."""
        return await self.post_info(EMOJI_UNICODE[":OK_hand:"])
