# Gokart Prompt 🏎

**Blazing fast and slick shell prompt, written in Go**

[![Build](https://github.com/LucaScorpion/gokart-prompt/actions/workflows/build.yml/badge.svg)](https://github.com/LucaScorpion/gokart-prompt/actions/workflows/build.yml)

<div align="center">
  <img src="screenshot.png" alt="Screenshot of a terminal with the Gokart prompt">
</div>

## Installation 🚀

First download the archive for your platform from the [latest release](https://github.com/LucaScorpion/gokart-prompt/releases/latest).
Next, follow the steps for your terminal setup:

<details>
<summary>Oh My Zsh</summary>

Extract the archive to `$ZSH_CUSTOM/themes`:

```shell
tar xzvf gokart-prompt-linux-amd64.tar.gz -C "$ZSH_CUSTOM/themes"
```

Make sure to replace `gokart-prompt-linux-amd64.tar.gz` with the filename of the archive you downloaded.

Next, symlink `gokart.zsh-theme` in your `$ZSH_CUSTOM/themes` directory:

```shell
ln -s "$ZSH_CUSTOM/themes/gokart-prompt/gokart.zsh-theme" "$ZSH_CUSTOM/themes/gokart.zsh-theme"
```

Finally, set `ZSH_THEME="gokart"` in your `.zshrc`.
</details>

<details>
<summary>Bash</summary>

Extract the archive, for example to your `$HOME` directory:

```shell
tar xzvf gokart-prompt-linux-amd64.tar.gz -C "$HOME"
```

Make sure to replace `gokart-prompt-linux-amd64.tar.gz` with the filename of the archive you downloaded.

Finally, source `gokart.bash-theme` in your `.bashrc`:

```shell
source "$HOME/gokart-prompt/gokart.bash-theme"
```
</details>

## Requirements ✅

### [Powerline](https://github.com/powerline/fonts) or [Nerd Font](https://www.nerdfonts.com)

To check if the font works, run:

```shell
echo -e "\xee\x82\xa0"
```

This should print a "branch" icon.

## Features ✨

- Git branch and status information
- Command exit code
- Command execution time[^zsh-only]
- Version information for various tools and languages:
  - Docker
  - .NET
  - Go
  - Java
  - Lua
  - Node.js
  - PHP
  - Python
  - Ruby
  - Rust
- Expected version and mismatch indicator (⚠️) for:
  - Go (`go.mod`)
  - Node.js (`.nvmrc`)

[^zsh-only]: Only in Zsh, due to limitations of Bash. 

Is something missing?
Feel free to [open an issue](https://github.com/LucaScorpion/gokart-prompt/issues/new) or send a pull request!
