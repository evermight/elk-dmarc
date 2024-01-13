#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR;
source ../.env

curl -k -X GET -u $ELASTIC_USER:$ELASTIC_PASS "$KIBANA_HOST/api/data_views" -H "kbn-xsrf: reporting"
