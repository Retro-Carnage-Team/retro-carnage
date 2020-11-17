# Game

Retro-Carnage is currently beeing implemented as a native application using the go programming language.

## Structure

TODO

## Build & Run

### The game

- Make sure to have a at least version 1.15 of go installed.
- Get the latest source code from [GitHub](https://github.com/huddeldaddel/retro-carnage).
- On Ubuntu Linux you'll need to install additional dependencies:
  - `sudo apt-get update`
  - `sudo apt-get install -y libgl1-mesa-dev xorg-dev libasound2-dev`
- Open the project folder using GoLand for best development experience.

TODO: Add commands that run a build on cmdline

### The documentation

The documentation uses MkDocs to transform Markdown documents into static HTML and JavaScript files.

- Make sure to have a recent version of MkDocs installed. Development of the game happens with version 1.0.4 running on
  Python 3.8.
- Get the latest source code from [GitHub](https://github.com/huddeldaddel/retro-carnage).
- Open you command line, navigate to the project folder.
- Run `mkdocs build` to create a production build of the documentation. This will result in a bunch of static files and
  assets in **./site**. Copy this folder to a web server, and you're ready to go.
- If you want to open the documentation locally or want to work on it, run `mkdocs serve` instead. This will build the
  docs and start a local server on [http://localhost:8000/](http://localhost:8000/). It will notice any changes you make
  to the documentation files and update the build accordingly.
