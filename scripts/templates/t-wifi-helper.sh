#!/bin/sh

# this is the helper script template for lora management
# modify function body [start, stop, restart, status]
# !!!!!!!!!!!!!!!! ATTENTION !!!!!!!!!!!!!!!
# Do NOT Modify the function name unless you understand what you are doing

set -e

start() {
    echo "did start $@"
}

stop() {
    echo "did stop $@"
}

restart() {
    echo "did restart $@"
}

# output should be `key=value` pairs
# key or value including spaces should be quoted with `"`
# such as
# "\"some spaced key here\"=\"some spaced value here\""
# !!!!!!!!!!!!!!!!!!! NOTICE !!!!!!!!!!!!!!!!!!!
# Any output not conforming the format will cause an error
# 
# VALID Fields:
# `on`      bool        mark device on or off
# `ssids`   string      scaned ssid list

status() {
    echo on=true
    echo "ssid_list=[test_ssid_1, test_ssid_2]"
}

"$@"
