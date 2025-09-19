# GCAI - AI Commit Message Generator

GCAI is a command-line tool that uses AI to automatically generate commit messages for your git repositories.

## Installation

To install GCAI, you can clone this repository and build the project using Go:

```bash
git clone https://github.com/eduardongomes/gcai.git
cd gcai
go build -o gcai ./cmd
```

This will create a `gcai` binary in the current directory. You can move this binary to a directory in your `PATH` to make it accessible from anywhere on your system.

## Usage

To use GCAI, simply run the `gcai` command in a git repository with staged changes. The tool will generate a commit message and create a commit with it.

```bash
gcai
```

## Configuration

The first time you run GCAI, it will create a configuration file at `~/.config/gcai/config.json`. You will need to add your Gemini API key to this file.

You can also use the `-config` flag to re-configure your API key.

## Flags

- `-config`: A boolean flag that, when present, allows you to set or update your Gemini API key.

## Development

To run the tests for this project, you can use the `go test` command:

```bash
go test ./...
```
