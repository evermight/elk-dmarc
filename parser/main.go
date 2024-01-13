package main

import (
	"dmarc-analyzer/common"
	"dmarc-analyzer/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	i := 0
	for _, s := range common.FindFiles("./logs/ingest", ".xml") {
		report := dto.UnmarshalDMARC(s)
		entries := dto.FlattenReport(report)

		for _, entry := range entries {

			jsonContent, _ := json.Marshal(entry)
			_ = ioutil.WriteFile("./output/"+fmt.Sprintf("%v", i)+".json", jsonContent, 0755)
			i = i + 1
		}
	}
}
