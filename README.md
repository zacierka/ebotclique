# EBotClique Discord Bot
This project is a bot for "my" personal discord. It ties in various aspects of the discord with other assets. This bot has many modules which may communicate with our minecraft server, retrieve analytics for various games, general fun commands, and more. This bot is a work in progress and will be maintained as best as possible.

This Project can be forked and edited. Just know that setup process may not be streamlined in the beginning stages as many things will change.

## Frameworks/Technologies
EBotClique Bot is centered around a few technologies which enable it to communicate with other assets.
* [MariaDB](https://mariadb.org/)                        - Database storage implementation
* [Redis](https://redis.io/)                             - Messaging and Shared Memory implementation
* [DiscordGo](https://github.com/bwmarrin/discordgo)     - Go Wrapper for Discord Bot
* [EBC-Plugin](https://github.com/zacierka/EBC-Plugin)* - Java Minecraft Plugin Extension built by me
* [demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang) - CSGO Analytics
* [Leetify](https://beta.leetify.com/app)                - CSGO Analytics

## Features
\* := Requires Authorized Role
|  Module   |  Command  |     Description    |            Usage             | ✅❌ |
| --------- | --------- | ------------------ | ---------------------------- | ------ |
|  common   | quote     | Quote from a user  | .user (add) \[quote\]        | ✅    |
|  common   | info      | Channel and Uptime | .info                        | ✅    |
|  common   | ping      | Reply Ping         | .ping                        | ✅    |
| minecraft | online    | Users Online list  | .online                      | ❌    |
| minecraft | mcstatus  | Server Status      | .mcstatus                    | ❌    |
| minecraft | mcreboot* | Reboot Server      | .mcreboot                    | ❌    |
| minecraft | mcstop*   | Stop Server        | .mcstop                      | ❌    |
| minecraft | mclag     | Clear Lag          | .mclag                       | ❌    |
| minecraft | mcwho     | Get Users MC Name  | .mcwho dname:mcname          | ❌    |
| minecraft | mcadd*    | Whitelist User     | .mcadd @dname mcname         | ❌    |
| minecraft | mcrm*     | Blacklist User     | .mcrm @dname:mcname          | ❌    |
| analytics | csgo*     | See CSGO Stats     | .csgo join:leave:reportcard  | ❌    |
| analytics | leetify   | Get Leetify Link   | .leetify team:user           | ❌    |
|  control  | enable*   | Enable a Module    | .enable module               | ❌    |
|  control  | disable*  | Disable a Module   | .disable module              | ❌    |

### Setup Process
The setup process for this bot is a work in progress. It goes something like...
1. Clone, Modify, build ```go build``` this repo.
2. Clone, Modify, build ```mvn package``` [EBC-Plugin](https://github.com/zacierka/EBC-Plugin) repo.
3. Install Required Technologies and configure.
4. Enjoy

#### Extended Goals
My current implementation of this is hosted locally on my home network. The desired goal is to host on some cloud provider such as AWS or Google Cloud. Another goal is to have a companion website. This website may interface with the discord and offer more visuals using the csgo module such as better statistics and graphs. There is a current website, but I would rather have a working minimal viable product till I start developing that.

#### Contribution 
Closed. This is a personal project to express backend as well as frontend skills.

Thanks\
\- switch
