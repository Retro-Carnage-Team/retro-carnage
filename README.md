[![Go](https://github.com/Retro-Carnage-Team/retro-carnage/actions/workflows/build.yml/badge.svg)](https://github.com/Retro-Carnage-Team/retro-carnage/actions/workflows/build.yml) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage)

# RETRO CARNAGE

The goal of this project is to take you back to the best part of your childhood. To do this, we are building a modern
multidirectional scrolling shooter. Once finished, Retro-Carnage is going to be a worthy successor of classic video
games like [Ikari Warriors](https://en.wikipedia.org/wiki/Ikari_Warriors) by [SNK](http://www.snk-corp.co.jp/),
[War Zone](https://core-design.com/warzone.html) by [Core Design](https://core-design.com/), or
[Dogs of War](https://en.wikipedia.org/wiki/Dogs_of_War_(1989_video_game))
by [Elite Systems](http://www.elite-systems.co.uk).

This game is currently being developed - but not ready to get played, yet.

[![Watch the video](youtube-2021-06-03.png)](https://youtu.be/PqWghPZvIy4)
Development status as of 2021-08-09

## Build & Run

Make sure you have Golang (>= 1.20) and git installed

### Install dependencies

#### On Ubuntu

Install the required libraries: `sudo apt-get install -y libgl1-mesa-dev xorg-dev libasound2-dev`

#### On Fedora

Install the required
libraries: `sudo dnf install libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel mesa-libGL-devel xorg-x11-server-devel alsa-lib-devel libXxf86vm-devel`

#### On Windows

Install [tdm-gcc](https://jmeubank.github.io/tdm-gcc/) (so that various go-bindings can be compiled).

### Build and run the game

- Get the code: `git clone https://github.com/Retro-Carnage-Team/retro-carnage.git`
- Get the assets: `git clone https://github.com/Retro-Carnage-Team/retro-carnage-assets`
- Change into the src directory: `cd retro-carnage`
- Install required modules: `go get -d`
- Build the application: `go build`
- Start the game: `./retro-carnage ../retro-carnage-assets`

The repository contains IDE settings for Visual Studio Code to debug, run, and test the game.
