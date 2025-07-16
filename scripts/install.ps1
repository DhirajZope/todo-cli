<#
.SYNOPSIS
  Installs the todo-cli into %USERPROFILE%\bin and updates User PATH.

.PARAMETER Version
  The release tag to install (e.g. v1.2.3). Defaults to "latest".

.EXAMPLE
  # Install latest
  .\install.ps1

  # Install specific version
  .\install.ps1 -Version v1.0.0
#>

param(
    [string]$Version = "latest"
)

# Repo and binary details
$repo   = "dhirajzope/todo-cli"
$binary = "todo.exe"
$os      = "windows"
$arch    = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

# Determine download URL
if ($Version -eq "latest") {
    Write-Host "Fetching latest release info..."
    $rel   = Invoke-RestMethod "https://api.github.com/repos/$repo/releases/latest"
    $asset = $rel.assets |
        Where-Object { $_.name -like "*_${os}_${arch}.zip" } |
        Select-Object -First 1 -ExpandProperty browser_download_url
} else {
    $asset = "https://github.com/$repo/releases/download/$Version/todo_${Version}_${os}_${arch}.zip"
}

if (-not $asset) {
    Write-Error "Could not find a matching asset for $os/$arch. Exiting."
    exit 1
}

Write-Host "Downloading $asset..."
$zip = Join-Path $env:TEMP "todo.zip"
Invoke-WebRequest -Uri $asset -OutFile $zip -UseBasicParsing

# Prepare destination folder
$dest = Join-Path $env:USERPROFILE "bin"
if (-not (Test-Path $dest)) {
    Write-Host "Creating install directory $dest..."
    New-Item -ItemType Directory -Path $dest | Out-Null
}

# Extract and clean up
Write-Host "Extracting to $dest..."
Expand-Archive -Path $zip -DestinationPath $dest -Force
Remove-Item $zip

# Ensure $dest is on the user PATH
$currentPath = [Environment]::GetEnvironmentVariable('Path','User') -split ';' | Where-Object { $_ -ne "" }
if ($currentPath -notcontains $dest) {
    Write-Host "Adding $dest to your User PATH..."
    $newPath = ($currentPath + $dest) -join ';'
    [Environment]::SetEnvironmentVariable('Path', $newPath, 'User')
    Write-Host "✅ Added $dest to your User PATH. Restart your shell or log out/in to use 'todo' command."
} else {
    Write-Host "$dest is already in your User PATH."
}

Write-Host "✅ Installed $binary to $dest"
