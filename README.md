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
```powershell
# Download and extract (replace VERSION with latest release version)
curl -L -o todo-windows-amd64.zip https://github.com/DhirajZope/todo-cli/releases/download/vVERSION/todo-windows-amd64.zip
Expand-Archive todo-windows-amd64.zip
move todo-windows-amd64\todo.exe C:\Windows\System32\todo.exe

# Or manually:
# 1. Download todo-windows-amd64.zip from releases
# 2. Extract todo.exe
# 3. Move to a directory in your PATH (e.g., C:\Windows\System32)
```

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