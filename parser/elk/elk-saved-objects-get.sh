#!/bin/bash
while getopts t:i: option
do
    case "${option}" in
        i) dvid=${OPTARG};;
        t) stype=${OPTARG};;
    esac
done

if [ -z $dvid ]; then
  echo '-i for dashboard id required'
  exit
fi

if [ -z $stype ]; then
  echo '-t for object type required: visualization, dashboard, search, index-pattern'
  exit
fi

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR;
source ../.env

curl -k -X POST -u $ELASTIC_USER:$ELASTIC_PASS "$KIBANA_HOST/api/saved_objects/_export" -H "kbn-xsrf: reporting" -H "Content-Type: application/json" \
-d '{
  "objects": [
    {
      "type": "'$stype'",
      "id": "'$dvid'"
    }
  ]
}'
