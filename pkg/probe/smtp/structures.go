package smtp

import "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	Port   int                `json:"port,omitempty"`
	Limits *helper.LimitsInfo `json:"limits,omitempty"`
}
