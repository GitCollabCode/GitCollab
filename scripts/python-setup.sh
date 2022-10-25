#!/bin/sh


PYTHON_VENV_DIR="$(pwd)/gitcollab_pyenv"

'which python3'

if [ $? -eq 0 ]; then
    echo "Python not installed, installing python3"
    apt install python3 -y
fi

# install venv, and postgres requirements
apt install python3.8-venv -y
#apt install libpq-dev postgresql

if [ ! -d "$PYTHON_VENV_DIR" ]; then
    echo "Creating virtual env"
    python3 -m venv "$PYTHON_VENV_DIR"
fi

# install requirements in virtual environment, pytest
echo "$PYTHON_VENV_DIR/bin/pip3"
"$PYTHON_VENV_DIR/bin/pip3" install -r "$(pwd)/scripts/requirements.txt"
