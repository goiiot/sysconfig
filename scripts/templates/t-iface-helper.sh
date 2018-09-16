#!/bin/sh

# this is the helper script template for network interface management
# modify function body [start, stop, restart, status]
# !!!!!!!!!!!!!!!!!!!!!!!!!!!! ATTENTION !!!!!!!!!!!!!!!!!!!!!!!!!!!!
# Do NOT Modify the function name unless you understand what you are doing

set -e

start() {
    echo "did start $@"
}

stop() {
    echo "did stop $@"
}

restart() {
    stop
    sleep 2
    start
    echo "did restart $@"
}

# output should be `key=value` pairs
# key or value including spaces should be quoted with `"`
# such as
# "\"some spaced key here\"=\"some spaced value here\""
# !!!!!!!!!!!!!!!!!!!!!!!!!!!! NOTICE !!!!!!!!!!!!!!!!!!!!!!!!!!!!
# Any output not conforming the format will cause an error
status() {
    echo "name=$@"
    t=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
    echo time=$t
    echo "\"some spaced key here\"=\"some spaced value here\""
    echo num=12345
}

# input params will be `key=value` pairs
# export these params and you can get these values by key
# e.g. 
# $ ./t-iface-helper.sh test=a data=1
# Output: a 1
# Available
config() {
    export $@
    echo $test $data
}

"$@"
