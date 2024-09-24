# dutip

A simple command-line tool to change the default application for file extensions using [duti](https://github.com/moretension/duti) on macOS. 🍎

## Installation

### Using Go

```shell
go install github.com/albinekb/dutip@latest
```

## Usage

```shell
dutip
```

When trying out different code editors, you can use the following commands to switch between Cursor and VSCode:

```sh
dutip -from com.microsoft.VSCodeInsiders -to Cursor.app
```

```sh
dutip -from Cursor.app -to com.microsoft.VSCodeInsiders
```

## Features ✨

- Change default applications for multiple file extensions at once 📂🔀
- Support for both application names and bundle identifiers 🏷️
- Force mode to skip confirmation prompts ⚡
- Version information display ℹ️

## Installation 🛠️

1. Ensure you have [duti](https://github.com/moretension/duti) & [Go](https://go.dev/doc/install) installed on your system. 🔧

```sh
brew install duti go
```

2. Install dutip: 📥

```sh
go install github.com/albinekb/dutip@latest
```

3. Run dutip: 🚀

```sh
dutip
```
