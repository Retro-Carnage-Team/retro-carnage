# This script creates a platform specific build of a specific release of Retro-Carnage.
#
# This script can be executed on a target platform of Retro-Carnage to create a build. Cross compiling is a little 
# tricky, as the game depends on OpenGL and getting x-platform builds right is not that easy. So currently my attempt
# is to create builds on the target platforms. This script is designed to do just that.
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


param(
    [Parameter(Mandatory = $true, ValueFromPipelineByPropertyName = $true, Position = 0)]
    [string]$releaseTag
)

class ApplicationBuilder {

    [string] $outFile
    [string] $releaseTag
    [string] $workFolder    

    ApplicationBuilder($tag) {
        $this.releaseTag = $tag
        $this.outFile = $this.GetOutputFileName()        
        $this.workFolder = $this.CreateTempFolder()
    }

    [void] BuildRelease() {
        $this.DownloadAssets()
        $this.BuildBinary()
        $this.BuildArchive($this.outFile)
    }

    [void] BuildBinary() {
        Write-Host "Building Retro-Carnage binary"
        Set-Location $this.workFolder        
        git clone --depth 1 --branch $this.releaseTag https://github.com/Retro-Carnage-Team/retro-carnage.git
        Set-Location ./retro-carnage
        go build -v
        Move-Item -Path ./retro-carnage* -Destination (Join-Path $this.workFolder $this.releaseTag)
        Set-Location $this.workFolder
        Remove-Item -Recurse -Force (Join-Path $this.workFolder "retro-carnage")
    }

    [void] DownloadAssets() {
        Write-Host "Loading assets"
        Set-Location $this.workFolder
        git clone --depth 1 --branch $this.releaseTag https://github.com/Retro-Carnage-Team/retro-carnage-assets.git $this.releaseTag
        Remove-Item -Recurse -Force (Join-Path $this.workFolder $this.releaseTag ".git")
    }

    [void] BuildArchive($path) {                
        $compress = @{
            Path = Join-Path $this.workFolder $this.releaseTag
            CompressionLevel = "Optimal"
            DestinationPath = $path
        }
        Compress-Archive @compress        
    }

    [string] CreateTempFolder() {
        $guid = [System.Guid]::NewGuid()
        $tmpFolder = [System.IO.Path]::GetTempPath()
        $folder = Join-Path $tmpFolder $guid
        New-Item -ItemType Directory -Path $folder
        return $folder
    }

    [string] GetOutputFileName() {        
        $location = Get-Location
        if([System.Environment]::OSVersion.Platform -eq "Win32NT") {
            return Join-Path $location "Retro-Carnage-$($this.releaseTag)-Windows.zip"
        }
        return Join-Path $location "Retro-Carnage-$($this.releaseTag)-Linux.zip"
    }
}

$location = Get-Location
$appBuilder = [ApplicationBuilder]::new($releaseTag)
$appBuilder.BuildRelease()
Write-Host "Release created: $($appBuilder.outFile)"
Set-Location $location
