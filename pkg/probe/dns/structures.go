package dns

import "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	OwnerName string             `json:"ownerName,omitempty"`
	Type      string             `json:"type,omitempty"`
	Port      int                `json:"port,omitempty"`
	TCPOnly   bool               `json:"tcpOnly,omitempty"`
	Limits    *helper.LimitsInfo `json:"limits,omitempty"`
}
