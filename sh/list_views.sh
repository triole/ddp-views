#!/bin/bash

scriptdir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
basedir=$(echo "${scriptdir}" | grep -Po ".*(?=/)")
conf="${basedir}/conf/conf.toml"

curl=$(stoml "${conf}" "curl")
burl=$(stoml "${conf}" "base_url")

cmd="$(echo ${curl} ${burl} | sd "<METHOD>" "GET")"

echo "<yellow>${cmd}</yellow>" | tml
eval ${cmd} | jq ".[] | {title_en, id}"
