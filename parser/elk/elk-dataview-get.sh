#!/bin/bash
while getopts i: option
do
    case "${option}" in
        i) dvid=${OPTARG};;
    esac
done

if [ -z $dvid ]; then
  echo '-i for data view id required'
  exit
fi
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR
source ../.env

curl -k -X GET -u $ELASTIC_USER:$ELASTIC_PASS "$KIBANA_HOST/api/data_views/data_view/${dvid}" -H "kbn-xsrf: reporting"
