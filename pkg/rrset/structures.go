package rrset

import (
	"fmt"
	"net/url"
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

func (r RRSetKey) URI() string {
	r.Name = url.PathEscape(r.Name)
	r.Zone = url.PathEscape(r.Zone)

	return fmt.Sprintf("zones/%s/rrsets/%s/%s", r.Zone, r.Type, r.Name)
}

func (r RRSetKey) ID() string {
	return fmt.Sprintf("%s:%s:%s", r.Name, r.Zone, r.Type)
}
