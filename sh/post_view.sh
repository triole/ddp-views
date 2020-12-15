#!/bin/bash

scriptdir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
basedir=$(echo "${scriptdir}" | grep -Po ".*(?=/)")
conf="${basedir}/conf/conf.toml"

curl=$(stoml "${conf}" "curl")
burl=$(stoml "${conf}" "base_url")

id="${1}"
fil="${2}"

if [[ ! ${id} =~ ^[0-9]+$ ]]; then
    echo "First arg needs to be an integer: post_view.sh 20 tmp/20.json"
    exit
fi

cmd="$(echo ${curl} --data-binary \"@${fil}\" ${burl}${id}/ | sd "<METHOD>" "PUT")"

echo "<yellow>${cmd}</yellow>" | tml
eval ${cmd} || (eval ${cmd} | html2data '.exception_value')
echo ""
