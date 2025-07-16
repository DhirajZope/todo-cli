# todo-cli

A simple terminal-based TODO application built in Go, distributed as pre-built binaries and packages via GitHub Releases.

---

## 🚀 Quick Install

### Linux & macOS

```bash
# install latest
curl -sL https://raw.githubusercontent.com/yourusername/todo-cli/main/scripts/install.sh | bash

# install specific version
curl -sL https://raw.githubusercontent.com/yourusername/todo-cli/main/scripts/install.sh | bash -s v1.2.0

```
### Windows

```bash
# install latest
irm https://raw.githubusercontent.com/yourusername/todo-cli/main/scripts/install.ps1 | iex

# install specific version
irm https://raw.githubusercontent.com/yourusername/todo-cli/main/scripts/install.ps1 | iex -Version v1.2.0
```

### Or Download manually

1. Go to [Releases](https://github.com/dhirajzope/todo-cli/releases)

2. Download the appropriate asset:

    - ```*.tar.gz``` or ```*.zip``` for generic binaries

    - ```*.deb``` for Debian/Ubuntu

    - ```*.rpm``` for Fedora/RHEL

3. Extract or install:

```bash
# example for Debian/Ubuntu:
sudo dpkg -i todo_1.2.0_amd64.deb
```

### Usage

```bash
todo --help
```

List Management
```bash
todo list create "Work"
todo list all
todo list delete 1
```

Task Management
```bash
todo task add 1 "Finish report"
todo task list 1
todo task complete 2
todo task delete 2
```