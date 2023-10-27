# UltraDNS SDK for Go

ultradns-go-sdk is the official UltraDNS SDK for the Go programming language.<br/>
Golang-Version : `1.19`

Jump To:
* [Getting Started](#Getting-Started)
* [Quick Examples](#Quick-Examples)

## Getting Started

### Installing

Use `go get` to retrieve the latest version of SDK to add it to your `GOPATH` workspace.

	go get github.com/ultradns/ultradns-go-sdk@latest

## Quick Examples

### Complete SDK Example

This example shows a complete working Go file which will create a primary zone in UltraDNS. 
This example highlights how to get services using client and make requests.

```go
package main

import (
	"fmt"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func main() {
	conf := client.Config{
		Username: "username",
		Password: "password",
		HostURL:  "https://ultradns.com",
	}

	client, err := client.NewClient(conf)

	if err != nil {
		fmt.Println(err)
		return
	}

	zoneService, err := zone.Get(client)

	if err != nil {
		fmt.Println(err)
		return
	}

	zoneProp := &zone.Properties{
		Name:        "zone_name",
		AccountName: "account_name",
		Type:        "PRIMARY",
	}

	primaryZone := &zone.PrimaryZone{
		CreateType: "NEW",
	}

	zone := &zone.Zone{
		Properties:        zoneProp,
		PrimaryCreateInfo: primaryZone,
	}

	res, err := zoneService.CreateZone(zone)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Zone Created Successfully")
}
```