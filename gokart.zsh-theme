#!/usr/bin/env zsh

# Get the path to this file, make it absolute and strip the filename.
# See:  https://stackoverflow.com/questions/9901210/bash-source0-equivalent-in-zsh/28336473#28336473
GOKART_PROMPT_DIR=${${(%):-%x}:A:h}

gokart_prompt_precmd() {
    export EXIT_CODE=$?
    PS1=$("$GOKART_PROMPT_DIR/gokart" ps1)
    PS2=$("$GOKART_PROMPT_DIR/gokart" ps2)
}

gokart_prompt_setup() {
    autoload -Uz add-zsh-hook

    add-zsh-hook precmd gokart_prompt_precmd
}

gokart_prompt_setup
