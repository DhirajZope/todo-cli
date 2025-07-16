param([string]$Version = "latest")

$repo = "dhirajzope/todo-cli"
$os   = "windows"
$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

if ($Version -eq "latest") {
  $rel = Invoke-RestMethod "https://api.github.com/repos/$repo/releases/latest"
  $asset = $rel.assets | Where-Object { $_.name -like "*_${os}_${arch}.zip" } |
           Select-Object -First 1 -ExpandProperty browser_download_url
} else {
  $asset = "https://github.com/$repo/releases/download/$Version/todo_${Version}_${os}_${arch}.zip"
}

Write-Host "Downloading $asset"
$zip = "$env:TEMP\todo.zip"
Invoke-WebRequest -Uri $asset -OutFile $zip

$dest = "$env:USERPROFILE\bin"
if (-not (Test-Path $dest)) { New-Item -ItemType Directory -Path $dest | Out-Null }
Expand-Archive -Path $zip -DestinationPath $dest -Force

Write-Host "✅ Installed todo.exe to $dest"
Write-Host "Add `$dest` to your PATH if not already."
