package main

import (
	"fmt"
	"log"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
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

	// Enabling the default debug logger
	client.EnableDefaultDebugLogger()

	recordService, err := record.Get(client)

	if err != nil {
		log.Fatalf("Record service initialization failed with error - %v", err)
	}

	// creating DNS record of type A
	rrSetKey, rrSet := newRecordData("www", "<zone name>", "A", "1.1.1.1")
	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to create record - '%v' : error - %v", rrSetKey, er)
	}

	// updating DNS record of type A
	rrSet.RData = []string{"2.2.2.2"}
	if _, er := recordService.Update(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to update record - '%v' : error - %v", rrSetKey, er)
	}

	// reading DNS record of type A
	if _, res, er := recordService.Read(rrSetKey); er != nil {
		log.Fatalf("Unable to read record - '%v' : error - %v", rrSetKey, er)
	} else {
		fmt.Printf("Record Data: \n%+v\n", res.RRSets[0])
	}

	// deleting DNS record of type A
	if _, er := recordService.Delete(rrSetKey); er != nil {
		log.Fatalf("Unable to delete record - '%v' : error - %v", rrSetKey, er)
	}

	// creating DNS record of type CAA
	rrSetKey, rrSet = newRecordData("<zone name>", "<zone name>", "CAA", "0 issue ultradns")
	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to create record - '%v' : error - %v", rrSetKey, er)
	}

	// creating DNS record of type DS
	rrSetKey, rrSet = newRecordData("ds", "<zone name>", "DS", "25286 1 1 340437DC66C3DFAD0B3E849740D2CF1A4151671D")
	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to create record - '%v' : error - %v", rrSetKey, er)
	}

	// creating DNS record of type HTTPS
	rrSetKey, rrSet = newRecordData("https", "<zone name>", "HTTPS", "1 www.test.com. alpn=h2")
	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		log.Fatalf("Unable to create record - '%v' : error - %v", rrSetKey, er)
	}

}

func newRecordData(ownerName, zoneName, rrType, rdata string) (*rrset.RRSetKey, *rrset.RRSet) {
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
