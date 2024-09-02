package main

import (
	"fmt"
	"log"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
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

	ipService, err := ip.Get(client)

	if err != nil {
		log.Fatalf("IP service initialization failed with error - %v", err)
	}

	ipdata := &ip.DirGroupIP{
		Name:        "<group name>",
		AccountName: "<account name>",
		IPs: []*ip.IPAddress{
			&ip.IPAddress{
				Start: "192.168.1.1",
				End:   "192.168.1.10",
			},
			&ip.IPAddress{
				Cidr: "192.168.2.0/24",
			},
			&ip.IPAddress{
				Address: "192.168.3.4",
			},
		},
		Description: "Description of IP",
	}

	// create ip group
	if _, er := ipService.Create(ipdata); er != nil {
		log.Fatalf("Unable to create ip group - '%v' : error - %v", ipdata.DirGroupIPID, er)
	}

	// Read ip group
	if _, res, _, er := ipService.Read(ipdata.DirGroupIPID()); er != nil {
		log.Fatalf("Unable to read ip group - '%v' : error - %v", ipdata.DirGroupIPID, er)
	} else {
		fmt.Println(res)
	}

	// Update ip group
	ipdata.IPs = []*ip.IPAddress{
		&ip.IPAddress{
			Address: "1.1.1.1",
		},
	}

	if _, er := ipService.Update(ipdata); er != nil {
		log.Fatalf("Unable to update ip group - '%v' : error - %v", ipdata.DirGroupIPID, er)
	}

	// Delete ip group
	if _, er := ipService.Delete(ipdata.DirGroupIPID()); er != nil {
		log.Fatalf("Unable to update ip group - '%v' : error - %v", ipdata.DirGroupIPID, er)
	}
}
