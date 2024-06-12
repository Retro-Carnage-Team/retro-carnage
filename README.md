[![Go](https://github.com/Retro-Carnage-Team/retro-carnage/actions/workflows/build.yml/badge.svg)](https://github.com/Retro-Carnage-Team/retro-carnage/actions/workflows/build.yml) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage) [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=Retro-Carnage-Team_retro-carnage&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=Retro-Carnage-Team_retro-carnage)

# RETRO CARNAGE

The goal of this project is to take you back to the best part of your childhood. To do this, we are building a modern
multidirectional scrolling shooter. Once finished, Retro-Carnage is going to be a worthy successor of classic video
games like [Ikari Warriors](https://en.wikipedia.org/wiki/Ikari_Warriors) by [SNK](http://www.snk-corp.co.jp/),
[War Zone](https://core-design.com/warzone.html) by [Core Design](https://core-design.com/), or
[Dogs of War](https://en.wikipedia.org/wiki/Dogs_of_War_(1989_video_game))
by [Elite Systems](http://www.elite-systems.co.uk).

This game is currently in development and not ready to get played, yet.

[![Watch the video](youtube-2021-06-03.png)](https://youtu.be/PqWghPZvIy4)
Development status as of 2021-08-09

## Getting the latest release

You can find the latest builds for various platforms on the [release page](https://github.com/Retro-Carnage-Team/retro-carnage/releases).

- Latest version for [Linux x86-64](https://rcpublic.blob.core.windows.net/rcrelease/v2024.05/Retro-Carnage-v2024.05-Linux.zip)
- Latest version for [Windows x86-64](https://rcpublic.blob.core.windows.net/rcrelease/v2024.05/Retro-Carnage-v2024.05-Windows.zip)

### Run the game

- Download the application archive for your platform
- Unzip the application archive
- On Linux only: make application executeable `chmod +x retro-carnage`
- **Double click on application** or use terminal to run `.\retro-carnage` in application folder


## Getting Started

### Prerequisites

#### All platforms

- git
- Golang (>= 1.22) 
- PowerShell

#### On Ubuntu

Install the required libraries:

`sudo apt-get install -y libgl1-mesa-dev xorg-dev libasound2-dev`

#### On Fedora

Install the required libraries:

`sudo dnf install libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel mesa-libGL-devel xorg-x11-server-devel alsa-lib-devel libXxf86vm-devel`

#### On Windows

Install [tdm-gcc](https://jmeubank.github.io/tdm-gcc/) so that various go-bindings can be compiled. A installation with default option will do fine.

### Installing

Get the code and assets

`git clone https://github.com/Retro-Carnage-Team/retro-carnage.git`  
`git clone https://github.com/Retro-Carnage-Team/retro-carnage-assets`

Change into the src directory, install required modules, compile the application

`cd retro-carnage`  
`go get -d`  
`go build`

Start the game

`./retro-carnage ../retro-carnage-assets`

The repository contains IDE settings for Visual Studio Code to debug, run, and test the game.

## Running the tests

Run the steps to install the development environment first (see previous chapter).
Open a terminal, navigate into the application folder and run the test script:

`pwsh .\test.ps1`

## Authors

- **[Thomas Werner](https://github.com/huddeldaddel)**

## License

This project is licensed under the MIT License. See the [LICENSE.md](./LICENSE.md) file for details.

## Acknowledgments

This game is based on the work of many great artists who share their work free of charge.
See the [ATTRIBUTION.md](ATTRIBUTION.md) file for details.