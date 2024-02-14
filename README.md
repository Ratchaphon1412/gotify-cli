# Command line Gotify client with Socket tcp

send message to Gotify server using tcp socket

# Workflow

![](https://cdn.discordapp.com/attachments/1037613428132560896/1207357998498975834/Screenshot_2567-02-14_at_18.58.20.png?ex=65df5ac2&is=65cce5c2&hm=706f87d964a54cfabeeb4d299d6e1c1b87123a3ce96b6aaf9511b65908408666&)

Command line Gotify client will send message to Gotify server using tcp socket. The server will then send the message to the Gotify server using the REST API. and then the Gotify server will send the message to the mobile app or web server. using WebSocket.

## Download Mobile App or Use on Web Server

[Ratchaphon Gotify](https://noti.ratchaphon1412.co/) click this link to go to the website. Gotify Demo

> Domain : https://noti.ratchaphon1412.co/

> User : guest

> Password : pass

or download the mobile app from the link below Supported on Android Only

[Gotify Mobile App](https://gotify.net/)

## Usage Command Line in Linux

clone the repo

```bash
$ git clone https://github.com/Ratchaphon1412/gotify-cli.git
$ cd gotify-cli
```

```bash
$ ./gotify-cli -h
```

for help and usage

```bash
$ ./gotify-cli sendNotifyMessage "Message" host:port
```

## Test with Server

using the ip 119.59.99.143:8998

```bash

./gotify-cli sendNotifyMessage "Hello World" 119.59.99.143:8998

```

## Server Repo

[Click this link](https://github.com/Ratchaphon1412/gotify-cli-server)

## Demo

[Click this link](https://youtu.be/qDsILvQNk2s)
