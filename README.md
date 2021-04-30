[![Build Status](https://dev.azure.com/huddeldaddel/Personal%20Projects/_apis/build/status/huddeldaddel.retro-carnage?branchName=main)](https://dev.azure.com/huddeldaddel/Personal%20Projects/_build/latest?definitionId=12&branchName=main)
[![Code Inspector](https://www.code-inspector.com/project/15536/score/svg)](https://frontend.code-inspector.com/public/project/15536/retro-carnage/dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/huddeldaddel/retro-carnage)](https://goreportcard.com/report/github.com/huddeldaddel/retro-carnage)

# RETRO CARNAGE

The goal of this project is to build a modern multi-directional scrolling shooter - a worthy successor of the classic
1989 video game [Dogs of War](https://gamesdb.launchbox-app.com/games/details/41090) by
[Elite Systems](http://www.elite-systems.co.uk).

This game is currently under active development. At the moment you can test different concepts of the game but there is
no gripping gameplay.

An unfinished game does not deter you? Then you can find the current state of development on the official homepage of
the game: [http://www.retro-carnage.net](http://www.retro-carnage.net).

[![Watch the video](docs/images/youtube-2020-08-25.png)](https://youtu.be/IeUowwMaIB4)
Development status 2020-08-25

[![Watch the video](docs/images/youtube-first-impression.png)](https://youtu.be/W5dJvoZUGt8)
Development status 2020-04-17

## Build & Run

Retro-Carnage is being developed on Ubuntu Linux (latest). Follow these steps to get the code up & running.  

 - Make sure you have go (>= 1.15) and git installed
 - Install the required development libraries: `sudo apt-get install -y libgl1-mesa-dev xorg-dev libasound2-dev`
 - Get the code: `git clone https://github.com/huddeldaddel/retro-carnage.git`
 - Change into the src directory: `cd retro-carnage.net/src`
 - Install required modules: `go get -d`
 - Build the application: `go build`
 - Run the tests: `go test -v ./...`
 - Move the binary one level up: `mv retro-carnage.net ./..`
 - Change into the main directory: `cd ..`
 - Finally: start the game! `./retro-carnage.net`

## Usage statistics

![Usage statistics](https://backend.retro-carnage.net/usage/chart "Usage statistics")
