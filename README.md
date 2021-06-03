[![Build Status](https://dev.azure.com/huddeldaddel/Personal%20Projects/_apis/build/status/huddeldaddel.retro-carnage?branchName=main)](https://dev.azure.com/huddeldaddel/Personal%20Projects/_build/latest?definitionId=12&branchName=main)
[![Code Inspector](https://www.code-inspector.com/project/15536/score/svg)](https://frontend.code-inspector.com/public/project/15536/retro-carnage/dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/huddeldaddel/retro-carnage)](https://goreportcard.com/report/github.com/huddeldaddel/retro-carnage)

# RETRO CARNAGE

The goal of this project is to take you back to your childhood by building a modern multi-directional scrolling shooter.
Once finished, Retro-Carnage is going to be a worthy successor of classic video games like
[Ikari Warriors](https://en.wikipedia.org/wiki/Ikari_Warriors) by [SNK](http://www.snk-corp.co.jp/),
[Warzone](https://core-design.com/warzone.html) by [Core Design](https://core-design.com/), or
[Dogs of War](https://en.wikipedia.org/wiki/Dogs_of_War_(1989_video_game))
by [Elite Systems](http://www.elite-systems.co.uk).

This game is currently being developed - but not ready to get played, yet.

[![Watch the video](docs/images/youtube-2021-06-03.png)](https://youtu.be/7Id8BZXu6S0)
Development status 2020-08-25

## Build & Run

Retro-Carnage is being developed on Ubuntu Linux (latest). Follow these steps to get the code up & running (manually).

- Make sure you have go (>= 1.16) and git installed
- Then install the required development libraries: `sudo apt-get install -y libgl1-mesa-dev xorg-dev libasound2-dev`
- Get the code: `git clone https://github.com/huddeldaddel/retro-carnage.git`
- Change into the src directory: `cd retro-carnage/src`
- Install required modules: `go get -d`
- Build the application: `go build`
- Move the binary one level up: `mv retro-carnage ./..`
- Change into the main directory: `cd ..`
- Finally: start the game! `./retro-carnage`

The repository contains IDE configurations for JetBrains Goland to test and run the game.