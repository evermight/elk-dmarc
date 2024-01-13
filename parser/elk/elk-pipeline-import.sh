#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR
source ../.env

curl -k -X PUT -u $ELASTIC_USER:$ELASTIC_PASS "$ELASTIC_HOST/_ingest/pipeline/geoip" \
-H "Content-Type: application/json" \
-d @./pipeline/geoip.json
