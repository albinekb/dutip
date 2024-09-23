# dutip 🔄🖥️

Only works on macOS.

A simple command-line tool to change the default application for file extensions using [duti](https://github.com/moretension/duti) on macOS. 🍎

For example when switching editors from VSCode to Cursor.

## Features ✨

- Change default applications for multiple file extensions at once 📂🔀
- Support for both application names and bundle identifiers 🏷️
- Force mode to skip confirmation prompts ⚡
- Version information display ℹ️

## Installation 🛠️

1. Ensure you have [duti](https://github.com/moretension/duti) installed on your system. 🔧
2. Install dutip: 📥

## Usage 🚀

```sh
dutip -from com.microsoft.VSCodeInsiders -to Cursor.app
```

```sh
dutip -from Cursor.app -to com.microsoft.VSCodeInsiders
```
