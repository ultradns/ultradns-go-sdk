//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
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

	geoService, err := geo.Get(client)

	if err != nil {
		log.Fatalf("Geo service initialization failed with error - %v", err)
	}

	geodata := &geo.DirGroupGeo{
		Name:        "<group name>",
		AccountName: "<account name>",
		Codes:       []string{"CA", "US", "MX"},
		Description: "Description of GEO",
	}

	// create geo group
	if _, er := geoService.Create(geodata); er != nil {
		log.Fatalf("Unable to create geo group - '%v' : error - %v", geodata.DirGroupGeoID, er)
	}

	// Read geo group
	if _, res, _, er := geoService.Read(geodata.DirGroupGeoID()); er != nil {
		log.Fatalf("Unable to read geo group - '%v' : error - %v", geodata.DirGroupGeoID, er)
	} else {
		fmt.Println(res)
	}

	// Update geo group
	geodata.Codes = []string{"CA", "MX"}
	if _, er := geoService.Update(geodata); er != nil {
		log.Fatalf("Unable to update geo group - '%v' : error - %v", geodata.DirGroupGeoID, er)
	}

	// Delete geo group
	if _, er := geoService.Delete(geodata.DirGroupGeoID()); er != nil {
		log.Fatalf("Unable to update geo group - '%v' : error - %v", geodata.DirGroupGeoID, er)
	}

}
