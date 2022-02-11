package ping

import "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	Packets    int                `json:"packets,omitempty"`
	PacketSize int                `json:"packetSize,omitempty"`
	Limits     *helper.LimitsInfo `json:"limits,omitempty"`
}
