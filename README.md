# Todo CLI

A beautiful and interactive terminal-based TODO application written in Go.

## ✨ Features

- **Pure Go**: No CGO dependencies - builds everywhere Go does
- **Interactive CLI**: Beautiful terminal interface with modern UX
- **Cross-platform**: Works on Linux, macOS, Windows, FreeBSD
- **SQLite Database**: Persistent storage with pure Go SQLite driver
- **Multiple Architectures**: Supports amd64, arm64, 386, and ARM

## 🚀 Installation

### From Releases
Download the latest binary for your platform from the [releases page](https://github.com/dhirajzope/todo-cli/releases).

### Using Go
```bash
go install github.com/dhirajzope/todo@latest
```

### Using Homebrew (macOS/Linux)
```bash
brew tap dhirajzope/tap
brew install todo
```

### Using Docker
```bash
docker run -it -v $(pwd)/data:/data dhirajzope/todo:latest
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