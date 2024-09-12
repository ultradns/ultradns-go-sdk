//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
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

	// Enabling the default trace logger
	client.EnableDefaultTraceLogger()

	zoneService, err := zone.Get(client)

	if err != nil {
		log.Fatalf("Zone service initialization failed with error - %v", err)
	}

	zoneName, accountName := "<primary zone name>", "<account name>"
	secondaryZone, namerserver := "<secondary zone name >", "<master nameserver>"

	// creating a primary zone
	zoneData := newPrimaryZoneData(zoneName, accountName)
	if _, err := zoneService.CreateZone(zoneData); err != nil {
		log.Fatalf("Unable to create primary zone - '%v' : error - %v", zoneName, err)
	}

	// creating a secondary zone
	zoneData = newSecondaryZoneData(secondaryZone, accountName, namerserver)
	if _, err := zoneService.CreateZone(zoneData); err != nil {
		log.Fatalf("Unable to create secondary zone - '%v' : error - %v", secondaryZone, err)
	}

	// Reading a primary zone
	if _, res, err := zoneService.ReadZone(zoneName); err != nil {
		log.Fatalf("Unable to get primary zone - '%v' : error - %v", zoneName, err)
	} else {
		fmt.Printf("Primary zone properties:\n '%+v'\n", res.Properties)
	}

	// Reading a secondary zone
	if _, res, err := zoneService.ReadZone(secondaryZone); err != nil {
		log.Fatalf("Unable to get primary zone - '%v' : error - %v", secondaryZone, err)
	} else {
		fmt.Printf("Secondary zone properties:\n '%+v'\n", res.Properties)
	}

	// Listing zones with limit
	query := &helper.QueryInfo{
		Limit: 2,
	}
	if _, res, err := zoneService.ListZone(query); err != nil {
		log.Fatalf("Unable to list zones : error - %v", err)
	} else {
		for _, v := range res.Zones {
			fmt.Println(v.Properties.Name)
		}
	}

	// Move zone from one account to another
	if _, er := zoneService.MigrateZoneAccount([]string{zoneName}, "<old account>", "<new account>"); er != nil {
		log.Fatalf("unable to move zone - %s : error - %s", zoneName, er.Error())
	}

	// Deleting a primary zone
	if _, err := zoneService.DeleteZone(zoneName); err != nil {
		log.Fatalf("Unable to delete primary zone - '%v' : error - %v", zoneName, err)
	}

	// Deleting a Secondary zone
	if _, err := zoneService.DeleteZone(secondaryZone); err != nil {
		log.Fatalf("Unable to delete secondary zone - '%v' : error - %v", secondaryZone, err)
	}

	fmt.Println("Completed")
}

func newPrimaryZoneData(zoneName, accountName string) *zone.Zone {
	zoneProp := &zone.Properties{
		Name:        zoneName,
		AccountName: accountName,
		Type:        "PRIMARY",
	}

	primaryZone := &zone.PrimaryZone{
		CreateType: "NEW",
	}

	return &zone.Zone{
		Properties:        zoneProp,
		PrimaryCreateInfo: primaryZone,
	}
}

func newSecondaryZoneData(zoneName, accountName, nameserver string) *zone.Zone {
	nameServerIP := &zone.NameServer{
		IP: nameserver,
	}
	nameServerIPList := &zone.NameServerIPList{
		NameServerIP1: nameServerIP,
	}

	primaryNameServer := &zone.PrimaryNameServers{
		NameServerIPList: nameServerIPList,
	}

	secondaryZone := &zone.SecondaryZone{
		PrimaryNameServers: primaryNameServer,
	}

	zoneProp := &zone.Properties{
		Name:        zoneName,
		AccountName: accountName,
		Type:        "SECONDARY",
	}

	return &zone.Zone{
		Properties:          zoneProp,
		SecondaryCreateInfo: secondaryZone,
	}
}
