#!/usr/bin/env bash

if [ "$(basename "$(pwd)")" != "GitCollab" ]
then
    echo "Please run gitcollab.sh from the GitCollab repo directory!"
    echo "i.e. ./scripts/gitcollab.sh [options] arguments"
    exit 1
fi

# Enable xtrace if the DEBUG environment variable is set
if [[ ${DEBUG-} =~ ^1|yes|true$ ]]; then
    set -o xtrace       # Trace the execution of the script (debug)
fi

if ! (return 0 2> /dev/null); then
    # A better class of script...
    set -o errexit      # Exit on most errors (see the manual)
    set -o nounset      # Disallow expansion of unset variables
    set -o pipefail     # Use last non-zero exit code in a pipeline
fi

# Enable errtrace or the error trap handler will not work as expected
set -o errtrace         # Ensure the error trap handler is inherited

function script_usage() {
    cat << EOF
GitCollab deployment script, used to control the docker functions of the project
    Usage:
            gitcollab.sh [options] arguments
    Options:
        -h | --help              Displays this help
        -v | --verbose           Displays verbose output, display docker logs
    Arugments:
        build                    Disables colour output
        start                    Run silently unless we encounter an error
EOF
}

function build() {
    echo "Building GitCollab docker images..."
    docker compose -f "$(pwd)/docker-compose-convert.yaml" build
}

function start() {
    if [ "$is_verbose" = "true" ]; then
        echo "Starting GitCollab docker containers [verbose]..."
        docker compose -f "$(pwd)/docker-compose-convert.yaml" up
    else
        echo "Starting GitCollab docker containers..."
        docker compose -f "$(pwd)/docker-compose-convert.yaml" up -d 
    fi
}

function restart() {
    echo "Restarting GitCollab docker containers..."
    docker compose down
    docker compose -f "$(pwd)/docker-compose-convert.yaml" up -d 
}

function stop() {
    echo "Stopping active GitCollab docker containers..."
    docker compose stop
}

function clean() {
    echo "Taking down active GitCollab docker containers..."
    docker compose down
    echo "Docker system prune..."
    docker system prune -a
}

function clean-db() {
    echo "Removing saved postgres data from $(pwd)/db_data..."
    sudo rm -rfd "$(pwd)/db_data"
}

function refresh-env-file() {
    echo "Refreshing $(pwd)/.env..."
    cp "$(pwd)/env" "$(pwd)/.env"
    chmod 777 "$(pwd)/.env" # change to write permission
}

function parse_params() {
    local param
    while [[ $# -gt 0 ]]; do
        param="$1"
        shift
        case $param in
            -h | --help)
                script_usage
                exit 0
                ;;
            -v | --verbose)
                is_verbose=true
                ;;
            build)
                build
                ;;
            start)
                start
                ;;
            restart)
                restart
                ;;
            stop)
                stop
                ;;
            clean)
                clean
                echo "done!"
                exit 0
                ;;
            clean-db)
                clean-db
                echo "done!"
                exit 0
                ;;
            refresh-env-file)
                refresh-env-file
                ;;
            *)
                echo "Invalid parameter was provided: $param"
                exit 1
                ;;
        esac
    done
}

function main() {

    if [ $# -eq 0 ] || [[ "$1" = "-v" && $# -eq 1 ]]
    then
        echo "No arguments supplied!"
        echo ""
        script_usage
    fi

    is_verbose=false

    docker compose convert > "$(pwd)/docker-compose-convert.yaml"
    parse_params "$@"
}

if [ ! -f "$(pwd)/env" ]; then
    echo "env missing, what did you do!"
    exit 1
fi

if [ ! -f "$(pwd)/.env" ]; then
    cp "$(pwd)/env" "$(pwd)/.env"
    chmod 777 "$(pwd)/.env" # change to write permission
fi

main "$@"
echo "done!"