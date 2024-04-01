# This script creates the monthly release tag for this repository.
#
# Usage:
# 
# PS> .\create-release.ps1
#
# This script has dependencies that have to be installed first:
#
# All platforms:
# * git

$tagName = Get-Date -Format "vyyyy.MM"
git tag -a $tagName -m "Pre release for this month"
git push origin $tagName
