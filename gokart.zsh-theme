#!/usr/bin/env zsh

# Get the path to this file, make it absolute and strip the filename.
# See:  https://stackoverflow.com/questions/9901210/bash-source0-equivalent-in-zsh/28336473#28336473
GOKART_DIR=${${(%):-%x}:A:h}

gokart_powerline_precmd() {
    export EXIT_CODE=$?
    PS1=$("$GOKART_DIR/gokart" ps1)
    PS2=$("$GOKART_DIR/gokart" ps2)

    local NEWLINE=$'\n'
    PROMPT="${PS1}${NEWLINE}${PS2}"
}

gokart_powerline_setup() {
    autoload -Uz add-zsh-hook

    add-zsh-hook precmd gokart_powerline_precmd
}

gokart_powerline_setup
