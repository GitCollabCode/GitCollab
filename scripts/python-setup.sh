#!/bin/sh


PYTHON_VENV_DIR="$(pwd)/gitcollab_pyenv"
PYTHON_VENV_ACIVATE="$(pwd)/gitcollab_pyenv/bin/activate"

'which python3'

if [ $? -eq 0 ]; then
    echo "Python not installed, installing python3"
    apt install python3 -y
fi

# install venv
apt install python3.8-venv -y

if [ ! -d "$PYTHON_VENV_DIR" ]; then
    echo "Creating virtual env"
    python3 -m venv "$PYTHON_VENV_DIR"
fi

# install requirements in virtual environment, pytest
. "$PYTHON_VENV_ACIVATE"
pip3 install -r "$(find "$(pwd)" | grep requirements.txt)"
deactivate
