#!/bin/bash
#
# This is the InfluxDB test script.
# This script can run tests in different environments.
#
# Usage: ./test.sh <environment_index>
# Corresponding environments for environment_index:
#      0: normal 64bit tests
#      1: race enabled 64bit tests
#      2: normal 32bit tests
#      3: tsi build
#      count: print the number of test environments
#      *: to run all tests in parallel containers
#
# Logs from the test runs will be saved in OUTPUT_DIR, which defaults to ./test-logs
#

# Get dir of script and make it is our working directory.
DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
cd $DIR

# Set the default OUTPUT_DIR
OUTPUT_DIR=${OUTPUT_DIR-./build-logs}
# Set default parallelism
PARALLELISM=${PARALLELISM-1}
# Set default timeout
TIMEOUT=${TIMEOUT-1500s}

# Default to deleteing the container
DOCKER_RM=${DOCKER_RM-true}

# Update this value if you add a new test environment.
ENV_COUNT=3

# Default return code 0
rc=0

# Executes the given statement, and exits if the command returns a non-zero code.
function exit_if_fail {
    command=$@
    echo "Executing '$command'"
    $command
    rc=$?
    if [ $rc -ne 0 ]; then
        echo "'$command' returned $rc."
        exit $rc
    fi
}

# Convert dockerfile name to valid docker image tag name.
function filename2imagename {
    echo ${1/Dockerfile/influxdb}
}


# Build the docker image defined by given dockerfile.
function build_docker_image {
    local dockerfile=$1
    local imagename=$2

    echo "Building docker image $imagename"
    exit_if_fail docker build --rm=$DOCKER_RM -f "$dockerfile" -t "$imagename" .
}


if [ ! -d "$OUTPUT_DIR" ]
then
    mkdir -p "$OUTPUT_DIR"
fi


local dockerfile=$1
local imagename=$(filename2imagename "$dockerfile")
shift
local name=$1
shift
local logfile="$OUTPUT_DIR/${name}.log"

>&2 echo 'Build docker image'
build_docker_image "$dockerfile" "$imagename"
rc=$?

exit $rc
