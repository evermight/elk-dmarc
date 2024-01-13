#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR;

cd logs/zipped;
find . -name "*.gz" | while read filename; do gunzip $filename; done;
find . -name "*.zip" | while read filename; do unzip -o $filename; done;
mv *.xml $SCRIPT_DIR/logs/ingest/;

cd $SCRIPT_DIR;
go run main.go;
./elk/elk-start.sh

# Clean Up - delete data so it doesn't accidentally get imported again
rm -rf $SCRIPT_DIR/output/*.json;
rm -rf $SCRIPT_DIR/logs/ingest/*.xml;
rm -rf $SCRIPT_DIR/logs/zipped/*.zip;
