#!/bin/bash

CURL='/usr/bin/curl'
CURLARGS="-f -s -S -k -X GET"

examples() {
    # Array of prefixes as defined by code challenge
    declare -a Prefixes=("th" "fr" "pi" "sh" "wu" "ar" "il" "ne" "se" "pl")
 
    for val in ${Prefixes[@]}; do
        echo "Prefix:" $val
        $CURL $CURLARGS http://localhost:$1/autocomplete\?prefix\=$val
        echo ""
    done
}

examples $1