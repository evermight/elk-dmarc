SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR;
./elk-pipeline-import.sh
./elk-record-import.sh
./elk-dataview-import.sh
./elk-saved-objects-import.sh
