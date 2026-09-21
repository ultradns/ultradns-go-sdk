//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/webforward"
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

	webForwardService, err := webforward.Get(client)

	if err != nil {
		log.Fatalf("Web forward service initialization failed with error - %v", err)
	}

	zoneName := "<zone name>"

	// creating a web forward
	webForwardData := &webforward.WebForward{
		RequestTo:           "www.example.com",
		DefaultRedirectTo:   "https://example.com/",
		DefaultForwardType:  webforward.HTTP301Redirect,
		RelativeForwardType: webforward.ParameterAndPath,
	}
	if _, res, err := webForwardService.Create(zoneName, webForwardData); err != nil {
		log.Fatalf("Unable to create web forward - '%v' : error - %v", zoneName, err)
	} else {
		fmt.Printf("Created web forward:\n '%+v'\n", res)
	}

	// listing web forwards for a zone
	if _, res, err := webForwardService.List(zoneName, nil); err != nil {
		log.Fatalf("Unable to list web forwards - '%v' : error - %v", zoneName, err)
	} else {
		for _, v := range res.WebForwards {
			fmt.Println(v.RequestTo)
		}
	}

	// To create an HTTPS forward, use an https:// requestTo and either an
	// uploaded certificate ID or webforward.EECertificateManagedType:
	//
	// httpsWebForwardData := &webforward.WebForward{
	// 	RequestTo:          "https://secure.example.com",
	// 	DefaultRedirectTo:  "https://example.com/",
	// 	DefaultForwardType: webforward.HTTP301Redirect,
	// 	CertificateID:      "<certificate GUID>",
	// }

	fmt.Println("Completed")
}
