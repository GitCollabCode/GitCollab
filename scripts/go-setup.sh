#!/bin/sh

#Script only compatible with Ubuntu
if ! grep -q 'ubuntu' /etc/os-release
then
    echo "This Linux distro is not supported!";
    echo "Supported Linux distros: Ubuntu";
    exit 1;
fi

#Update and Upgrade the system
sudo apt -y update
sudo apt -y upgrade

#Fetch go tar into home dir
wget https://golang.org/dl/go1.19.1.linux-amd64.tar.gz -P /tmp/

#TODO: SHA256 checksum of tarball

#Unpack go tar
sudo tar -C /usr/local -xvf /tmp/go1.19.1.linux-amd64.tar.gz

#Setting up .profile export path
if ! grep -q "/usr/local/go/bin" ~/.profile
then
    echo "export PATH=\$PATH:/usr/local/go/bin" >>  ~/.profile
fi

#There is an issue with WSL paths in bash scripts for .profiles and .bashrc files
#https://askubuntu.com/questions/1354999/bad-variable-name-error-on-wsl
if grep -qi microsoft /proc/version
then
    echo "############################################################################"
    echo "  Please manually soruce ~/.profile for WSL environments"
    echo "  Bash is unable to properly process .bashrc and .profile paths within WSL"
    echo "  https://askubuntu.com/questions/1354999/bad-variable-name-error-on-wsl"
    echo "############################################################################"

    rm /tmp/go1.19.1.linux-amd64.tar.gz
    exit 0
fi

. "/$HOME/.profile"

go version

#go env -w GOPRIVATE=github.com/GitCollabCode/*
#git config --global url.git@github.com:.insteadOf https://github.com/

rm /tmp/go1.19.1.linux-amd64.tar.gz