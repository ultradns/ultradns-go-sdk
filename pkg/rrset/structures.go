package rrset

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

type RRSet struct {
	OwnerName string     `json:"ownerName,omitempty"`
	RRType    string     `json:"rrtype,omitempty"`
	TTL       int        `json:"ttl,omitempty"`
	RData     []string   `json:"rdata,omitempty"`
	Profile   RawProfile `json:"profile,omitempty"`
}

type RRSetKey struct {
	Zone string
	Type string
	Name string
}

type RawProfile interface {
	SetContext()
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
	r.Name = helper.GetOwnerFQDN(r.Name, r.Zone)
	r.Zone = helper.GetZoneFQDN(r.Zone)
	r.Type = helper.GetRecordTypeFullString(r.Type)

	return fmt.Sprintf("%s:%s:%s", r.Name, r.Zone, r.Type)
}

func GetRRSetKeyFromID(id string) *RRSetKey {
	rrSetKeyData := &RRSetKey{}
	splitStringData := strings.Split(id, ":")

	if len(splitStringData) == 3 {
		rrSetKeyData.Name = splitStringData[0]
		rrSetKeyData.Zone = splitStringData[1]
		rrSetKeyData.Type = helper.GetRecordTypeString(splitStringData[2])
	}

	return rrSetKeyData
}
