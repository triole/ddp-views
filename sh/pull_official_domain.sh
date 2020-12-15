#!/bin/bash

scriptdir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
basedir=$(echo "${scriptdir}" | grep -Po ".*(?=/)")
outfile="${basedir}/tmp/domain.xml"

curl -s \
    https://raw.githubusercontent.com/rdmorganiser/rdmo-catalog/master/rdmorganiser/domain/rdmo.xml \
    -o "${outfile}"
