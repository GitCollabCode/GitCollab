#!/bin/bash

#Script was written quickly and might need a bit more work

#You might need to set "StrictHostKeyChecking no" in your /etc/ssh/ssh_config file

if grep -q 'ubuntu' /etc/os-release; then
    if ! command -v mosquitto_sub &> /dev/null; then
        sudo apt-get -y install mosquitto-clients
    fi
    
elif grep -q 'centos' /etc/os-release; then
    if ! command -v mosquitto_sub &> /dev/null; then
        sudo yum -y install mosquitto
        sudo systemctl start mosquitto
        sudo systemctl enable mosquitto
    fi
fi

SERVER="$1"
USERNAME="$2"
PASSWORD="$3"
GICOLLAB_REPO_DIR="$4"

if [ -z "$SERVER" ]; then
    echo "First argument is not set, expecting mqtt server address!"
    exit 1
fi

if [ -z "$USERNAME" ]; then
    echo "Second argument is not set, expecting mqtt server username!"
    exit 1
fi

if [ -z "$PASSWORD" ]; then
    echo "Third argument is not set, expecting mqtt server password!"
    exit 1
fi

if [ ! -d "$GICOLLAB_REPO_DIR" ]; then
    echo "Fourth argument is not a path to directory, expecting path to GitCollab repo!"
    exit 1
fi

while true
do 
    DATE=$(date)
    echo "[$DATE] Waiting for MQTT message, from dev-server topic"
    message=""
    read -r message < <(mosquitto_sub -h "$SERVER" -C 1 -t "dev-server")
    if [ $? -ne 0 ]; then
        echo "mosquitto_sub returned an error, exiting!"
        exit 1
    fi

    if [ "$message" = "update" ]; then
        DATE=$(date)
        echo "[$DATE] Updating live server build!"
        (cd "$GICOLLAB_REPO_DIR" && docker compose down || exit)
        (cd "$GICOLLAB_REPO_DIR" && git pull || exit)
        (cd "$GICOLLAB_REPO_DIR" && docker compose up -d || exit)
    elif [ "$message" = "terminate" ]; then
        exit 0
    else
        echo "Unkown message [$message] received, ignoring!"
    fi
done