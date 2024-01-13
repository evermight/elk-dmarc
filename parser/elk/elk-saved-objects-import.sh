#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR;

source ../.env

curl -k -X POST -u $ELASTIC_USER:$ELASTIC_PASS "$KIBANA_HOST/api/saved_objects/_import" -H "kbn-xsrf: true" \
--form file=@./dashboard/dashboard.ndjson
