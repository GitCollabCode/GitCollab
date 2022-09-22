#!/usr/bin/env bash

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
    docker compose build
}

function start() {
    if [ "$is_verbose" = "true" ]; then
        docker compose up
    else
        docker compose up -d 
    fi
}

function restart() {
    docker compose down
    docker compose up -d 
}

function stop() {
    docker compose stop
}

function clean() {
    docker compose down
    docker system prune -a
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
                exit 0
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

    parse_params "$@"
}

main "$@"
