# Go Shell

A simple command-line shell implementation written in Go.

## Description

This project implements a basic interactive shell that allows users to execute system commands from a command-line interface. The shell provides a read-eval-print loop (REPL) that continuously prompts for user input and executes the entered commands.

## Features

- **Interactive Command Execution**: Run any system command available on your PATH
- **Built-in Commands**:
  - `cd <directory>` - Change the current working directory
  - `exit` - Terminate the shell
- **Error Handling**: Graceful error reporting for invalid commands or missing arguments
- **Real-time Output**: Commands display their output directly to the terminal

## Usage

1. **Compile and run the program**:
   ```bash
   go run main.go
   ```

2. **Use the shell**:
   ```
   Enter command: ls -la
   Enter command: cd /tmp
   Enter command: pwd
   Enter command: exit
   ```

## Examples

```bash
# List files in current directory
Enter command: ls

# Change to home directory
Enter command: cd ~

# Run commands with arguments
Enter command: echo "Hello, World!"

# Exit the shell
Enter command: exit
```


## Requirements

- Go 1.x or higher
- Unix-like operating system (Linux, macOS) or Windows with appropriate command-line tools
