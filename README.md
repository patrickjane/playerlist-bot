# Overview

This bot connects to one or more gameservers which use valves RCON protocol, to retrieve the list of active players on the server, and posts that list to a discord channel. In addition, the bot can add notifications to a channel whenever a player newly connects to a server, or leaves the server.

![Screenshot](./screenshots/screenshot1.png)

# Settings

The bot is configured using environment variables. The following variables exist:

| Variable                     | Required | Default  | Note |
|------------------------------|----------|----------|----------|
|DISCORD_BOT_TOKEN             | Yes      |          | The discord bot token |
|DISCORD_CHANNEL_ID_STATUS     | Yes      |          | The discord channel ID for posting the overall player list status |
|DISCORD_CHANNEL_ID_JOINLEAVE  | No       | DISCORD_CHANNEL_ID_STATUS          | The discord channel ID for posting the join/leave notifications |
|DISCORD_MESSAGE_TAG           | No       | playerlistbot | Small identifier to retrieve the message from the channel for editing (otherwise, every update would get a new message) |
|DISCORD_CACHE_PATH            | No       | ./discordplayerlistbot.txt | Cache file to store the message ID of the status message the bot is going to edit upon every player list update. Used to prevent complicated lookup of the message after a bot restart. |
|DISCORD_SHOW_JOINLEAVE        | No       |  true     | Send join/leave notifications to discord every time the bot detecs a player joining or leaving a server |
|DISCORD_PIN_PLAYERLIST        | No       |  true     | The discord bot token |
|RCON_QUERY_EVERY_S            | No       |  30     | The interval at which the bot queries the playerlist at the game server via RCON |
|RCON_SERVERS                  | YES       |          | IP,name,password tuples (separated by ;) for the game servers to connect to |

# Installation

