#!/usr/bin/env zsh

# Get the path to this file, make it absolute and strip the filename.
# See: https://stackoverflow.com/questions/9901210/bash-source0-equivalent-in-zsh/28336473#28336473
GOKART_PROMPT_DIR=${${(%):-%x}:A:h}

gokart_prompt_precmd() {
    # Store the current time, i.e. the previous command end time.
    export GOKART_CMD_END=$EPOCHREALTIME

    export EXIT_CODE=$?
    export GOKART_SHELL=zsh

    PS1=$("$GOKART_PROMPT_DIR/gokart" ps1)
    PS2=$("$GOKART_PROMPT_DIR/gokart" ps2)

    # Clear the entered command.
    export GOKART_CMD=
}

gokart_prompt_preexec() {
    # Get and store the command that was entered.
    # https://zsh.sourceforge.io/Doc/Release/Functions.html#index-preexec_005ffunctions
    export GOKART_CMD="$1"

    # Store the command start time.
    export GOKART_CMD_START=$EPOCHREALTIME
}

gokart_prompt_setup() {
    autoload -Uz add-zsh-hook

    add-zsh-hook precmd gokart_prompt_precmd
    add-zsh-hook preexec gokart_prompt_preexec
}

gokart_prompt_setup
