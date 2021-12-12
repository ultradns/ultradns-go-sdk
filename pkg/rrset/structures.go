package rrset

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

type RRSet struct {
	OwnerName string      `json:"ownerName,omitempty"`
	RRType    string      `json:"rrtype,omitempty"`
	TTL       int         `json:"ttl,omitempty"`
	RData     []string    `json:"rdata,omitempty"`
	Profile   interface{} `json:"profile,omitempty"`
}

type RRSetKey struct {
	Zone string
	Type string
	Name string
}

type ResponseList struct {
	ZoneName   string             `json:"zoneName,omitempty"`
	QueryInfo  *helper.QueryInfo  `json:"queryInfo,omitempty"`
	ResultInfo *helper.ResultInfo `json:"resultInfo,omitempty"`
	RRSets     []*RRSet           `json:"rrSets,omitempty"`
}

func (r RRSetKey) URI() string {
	r.Name = url.PathEscape(r.Name)
	r.Zone = url.PathEscape(r.Zone)

	if r.Type == "" {
		r.Type = "ANY"
	}

	return fmt.Sprintf("zones/%s/rrsets/%s/%s", r.Zone, r.Type, r.Name)
}

func (r RRSetKey) ID() string {
	r.Type = GetRRTypeFullString(r.Type)

	if !strings.Contains(r.Name, r.Zone) {
		r.Name += "." + r.Zone
	}

	return fmt.Sprintf("%s:%s:%s", r.Name, r.Zone, r.Type)
}

func GetRRTypeFullString(key string) string {
	var rrTypes = map[string]string{
		"A":         "A (1)",
		"1":         "A (1)",
		"NS":        "NS (2)",
		"2":         "NS (2)",
		"CNAME":     "CNAME (5)",
		"5":         "CNAME (5)",
		"SOA":       "SOA (6)",
		"6":         "SOA (6)",
		"PTR":       "PTR (12)",
		"12":        "PTR (12)",
		"HINFO":     "HINFO (13)",
		"13":        "HINFO (13)",
		"MX":        "MX (15)",
		"15":        "MX (15)",
		"TXT":       "TXT (16)",
		"16":        "TXT (16)",
		"RP":        "RP (17)",
		"17":        "RP (17)",
		"AAAA":      "AAAA (28)",
		"28":        "AAAA (28)",
		"SRV":       "SRV (33)",
		"33":        "SRV (33)",
		"NAPTR":     "NAPTR (35)",
		"35":        "NAPTR (35)",
		"DS":        "DS (43)",
		"43":        "DS (43)",
		"SSHFP":     "SSHFP (44)",
		"44":        "SSHFP (44)",
		"TLSA":      "TLSA (52)",
		"52":        "TLSA (52)",
		"SPF":       "SPF (99)",
		"99":        "SPF (99)",
		"CAA":       "CAA (257)",
		"257":       "CAA (257)",
		"APEXALIAS": "APEXALIAS (65282)",
		"65282":     "APEXALIAS (65282)",
	}

	return rrTypes[key]
}
