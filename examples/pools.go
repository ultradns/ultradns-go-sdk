//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/dirpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/rdpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/slbpool"
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

	recordService, err := record.Get(client)

	if err != nil {
		log.Fatalf("Record service initialization failed with error - %v", err)
	}

	// creating RD pool
	rrSetKey, rrSet := newPoolRecordData("rd", "<zone name>", "A", "1.1.1.1")
	rrSet.Profile = &rdpool.Profile{Order: "FIXED"}
	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to create pool - '%v' : error - %v", rrSetKey, er)
	}

	// updating RD pool
	rrSet.RData = []string{"1.1.1.1", "2.2.2.2"}
	rrSet.Profile = &rdpool.Profile{Order: "ROUND_ROBIN"}
	if _, er := recordService.Update(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to update pool - '%v' : error - %v", rrSetKey, er)
	}

	// reading RD pool
	rrSetKey.PType = pool.RD
	if _, res, er := recordService.Read(rrSetKey); er != nil {
		log.Fatalf("Unable to read pool - '%v' : error - %v", rrSetKey, er)
	} else {
		for _, v := range res.RRSets {
			fmt.Println("Record Data: %+v\n", v)
		}
	}

	// deleting RD pool
	if _, er := recordService.Delete(rrSetKey); er != nil {
		log.Fatalf("Unable to delete pool - '%v' : error - %v", rrSetKey, er)
	}

	// creating SB pool
	rrSetKey, rrSet = newPoolRecordData("sb", "<zone name>", "AAAA", "0:0:0:0:0:0:0:1")
	rrSet.Profile = newSbPoolProfile()
	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to create pool - '%v' : error - %v", rrSetKey, er)
	}

	// creating SLB pool
	rrSetKey, rrSet = newPoolRecordData("slb", "<zone name>", "A", "1.1.1.1")
	rrSet.Profile = newSlbPoolProfile()
	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to create pool - '%v' : error - %v", rrSetKey, er)
	}

	// creating Dir pool
	rrSetKey, rrSet = newPoolRecordData("dir", "<zone name>", "A", "1.1.1.1")
	rrSet.Profile = &dirpool.Profile{
		RDataInfo: []*dirpool.RDataInfo{&dirpool.RDataInfo{AllNonConfigured: true}},
	}
	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to create pool - '%v' : error - %v", rrSetKey, er)
	}

}

func newPoolRecordData(ownerName, zoneName, rrType, rdata string) (*rrset.RRSetKey, *rrset.RRSet) {
	rrSetKey := &rrset.RRSetKey{
		Owner:      ownerName,
		Zone:       zoneName,
		RecordType: rrType,
	}
	rrSet := &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    rrType,
		RData:     []string{rdata},
	}
	return rrSetKey, rrSet
}

func newSbPoolProfile() *sbpool.Profile {
	rdataInfo := &pool.RDataInfo{
		State:         "NORMAL",
		RunProbes:     true,
		Priority:      1,
		FailoverDelay: 0,
		Threshold:     1,
	}
	return &sbpool.Profile{
		RDataInfo:        []*pool.RDataInfo{rdataInfo},
		RunProbes:        true,
		ActOnProbes:      true,
		FailureThreshold: 0,
		Order:            "FIXED",
		MaxActive:        1,
		MaxServed:        1,
	}
}

func newSlbPoolProfile() *slbpool.Profile {
	rdataInfo := &slbpool.RDataInfo{
		ProbingEnabled: false,
	}
	allFailRecord := &slbpool.AllFailRecord{
		RData:   "192.168.0.1",
		Serving: false,
	}
	monitor := &pool.Monitor{
		Method: "GET",
		URL:    "https://test-non-existing-zone.com/",
	}

	return &slbpool.Profile{
		RDataInfo:                []*slbpool.RDataInfo{rdataInfo},
		Monitor:                  monitor,
		AllFailRecord:            allFailRecord,
		RegionFailureSensitivity: "HIGH",
		ServingPreference:        "AUTO_SELECT",
		ResponseMethod:           "ROUND_ROBIN",
	}
}
