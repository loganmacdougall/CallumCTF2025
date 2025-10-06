# CallumCTF2025

In this CTF, there will be three challenges. Overcome all challenges and win the prize

## Setup

This CTF requires that docker is working on the computer.

The CTF works with Windows and Linux. If there's any issue running the main executable (without revealing too much about it), you can open the game folder of this repo, open the project, and press the play button on the opened project. However, it is recommend to run the executable if possible.

You can call the following commands to start the CTF

```
git clone https://github.com/loganmacdougall/CallumCTF2025
cd CallumCTF2025
docker compose up --build -d
```

Once the servers are running, you can enter your browser and go to port `http://localhost:8080`

## Note

There are areas in the CTF where self DDos might be possible. Much effort went into preventing it when possible but in case it's not, I will provided the commands to restart the docker container here

```
docker compose up challenge1 --build -d # Challenge 1
docker compose up challenge2 --build -d # Challenge 2
docker compose up challenge3 --build -d # Challenge 3
```

It's also possibly the case that there are several minor bugs throughout this challenge. To that I say I do not care unless it's really taking away from the enjoyment of the CTF. This is meant to be a birthday gift and not production code.

## Have fun
Happy Hacking