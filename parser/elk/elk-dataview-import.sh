#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR;
source ../.env

curl -k -X POST -u $ELASTIC_USER:$ELASTIC_PASS "$KIBANA_HOST/api/data_views/data_view" -H "kbn-xsrf: reporting" -H "Content-Type: application/json" \
-d @./dataview/record.json
