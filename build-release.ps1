# This script creates a platform specific build of a specific release of Retro-Carnage.
#
# This script can be executed on a target platform of Retro-Carnage to create a build. Cross compiling is a little 
# tricky, as the game depends on OpenGL and getting x-platform builds right is not that easy. So currently my attempt
# is to create builds on the target platforms. This script is designed to do just that.
#
# Usage:
# Run script and specify a release tag (https://github.com/Retro-Carnage-Team/retro-carnage/tags) as parameter. 
# Example:
# 
# PS> .\build-release.ps1 v2024-03
#
# This script has dependencies that have to be installed first:
#
# All platforms:
# * git
# * go
#
# Ubuntu:
# * libgl1-mesa-dev 
# * xorg-dev
# * libasound2-dev
#
# Windows:
# * tdm-gcc (https://jmeubank.github.io/tdm-gcc/)

param(
    [Parameter(Mandatory = $true, ValueFromPipelineByPropertyName = $true, Position = 0)]
    [string]$releaseTag
)

class ApplicationBuilder {

    [string] $location
    [string] $outFile
    [string] $releaseTag
    [string] $workFolder    

    ApplicationBuilder($tag) {
        $this.location = Get-Location
        $this.releaseTag = $tag
        $this.outFile = $this.GetOutputFileName()        
        $this.workFolder = $this.CreateTempFolder()
    }

    [void] BuildRelease() {
        $this.DownloadAssets()
        $this.BuildBinary()
        $this.BuildArchive()
        $this.CleanUp()
    }

    [void] DownloadAssets() {
        Write-Host "Loading assets"
        Set-Location $this.workFolder
        git clone --depth 1 --branch $this.releaseTag https://github.com/Retro-Carnage-Team/retro-carnage-assets.git $this.releaseTag
        Remove-Item -Recurse -Force (Join-Path $this.workFolder (Join-Path $this.releaseTag ".git"))
    }

    [void] BuildBinary() {
        Write-Host "Building Retro-Carnage binary"        
        git clone --depth 1 --branch $this.releaseTag https://github.com/Retro-Carnage-Team/retro-carnage.git
        Set-Location (Join-Path $this.workFolder "retro-carnage")
        go build -v
        Move-Item -Path ./retro-carnage* -Destination (Join-Path $this.workFolder $this.releaseTag)
        Set-Location $this.location
        Remove-Item -Recurse -Force (Join-Path $this.workFolder "retro-carnage")
    }

    [void] BuildArchive() {                
        $compress = @{
            Path = Join-Path $this.workFolder $this.releaseTag
            CompressionLevel = "Optimal"
            DestinationPath = $this.outFile
        }
        Compress-Archive @compress        
    }

    [void] CleanUp() {
        Write-Host "Cleaning up"
        Remove-Item -Recurse -Force $this.workFolder
    }

    [string] CreateTempFolder() {
        $guid = [System.Guid]::NewGuid()
        $tmpFolder = [System.IO.Path]::GetTempPath()
        $folder = Join-Path $tmpFolder $guid
        New-Item -ItemType Directory -Path $folder
        return $folder
    }

    [string] GetOutputFileName() {                
        if([System.Environment]::OSVersion.Platform -eq "Win32NT") {
            return Join-Path $this.location "Retro-Carnage-$($this.releaseTag)-Windows.zip"
        }
        return Join-Path $this.location "Retro-Carnage-$($this.releaseTag)-Linux.zip"
    }
}

$appBuilder = [ApplicationBuilder]::new($releaseTag)
$appBuilder.BuildRelease()
Write-Host "Release created: $($appBuilder.outFile)"
