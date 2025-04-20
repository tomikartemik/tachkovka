#!/bin/bash

SESSION_NAME="tachkovka"
REPO_DIR="$HOME/tachkovka"
GIT_REPO_URL="https://github.com/tomikartemik/tachkovka.git"

# Экспортируем переменную окружения
export SERVER_IP="${SERVER_IP}"

echo "Starting deployment script"

tmux_send() {
    tmux send-keys -t $SESSION_NAME "$1" C-m
}

tmux kill-session -t $SESSION_NAME

echo "Cloning/updating repository"
if [ ! -d "$REPO_DIR" ]; then
    git clone $GIT_REPO_URL $REPO_DIR
fi

cd $REPO_DIR && git pull origin main

cd cmd

tmux new-session -d -s $SESSION_NAME
tmux_send "go run main.go"