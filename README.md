# ultradns-go-sdk

This is a golang SDK for the UltraDNS REST API.

The rest client object should be created with appropiate credentials. Using the client object CRUD operations of ultradns resources can be performed.

## Example

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

	fmt.Println(res.StatusCode)
}
```