#!/usr/bin/env zsh

# Get the path to this file, make it absolute and strip the filename.
# See: https://stackoverflow.com/questions/9901210/bash-source0-equivalent-in-zsh/28336473#28336473
GOKART_PROMPT_DIR=${${(%):-%x}:A:h}

gokart_prompt_precmd() {
    export EXIT_CODE=$?
    export GOKART_SHELL=zsh

    PS1=$("$GOKART_PROMPT_DIR/gokart" ps1)
    PS2=$("$GOKART_PROMPT_DIR/gokart" ps2)
}

gokart_prompt_preexec() {
    # TODO: Here we get the entered command.
    # https://zsh.sourceforge.io/Doc/Release/Functions.html#index-preexec_005ffunctions
    # If this is empty, we should also reset the timer.
    echo one
    echo $1
    echo two
    echo $2
    echo three
    echo $3
}

gokart_prompt_set_time() {
    export GOKART_TIME="$EPOCHREALTIME"
}

gokart_prompt_setup() {
    autoload -Uz add-zsh-hook

    add-zsh-hook precmd gokart_prompt_precmd
    add-zsh-hook preexec gokart_prompt_preexec
    add-zsh-hook preexec gokart_prompt_set_time
    add-zsh-hook chpwd gokart_prompt_set_time
}

gokart_prompt_setup
