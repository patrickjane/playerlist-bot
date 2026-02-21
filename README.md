# Overview

This bot connects to one or more gameservers which use valves RCON protocol, to retrieve the list of active players on the server, and posts that list to a discord channel. In addition, the bot can add notifications to a channel whenever a player newly connects to a server, or leaves the server.

<img title="Screenshot" alt="Screenshot" width="50%" border="1" src="screenshots/screenshot1.png">

## Supported games

This library uses the awesome github.com/gorcon/rcon library to interact with RCON servers. The bot should therefore work with _any_ game which uses the Source RCON protocol.

I have tested it with **Ark: Survival Ascended**, but it should work with other game servers as well.

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
|RCON_QUERY_EVERY_S            | No       |  30     | The interval (in seconds) at which the bot queries the playerlist at the game server via RCON |
|RCON_SERVERS                  | YES       |          | IP,name,password tuples (separated by ;) for the game servers to connect to |
|LOG_FILE                  | No      |          | File to store logmessages in. Logs to stdout if omitted (default). |

#### Note

`RCON_SERVERS` must be given in the following syntax:

```
RCON_SERVERS="10.0.0.1:27015,Server One,secret123;10.0.0.2:27015,Server two,backup456"
```

# Installation

## Binary

Download a binary from [Releases](https://github.com/patrickjane/playerlist-bot/releases) for your corresponding architecture, and run the binary. Make sure to set the environment variables as per the above definitions.

Example:

```
$ ./playerlistbot.linux-amd64
2026/02/19 18:37:30 INFO PlayerlistBot 1.0.0
2026/02/19 18:37:30 INFO https://github.com/patrickjane/playerlist-bot
2026/02/19 18:37:30 INFO Monitoring the following servers via RCON:
2026/02/19 18:37:30 INFO    Lost Colony at xx.xx.xx.xx:27020
2026/02/19 18:37:30 INFO    Extinction at xx.xx.xx.xx:27022
2026/02/19 18:37:30 INFO Query servers every 30 seconds
2026/02/19 18:37:30 INFO Connecting to discord
2026/02/19 18:37:30 INFO Creating RCON reader
2026/02/19 18:37:30 INFO Successfully started.

```

## Docker

The bot comes with 2 docker images, linux/amd64 and linux/arm64. Image path is `ghcr.io/patrickjane/playerlist-bot:latest` (or substitute `latest` with an actual version.).

You can run it directly, or e.g. use a `docker compose` file:

`docker-compose.yml`

```
services:
  playerlistbot:
    container_name: playerlist-bot
    image: ghcr.io/patrickjane/playerlist-bot:latest
    restart: unless-stopped
    environment:
      - DISCORD_CHANNEL_ID_STATUS=xxxx
      - DISCORD_BOT_TOKEN=xxxxx
      - DISCORD_CACHE_PATH=/data/cache.txt
      - RCON_QUERY_EVERY_S=30
      - RCON_SERVERS=xx.xx.xx.xx:27020,Lost Colony,MyPassword;xx.xx.xx.xx:27022,Extinction,MyPassword
    volumes:
      - ./:/data
```

And then:

```
$ docker compose up
[+] up 1/1
 ✔ Container playerlist-bot Recreated                                                                                                                                                                    0.1s
Attaching to playerlist-bot
playerlist-bot  | 2026/02/19 17:30:55 INFO PlayerlistBot main
playerlist-bot  | 2026/02/19 17:30:55 INFO https://github.com/patrickjane/playerlist-bot
playerlist-bot  | 2026/02/19 17:30:55 INFO Monitoring the following servers via RCON:
playerlist-bot  | 2026/02/19 17:30:55 INFO    Lost Colony at "xx.xx.xx.xx:27020
playerlist-bot  | 2026/02/19 17:30:55 INFO    Extinction at xx.xx.xx.xx:27022
playerlist-bot  | 2026/02/19 17:30:55 INFO Query servers every 30 seconds
playerlist-bot  | 2026/02/19 17:30:55 INFO Connecting to discord
playerlist-bot  | 2026/02/19 17:30:55 INFO Creating RCON reader
playerlist-bot  | 2026/02/19 17:30:55 INFO Successfully started.
```

#### Note
Use `docker compose up -d` to run the bot in background. Afterwards use `docker ps` to find the container, and `docker logs -f [ID]` to show the logs of the container.

## License
MIT License, see [LICENSE](LICENSE)