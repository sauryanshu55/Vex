# Vex - A CLI based Agentic AI to assist with Git Version Control system

Vex is an intelligent CLI tool that translates natural language commands into Git operations. Simply describe what you want to do with your repository, and Vex will handle the Git commands for you.

## Features
- **Fast**: Leverages Go's concurrency to process API calls, command generation, and execution in parallel, reducing floor times for Vex's usage
- **Natural Language Interface**: Use plain English to perform Git operations
- **Smart Prerequisites**: Automatically detects and handles missing setup steps
- **Safety**: Warns about potentially risky operations that could cause data loss

## Installation

Simply download the executables from the releases for your OS (configured for amd64 bits), add it to your PATH variable


## Requirements

- None, you could even ask Vex to install git in your system!

## Usage

The basic syntax is simple:

```bash
vex "describe what you want to do. Any git command, really"
```

## Examples - Non Exhaustive

### Initialize a Repository
```bash
$> vex "initialize a repository here"
# Executes: git init
```

### Commit Code
```bash
$> vex "commit this code"
$> What should be the commit message?
$> Fixed debug error
# Executes: git add .
#          git commit -m "Fixed debug error"
```

### Smart Prerequisites
```bash
$> vex "commit this code"
$> You need to initialize your repository first. Can I do that? (Y/N)
$> Y
$> What should be the commit message?
$> Fixed debug error
# Executes: git init
#          git add .
#          git commit -m "Fixed debug error"
```

### Branch Operations
```bash
$> vex "create new branch feature"
# Executes: git checkout -b feature

$> vex "switch to branch main"
$> You have uncommitted changes in this branch. Do you want me to stash these changes or commit them first?
$> Commit
$> What should be the commit message?
$> Work in progress
# Executes: git add .
#          git commit -m "Work in progress"
#          git checkout main
#          git pull origin main
```

### Repository Status
```bash
$> vex "show status"
# Executes: git status

$> vex "show recent commits"
# Executes: git log --oneline -10
```

### Push Changes
```bash
$> vex "push current branch"
# Executes: git push origin HEAD
```

### Pull Updates
```bash
$> vex "pull latest changes"
# Executes: git pull
```

### Undo Operations
```bash
$> vex "undo last commit"
$> This operation may cause data loss. Continue? (Y/N)
$> Y
# Executes: git reset --soft HEAD~1
```

## Safety and Data Loss warnings Features

Vex prioritizes data safety and will:

- Detect missing repository initialization
- Warn about uncommitted changes before branch switches
- Confirm potentially destructive operations
-  Ask for commit messages when needed
- Validate prerequisites before executing commands

