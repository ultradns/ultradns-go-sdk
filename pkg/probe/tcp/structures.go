package tcp

import "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	ControlIP string             `json:"controlIP,omitempty"`
	Port      int                `json:"port,omitempty"`
	Limits    *helper.LimitsInfo `json:"limits,omitempty"`
}
