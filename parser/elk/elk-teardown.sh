#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR;
source ../.env

curl -k -X DELETE -u $ELASTIC_USER:$ELASTIC_PASS "$ELASTIC_HOST/dmarc-record"

curl -k -X DELETE -u $ELASTIC_USER:$ELASTIC_PASS "$KIBANA_HOST/api/data_views/data_view/f4d7002a-ca4a-4f78-800d-04c64b3ab0e9" -H "kbn-xsrf: reporting"
curl -k -X DELETE -u $ELASTIC_USER:$ELASTIC_PASS "$KIBANA_HOST/api/saved_objects/dashboard/dad28d10-af6a-11ee-afa7-0de1eaf80728" -H "kbn-xsrf: reporting"

curl -k -X DELETE -u $ELASTIC_USER:$ELASTIC_PASS "$ELASTIC_HOST/_ingest/pipeline/geoip"
