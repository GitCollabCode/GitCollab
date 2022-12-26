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

#Uninstall old versions
sudo apt-get -y remove docker docker-engine docker.io containerd runc

#Install packages to allow apt to use a repository over HTTPS
sudo apt-get -y install \
    ca-certificates \
    curl \
    gnupg2 \
    lsb-release \
    apt-transport-https

#Add Docker’s official GPG key
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

#Set up repository
echo \
"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

#Install Docker Engine
sudo chmod a+r /etc/apt/keyrings/docker.gpg
sudo apt-get -y update

#NOTE: Should we install a select versions?
sudo apt-get -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin

#NOTE: WSL has an issue here, still wont allow docker commands without sudo
sudo usermod -aG docker "$USER"

#Start Docker and run hello-world
sudo service docker start
sudo systemctl enable docker #NOTE: systemctl enable does not work within WSL
sleep 3
sudo docker run hello-world
