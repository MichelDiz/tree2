$ErrorActionPreference = "Stop"

$repo = "MichelDiz/tree2"
$version = if ($args.Length -gt 0) { $args[0] } else { "latest" }

# Detect OS/ARCH (Windows/amd64)
$os = "windows"
$arch = "amd64"

# Get latest version if not provided
if ($version -eq "latest") {
    $response = Invoke-RestMethod -Uri "https://api.github.com/repos/$repo/releases/latest"
    $version = $response.tag_name
}

Write-Host "Installing tree2 version $version for $os/$arch..."

# Build download URL
$binaryUrl = "https://github.com/$repo/releases/download/$version/tree2-$os-$arch.exe"

# Download binary
$binDir = "$env:USERPROFILE\bin"
if (-Not (Test-Path $binDir)) { New-Item -ItemType Directory -Path $binDir | Out-Null }

$outputPath = "$binDir\tree2.exe"
Invoke-WebRequest -Uri $binaryUrl -OutFile $outputPath

Write-Host "tree2 downloaded to $outputPath"

# Add to PATH (current session)
$env:Path += ";$binDir"

Write-Host "tree2 installed successfully! Try running 'tree2 --help'"
