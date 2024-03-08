#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR
source ../.env

# curl -k -X PUT -u $ELASTIC_USER:$ELASTIC_PASS "$ELASTIC_HOST/dmarc-record"
# curl -k -X PUT -u $ELASTIC_USER:$ELASTIC_PASS "$ELASTIC_HOST/dmarc-record/_mapping" \
# -H "Content-Type: application/json" \
# -d @./mapping/record.json

payload=""
i=0
for file in ../output/*.json; do
  if [ -f "$file" ]; then

     payload="$payload"$'\n''{"index":{}}'$'\n'$(cat $file)
     i=$((i+1))
     if [[ "$i" -eq "$ELASTIC_BULK_LIMIT" ]]; then
       i=0
       curl -k -X POST -u $ELASTIC_USER:$ELASTIC_PASS "$ELASTIC_HOST/dmarc-record/_bulk?pipeline=geoip" \
         -H "Content-Type: application/json" \
         --data-binary @- <<< "$payload"
       payload=""
     fi
  fi
done
if [[ "$i" -ne 0 ]]; then
  curl -k -X POST -u $ELASTIC_USER:$ELASTIC_PASS "$ELASTIC_HOST/dmarc-record/_bulk?pipeline=geoip" \
    -H "Content-Type: application/json" \
    --data-binary @- <<< "$payload"
fi
