//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/dns"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/http"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/ping"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/tcp"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func main() {
	conf := client.Config{
		Username: "<username>",
		Password: "<password>",
		HostURL:  "https://api.ultradns.com/",
	}

	client, err := client.NewClient(conf)

	if err != nil {
		log.Fatalf("Client initialization failed with error - %v", err)
	}

	probeService, err := probe.Get(client)

	if err != nil {
		log.Fatalf("Probe service initialization failed with error - %v", err)
	}

	// creating DNS Probe
	rrSetKey := newRRSetKey("sb", "<zone name>", "AAAA", probe.DNS)
	if _, er := probeService.Create(rrSetKey, NewDNSProbeData()); er != nil {
		log.Fatalf("Unable to create probe - '%v' : error - %v", rrSetKey, er)
	}

	// creating TCP Probe
	rrSetKey = newRRSetKey("sb", "<zone name>", "A", probe.TCP)
	if _, er := probeService.Create(rrSetKey, NewTCPProbeData()); er != nil {
		log.Fatalf("Unable to create probe - '%v' : error - %v", rrSetKey, er)
	}

	// creating HTTP Probe
	rrSetKey = newRRSetKey("sb", "<zone name>", "AAAA", probe.HTTP)
	if _, er := probeService.Create(rrSetKey, NewHTTPProbeData()); er != nil {
		log.Fatalf("Unable to create probe - '%v' : error - %v", rrSetKey, er)
	}

	// creating PING Probe
	rrSetKey = newRRSetKey("sb", "<zone name>", "A", probe.PING)
	if _, er := probeService.Create(rrSetKey, NewPINGProbeData()); er != nil {
		log.Fatalf("Unable to create probe - '%v' : error - %v", rrSetKey, er)
	}

	// Listing Probes
	if _, res, er := probeService.List(rrSetKey, &probe.Query{}); er != nil {
		log.Fatalf("Unable to list probe - '%v' : error - %v", rrSetKey, er)
	} else {
		fmt.Println("Probe Data: %+v \n", res.Probes[0])
	}

	// Read a Probe which requires probe id
	rrSetKey.ID = "06084E00F38ADF64"
	if _, res, er := probeService.Read(rrSetKey); er != nil {
		log.Fatalf("Unable to read probe - '%v' : error - %v", rrSetKey, er)
	} else {
		fmt.Println("Probe Data: %+v\n", res)
	}

	// Update Probe
	probeData := NewTCPProbeData()
	probeData.Details.(*tcp.Details).Port = 54
	if _, er := probeService.Update(rrSetKey, probeData); er != nil {
		log.Fatalf("Unable to update probe - '%v' : error - %v", rrSetKey, er)
	}

	// Delete probe
	if _, er := probeService.Delete(rrSetKey); er != nil {
		log.Fatalf("Unable to delete probe - '%v' : error - %v", rrSetKey, er)
	}
}

func newRRSetKey(ownerName, zoneName, recordType, pType string) *rrset.RRSetKey {
	return &rrset.RRSetKey{
		Owner:      ownerName,
		Zone:       zoneName,
		RecordType: recordType,
		PType:      pType,
	}
}

func NewDNSProbeData() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		Run: limit,
	}
	details := &dns.Details{
		Port:   53,
		Limits: limitInfo,
	}
	return &probe.Probe{
		Type:      probe.DNS,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}
}

func NewTCPProbeData() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		Connect: limit,
	}
	details := &tcp.Details{
		Port:   53,
		Limits: limitInfo,
	}
	return &probe.Probe{
		Type:      probe.TCP,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}
}

func NewHTTPProbeData() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		Run:     limit,
		Connect: limit,
	}
	transction := &http.Transaction{
		Method:          "GET",
		ProtocolVersion: "HTTP/1.0",
		URL:             "https://test-unknown-host.com/",
		Limits:          limitInfo,
	}
	details := &http.Details{
		Transactions: []*http.Transaction{transction},
	}
	return &probe.Probe{
		Type:      probe.HTTP,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}
}

func NewPINGProbeData() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		LossPercent: limit,
		Total:       limit,
	}
	details := &ping.Details{
		Packets:    3,
		PacketSize: 56,
		Limits:     limitInfo,
	}
	return &probe.Probe{
		Type:      probe.PING,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}
}
