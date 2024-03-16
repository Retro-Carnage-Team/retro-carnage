# Small helper script that prepares the local environment for execution of unit tests.
# It sets the required environment variables, rebuilds the app and executes tests.

$assetFolder = Resolve-Path -Path "..\retro-carnage-assets"
Set-Item -Path Env:\RC-ASSETS -Value $assetFolder
Set-Item -Path Env:\sound -Value "no-fx;no-music"
go build -v
go test -v ./...
