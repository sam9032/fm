<p align="center">
  <img src="./assets/logo.svg" height="180" width="180" />
  <p align="center">
    Keep those files organized
  </p>
  <p align="center">
    <a href="https://github.com/knipferrc/fm/releases"><img src="https://img.shields.io/github/v/release/knipferrc/fm" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/knipferrc/fm?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/knipferrc/fm/actions"><img src="https://img.shields.io/github/workflow/status/knipferrc/fm/Release" alt="Build Status"></a>
  </p>
</p>

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
  <img src="./assets/default.png" width="450" alt="default screenshot">
</p>

## About The Project

A terminal based file manager

### Built With

- [Go](https://golang.org/)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [bubbles](https://github.com/charmbracelet/bubbles)
- [lipgloss](https://github.com/charmbracelet/lipgloss)
- [Glamour](https://github.com/charmbracelet/glamour)
- [Chroma](https://github.com/alecthomas/chroma)
- [Viper](https://github.com/spf13/viper)
- [Cobra](https://github.com/spf13/cobra)

## Installation

### Curl

```sh
curl -sfL https://raw.githubusercontent.com/knipferrc/fm/main/install.sh | sh
```

### Go

```
go install github.com/knipferrc/fm@latest
```

### AUR

Install through the Arch User Repository with your favorite AUR helper.
There are currently two possible packages:

- [fm-git](https://aur.archlinux.org/packages/fm-git/): Builds the package from the main branch

```sh
paru -S fm-git
```

- [fm-bin](https://aur.archlinux.org/packages/fm-bin/): Uses the github release package

```sh
paru -S fm-bin
```

## Features

- Double pane layout
- File icons
- Layout adjusts to terminal resize
- Syntax highlighting for source code with customizable themes using styles from [chroma](https://swapoff.org/chroma/playground/) (dracula, monokai etc.)
- Render pretty markdown
- Mouse support
- Themes (`default`, `gruvbox`, `spooky`, `nord`, `holiday`)
- Render PNG, JPG and JPEG as strings
- Colors adapt to terminal background
- Open selected file in editor set in EDITOR environment variable
- Preview a directory in the secondary pane
- Copy selected directory items path to the clipboard
- Read PDF files
- Experimental find files/directories (might be a little buggy)
- Simple mode (removes secondary box, hides borders, hide file icons and no colors)

## Themes

### Default

<img src="./assets/default.png" width="350" alt="default">

### Gruvbox

<img src="./assets/gruvbox.png" width="350" alt="gruvbox">

### Spooky

<img src="./assets/spooky.png" width="350" alt="spooky">

### Nord

<img src="./assets/nord.png" width="350" alt="nord">

### Holiday

<img src="./assets/holiday.png" width="350" alt="holiday">

## Simple Mode

<img src="./assets/simple_mode.png" width="350" alt="simple mode">

## Usage

- `fm` will start fm in the current directory
- `fm update` will update fm to the latest version
- `fm --start-dir=/some/start/dir` will start fm in the specified directory
- `fm --selection-path=/tmp/tmpfile` will write the selected items path to the selection path when pressing <kbd>E</kbd> and exit fm

## Navigation

| Key                   | Description                                                                                                                                                                                                                                                      |
| --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <kbd>h or left</kbd>  | Go back to previous directory                                                                                                                                                                                                                                    |
| <kbd> or down</kbd>   | Move down in the file tree or scroll pane down                                                                                                                                                                                                                   |
| <kbd>k or up</kbd>    | Move up in the file tree or scroll pane up                                                                                                                                                                                                                       |
| <kbd>l or right</kbd> | Opens the currently selected directory or file                                                                                                                                                                                                                   |
| <kbd>gg</kbd>         | Jump to bottom of file tree or pane                                                                                                                                                                                                                              |
| <kbd>G</kbd>          | Jump to top of file tree or pane                                                                                                                                                                                                                                 |
| <kbd>~</kbd>          | Go to home directory                                                                                                                                                                                                                                             |
| <kbd>/</kbd>          | Go to the root directory                                                                                                                                                                                                                                         |
| <kbd>.</kbd>          | Toggle hidden files and directories                                                                                                                                                                                                                              |
| <kbd>-</kbd>          | Go to previous directory                                                                                                                                                                                                                                         |
| <kbd>ctrl+c</kbd>     | Exit                                                                                                                                                                                                                                                             |
| <kbd>q</kbd>          | Exit if command bar is not open                                                                                                                                                                                                                                  |
| <kbd>m</kbd>          | Move the currently selected file or directory. Once pressed, the file manager enters move mode. Navigate the tree as usual and press enter in the desired destination directory. It will navigate back to the starting direcotry in which the move was initiated |
| <kbd>tab</kbd>        | Toggle between panes                                                                                                                                                                                                                                             |
| <kbd>esc</kbd>        | Reset FM to its initial state                                                                                                                                                                                                                                    |
| <kbd>z</kbd>          | Create a zip file of the currently selected directory item                                                                                                                                                                                                       |
| <kbd>u</kbd>          | Unzip a zip file                                                                                                                                                                                                                                                 |
| <kbd>c</kbd>          | Create a copy of a file or directory                                                                                                                                                                                                                             |
| <kbd>ctrl+d</kbd>     | Delete the currently selected file or directory                                                                                                                                                                                                                  |
| <kbd>n</kbd>          | Create a new file in the current directory                                                                                                                                                                                                                       |
| <kbd>N</kbd>          | Create a new directory in the current directory                                                                                                                                                                                                                  |
| <kbd>r</kbd>          | Rename the currently selected file or directory                                                                                                                                                                                                                  |
| <kbd>E</kbd>          | Open in editor set in EDITOR environment variable                                                                                                                                                                                                                |
| <kbd>p</kbd>          | Preview a directory in the secondary pane                                                                                                                                                                                                                        |
| <kbd>y</kbd>          | Copy selected directory items path to the clipboard                                                                                                                                                                                                              |
| <kbd>ctrl+f</kbd>     | Find files and directories in working directory                                                                                                                                                                                                                  |
| <kbd>?</kbd>          | Toggle help screen in simple mode                                                                                                                                                                                                                                |
| <kbd>O</kbd>          | Open logger when logging is enabled                                                                                                                                                                                                                              |
| <kbd>rc</kbd>         | Reload config                                                                                                                                                                                                                                                    |

## Configuration

- A config file will be generated at `~/.config/fm.yml` when you first run `fm` (On windows it will be `C\:\\Users\\username\\fm.yml`)

```yml
settings:
  borderless: false
  enable_logging: false
  enable_mousewheel: true
  pretty_markdown: true
  show_icons: true
  simple_mode: false
  start_dir: .
  syntax_theme: dracula
  theme: default
```

## Local Development

Follow the instructions below to get setup for local development

1. Clone the repo

```sh
git clone https://github.com/knipferrc/fm
```

2. Run

```sh
make
```

3. Build a binary

```sh
make build
```

## Credit

- Thank you to this repo https://github.com/Yash-Handa/logo-ls for the icons
