package rrset

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

var rrTypes = map[string]string{
	"A":     "A (1)",
	"1":     "A (1)",
	"AAAA":  "AAAA (28)",
	"28":    "AAAA (28)",
	"CNAME": "CNAME (5)",
	"5":     "CNAME (5)",
	"MX":    "MX (15)",
	"15":    "MX (15)",
	"SRV":   "SRV (33)",
	"33":    "SRV (33)",
	"TXT":   "TXT (16)",
	"16":    "TXT (16)",
	"PTR":   "PTR (12)",
	"12":    "PTR (12)",
}

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
	r.Type = rrTypes[r.Type]

	if !strings.Contains(r.Name, r.Zone) {
		r.Name += "." + r.Zone
	}

	return fmt.Sprintf("%s:%s:%s", r.Name, r.Zone, r.Type)
}
