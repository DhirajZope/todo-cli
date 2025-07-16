# Todo CLI

A beautiful and interactive terminal-based TODO application written in Go.

## ✨ Features

- **Pure Go**: No CGO dependencies - builds everywhere Go does
- **Interactive CLI**: Beautiful terminal interface with modern UX
- **Cross-platform**: Works on Linux, macOS, Windows, FreeBSD
- **SQLite Database**: Persistent storage with pure Go SQLite driver
- **Multiple Architectures**: Supports amd64, arm64, 386, and ARM

## 🚀 Installation

### 📦 Binary Releases (Recommended)

Download pre-built binaries for your platform from the [releases page](https://github.com/DhirajZope/todo-cli/releases).

#### Windows

##### PowerShell (Recommended)
```powershell
# Download and extract (replace VERSION with latest release version)
$version = "VERSION"  # Replace with actual version like "v1.0.0"
$url = "https://github.com/DhirajZope/todo-cli/releases/download/$version/todo-windows-amd64.zip"
Invoke-WebRequest -Uri $url -OutFile "todo-windows-amd64.zip"
Expand-Archive -Path "todo-windows-amd64.zip" -DestinationPath "."
Move-Item "todo-windows-amd64\todo.exe" "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps\todo.exe"

# Alternative: Install to a custom directory in PATH
# New-Item -Path "C:\tools" -ItemType Directory -Force
# Move-Item "todo-windows-amd64\todo.exe" "C:\tools\todo.exe"
# [Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\tools", "User")
```

##### Command Prompt (CMD)
```cmd
REM Download using curl (if available) - replace VERSION with actual version
curl.exe -L -o todo-windows-amd64.zip https://github.com/DhirajZope/todo-cli/releases/download/vVERSION/todo-windows-amd64.zip

REM Extract (requires 7-zip, WinRAR, or Windows 10+ built-in support)
tar -xf todo-windows-amd64.zip

REM Move to PATH directory
move todo-windows-amd64\todo.exe %USERPROFILE%\AppData\Local\Microsoft\WindowsApps\
```

##### Using Chocolatey (Alternative)
```powershell
# Install Chocolatey first if not installed:
# Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))

# Then install todo-cli (when available)
# choco install todo-cli
```

##### Manual Installation (No Command Line)
1. Go to [releases page](https://github.com/DhirajZope/todo-cli/releases)
2. Download `todo-windows-amd64.zip` (or `todo-windows-386.zip` for 32-bit)
3. Right-click the zip file → "Extract All..."
4. Copy `todo.exe` from the extracted folder
5. Paste it into one of these directories:
   - `%USERPROFILE%\AppData\Local\Microsoft\WindowsApps\` (recommended)
   - `C:\Windows\System32\` (requires admin rights)
   - Or create `C:\tools\` and add it to your PATH

#### macOS
```bash
# Intel Macs
curl -L -o todo-darwin-amd64.tar.gz https://github.com/DhirajZope/todo-cli/releases/download/vVERSION/todo-darwin-amd64.tar.gz
tar -xzf todo-darwin-amd64.tar.gz
sudo mv todo /usr/local/bin/

# Apple Silicon Macs (M1/M2)
curl -L -o todo-darwin-arm64.tar.gz https://github.com/DhirajZope/todo-cli/releases/download/vVERSION/todo-darwin-arm64.tar.gz
tar -xzf todo-darwin-arm64.tar.gz
sudo mv todo /usr/local/bin/
```

#### Linux
```bash
# x86_64
curl -L -o todo-linux-amd64.tar.gz https://github.com/DhirajZope/todo-cli/releases/download/vVERSION/todo-linux-amd64.tar.gz
tar -xzf todo-linux-amd64.tar.gz
sudo mv todo /usr/local/bin/

# ARM64
curl -L -o todo-linux-arm64.tar.gz https://github.com/DhirajZope/todo-cli/releases/download/vVERSION/todo-linux-arm64.tar.gz
tar -xzf todo-linux-arm64.tar.gz
sudo mv todo /usr/local/bin/

# 32-bit x86
curl -L -o todo-linux-386.tar.gz https://github.com/DhirajZope/todo-cli/releases/download/vVERSION/todo-linux-386.tar.gz
tar -xzf todo-linux-386.tar.gz
sudo mv todo /usr/local/bin/
```

#### FreeBSD
```bash
# x86_64
curl -L -o todo-freebsd-amd64.tar.gz https://github.com/DhirajZope/todo-cli/releases/download/vVERSION/todo-freebsd-amd64.tar.gz
tar -xzf todo-freebsd-amd64.tar.gz
sudo mv todo /usr/local/bin/
```

### 🛠 Build from Source

#### Prerequisites
- Go 1.23 or later
- Git

#### Using Go Install
```bash
go install github.com/DhirajZope/todo-cli@latest
```

#### Manual Build
```bash
# Clone the repository
git clone https://github.com/DhirajZope/todo-cli.git
cd todo-cli

# Build for your platform
go build -o todo

# Or build with version info
go build -ldflags "-X main.version=$(git describe --tags --always)" -o todo

# Install to Go bin directory
go install
```

### 🔧 Manual Installation

1. **Download**: Go to [releases page](https://github.com/DhirajZope/todo-cli/releases)
2. **Choose**: Select the appropriate binary for your OS and architecture:
   - `todo-windows-amd64.zip` - Windows 64-bit
   - `todo-windows-386.zip` - Windows 32-bit
   - `todo-darwin-amd64.tar.gz` - macOS Intel
   - `todo-darwin-arm64.tar.gz` - macOS Apple Silicon
   - `todo-linux-amd64.tar.gz` - Linux 64-bit
   - `todo-linux-arm64.tar.gz` - Linux ARM64
   - `todo-linux-386.tar.gz` - Linux 32-bit
   - `todo-freebsd-amd64.tar.gz` - FreeBSD 64-bit
3. **Extract**: Unzip/untar the downloaded file
4. **Install**: Move the binary to a directory in your PATH
5. **Verify**: Run `todo --version` to confirm installation

### ✅ Verify Installation

After installation, verify that todo is working:

```bash
# Check version
todo --version

# Show help
todo --help

# Create your first list
todo create "My First List"
```

## 📖 Usage

### Basic Commands

```bash
# List all todo lists
todo list all

# Create a new list
todo create "My Tasks"

# Add a task to a list
todo task add 1 "Complete the project"

# List tasks in a list
todo task list 1

# Complete a task
todo complete 1

# Delete a task
todo task delete 1
```

### Environment Variables

- `DB_PATH`: Set custom database location (default: `todo.db`)

## 🛠 Development

### Prerequisites
- Go 1.23+

### Building
```bash
go build
```

### Running Tests
```bash
go test ./...
```

### Cross-compilation
Pure Go build means easy cross-compilation:
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o todo.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o todo

# macOS
GOOS=darwin GOARCH=amd64 go build -o todo
```

## 🏗 Architecture

- **Pure Go SQLite**: Uses `modernc.org/sqlite` for CGO-free database operations
- **Cobra CLI**: Modern CLI framework for command structure
- **Bubble Tea**: Interactive terminal UI framework (coming soon)
- **Lipgloss**: Styling and layout for beautiful terminal interfaces (coming soon)

## 📦 Release Process

The project uses GitHub Actions for automated releases:
- Triggered on git tags (v*)
- Builds for multiple platforms automatically
- Publishes to GitHub Releases, Package registries, and Docker Hub
- No CGO complications - builds everywhere!

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.