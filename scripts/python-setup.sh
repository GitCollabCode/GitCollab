#!/bin/sh


if [ "$(basename "$(pwd)")" != "GitCollab" ]
then
    echo "Please run python-setup.sh from the GitCollab repo directory!"
    echo "i.e. ./scripts/python-setup.sh"
    exit 1
fi

PYTHON_VENV_DIR="$(pwd)/gitcollab_pyenv"

'which python3'

if [ $? -eq 0 ]; then
    echo "Python not installed, installing python3"
    apt install python3 -y
fi

# install venv, and postgres requirements
apt-get install -y python3-venv
#apt install libpq-dev postgresql

# Create virtual env. if folder doesn't already exists
if [ ! -d "$PYTHON_VENV_DIR" ]; then
    echo "Creating virtual env"
    python3 -m venv "$PYTHON_VENV_DIR"
fi

# install requirements in virtual environment, pytest
echo "$PYTHON_VENV_DIR/bin/pip3"
"$PYTHON_VENV_DIR/bin/pip3" install -r "$(pwd)/scripts/requirements.txt"