# CallumCTF2025

In this CTF, there will be three challenges. Overcome all challenges and win the prize

## Setup

The following are pre-requisites to playing this CTF

 - Docker
 - GNU Tools (gcc)
 - Python

The CTF is restricted to windows but if demand to make it cross-platform, it can be done to make the CTF cross-platform

You can call the following commands to start the CTF

```
git clone https://github.com/loganmacdougall/CallumCTF2025
cd CallumCTF2025
docker compose up --build -d
```

## Note

There are areas in the CTF where self DDos might be possible. Much effort went into preventing it when possible but in case it's not, I will provided the commands to restart the docker container here

```
docker compose up c_server --build -d # Challenge 1
docker compose up puzzle_castle --build -d # Challenge 1
```

## Have fun
Happy Hacking