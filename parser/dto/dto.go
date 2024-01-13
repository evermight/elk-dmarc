package dto

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type FlatRecord struct {
	Version                string
	ReportOrgName          string
	ReportEmail            string
	ReportExtraContactInfo string
	ReportId               string
	ReportDateBegin        int
	ReportDateEnd          int
	PolicyPublishedDomain  string
	PolicyPublishedADkim   string
	PolicyPublishedASpf    string
	PolicyPublishedP       string
	PolicyPublishedSP      string
	PolicyPublishedNP      string
	PolicyPublishedPCT     string
	PolicyPublishedFO      string
	IdentifiersHeaderFrom  string
	RowSourceIp            string
	RowCount               int
	RowPEDisposition       string
	RowPEDkim              string
	RowPESpf               string
	AuthResultSpfDomain    string `xml:"auth_results>spf>domain"`
	AuthResultSpfResult    string `xml:"auth_results>spf>result"`
	AuthResultSpfScope     string `xml:"auth_results>spf>scope"`
	AuthResultSpfSelector  string `xml:"auth_results>spf>selector"`
	AuthResultDkimDomain   string `xml:"auth_results>dkim>domain"`
	AuthResultDkimResult   string `xml:"auth_results>dkim>result"`
	AuthResultDkimScope    string `xml:"auth_results>dkim>scope"`
	AuthResultDkimSelector string `xml:"auth_results>dkim>selector"`
}
type Report struct {
	Version                string   `xml:"version"`
	ReportOrgName          string   `xml:"report_metadata>org_name"`
	ReportEmail            string   `xml:"report_metadata>email"`
	ReportExtraContactInfo string   `xml:"report_metadata>extra_contact_info"`
	ReportId               string   `xml:"report_metadata>report_id"`
	ReportDateBegin        int      `xml:"report_metadata>date_range>begin"`
	ReportDateEnd          int      `xml:"report_metadata>date_range>end"`
	PolicyPublishedDomain  string   `xml:"policy_published>domain"`
	PolicyPublishedADkim   string   `xml:"policy_published>adkim"`
	PolicyPublishedASpf    string   `xml:"policy_published>aspf"`
	PolicyPublishedP       string   `xml:"policy_published>p"`
	PolicyPublishedSP      string   `xml:"policy_published>sp"`
	PolicyPublishedNP      string   `xml:"policy_published>np"`
	PolicyPublishedPCT     string   `xml:"policy_published>pct"`
	PolicyPublishedFO      string   `xml:"policy_published>fo"`
	Records                []Record `xml:"record"`
}

type Record struct {
	IdentifiersHeaderFrom  string `xml:"identifiers>header_from"`
	RowSourceIp            string `xml:"row>source_ip"`
	RowCount               int    `xml:"row>count"`
	RowPEDisposition       string `xml:"row>policy_evaluated>disposition"`
	RowPEDkim              string `xml:"row>policy_evaluated>dkim"`
	RowPESpf               string `xml:"row>policy_evaluated>spf"`
	AuthResultSpfDomain    string `xml:"auth_results>spf>domain"`
	AuthResultSpfResult    string `xml:"auth_results>spf>result"`
	AuthResultSpfScope     string `xml:"auth_results>spf>scope"`
	AuthResultSpfSelector  string `xml:"auth_results>spf>selector"`
	AuthResultDkimDomain   string `xml:"auth_results>dkim>domain"`
	AuthResultDkimResult   string `xml:"auth_results>dkim>result"`
	AuthResultDkimScope    string `xml:"auth_results>dkim>scope"`
	AuthResultDkimSelector string `xml:"auth_results>dkim>selector"`
}

func FlattenReport(report Report) []FlatRecord {
	flattenedRecords := []FlatRecord{}
	for i := range report.Records {
		flat := FlatRecord{
			Version:                report.Version,
			ReportOrgName:          report.ReportOrgName,
			ReportEmail:            report.ReportEmail,
			ReportExtraContactInfo: report.ReportExtraContactInfo,
			ReportId:               report.ReportId,
			ReportDateBegin:        report.ReportDateBegin,
			ReportDateEnd:          report.ReportDateEnd,
			PolicyPublishedDomain:  report.PolicyPublishedDomain,
			PolicyPublishedADkim:   report.PolicyPublishedADkim,
			PolicyPublishedASpf:    report.PolicyPublishedASpf,
			PolicyPublishedP:       report.PolicyPublishedP,
			PolicyPublishedSP:      report.PolicyPublishedSP,
			PolicyPublishedNP:      report.PolicyPublishedNP,
			PolicyPublishedPCT:     report.PolicyPublishedPCT,
			PolicyPublishedFO:      report.PolicyPublishedFO,
			IdentifiersHeaderFrom:  report.Records[i].IdentifiersHeaderFrom,
			RowSourceIp:            report.Records[i].RowSourceIp,
			RowCount:               report.Records[i].RowCount,
			RowPEDisposition:       report.Records[i].RowPEDisposition,
			RowPEDkim:              report.Records[i].RowPEDkim,
			RowPESpf:               report.Records[i].RowPESpf,
			AuthResultSpfDomain:    report.Records[i].AuthResultSpfDomain,
			AuthResultSpfResult:    report.Records[i].AuthResultSpfResult,
			AuthResultSpfScope:     report.Records[i].AuthResultSpfScope,
			AuthResultSpfSelector:  report.Records[i].AuthResultSpfSelector,
			AuthResultDkimDomain:   report.Records[i].AuthResultDkimDomain,
			AuthResultDkimResult:   report.Records[i].AuthResultDkimResult,
			AuthResultDkimScope:    report.Records[i].AuthResultDkimScope,
			AuthResultDkimSelector: report.Records[i].AuthResultDkimSelector,
		}
		flattenedRecords = append(flattenedRecords, flat)
	}
	return flattenedRecords
}
func UnmarshalDMARC(dmarcFileName string) Report {
	var report Report
	file, err := os.Open(dmarcFileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return report
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return report
	}

	err = xml.Unmarshal(content, &report)
	if err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return report
	}
	return report
}
