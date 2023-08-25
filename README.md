# Gokart Prompt 🏎

**Blazing fast and slick shell prompt, written in Go**

[![Build](https://github.com/LucaScorpion/gokart-prompt/actions/workflows/build.yml/badge.svg)](https://github.com/LucaScorpion/gokart-prompt/actions/workflows/build.yml)

<div align="center">
  <img src="screenshot.png" alt="Screenshot of a terminal with the Gokart prompt">
</div>

## Installation 🚀

<details>
<summary>Oh My Zsh</summary>

Download and extract the latest release:

```shell
curl -fsSL https://github.com/LucaScorpion/gokart-prompt/releases/latest/download/gokart-prompt.tar.gz | tar xzvf - -C "$ZSH_CUSTOM/themes"
```

Symlink `gokart.zsh-theme` in your themes directory:

```shell
ln -s "$ZSH_CUSTOM/themes/gokart-prompt/gokart.zsh-theme" "$ZSH_CUSTOM/themes/gokart.zsh-theme"
```

Set `ZSH_THEME="gokart"` in your `.zshrc`.
</details>

<details>
<summary>Bash</summary>

Download and extract the latest release:

```shell
curl -fsSL https://github.com/LucaScorpion/gokart-prompt/releases/latest/download/gokart-prompt.tar.gz | tar xzvf - -C "$HOME"
```

Source `gokart.bash-theme` in your `.bashrc`:

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
