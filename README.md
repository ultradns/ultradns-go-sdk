# ultradns-go-sdk

This is a golang SDK for the UltraDNS REST API.

The rest client object should be created with appropiate credentials. Using the client object CRUD operations of ultradns resources can be performed.

## Example

```go
package main

import (
	"fmt"
	
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func main() {

    client, err := ultradns.NewClient("username", "password", "hosturl", "version", "client name")
	if err != nil {
		fmt.Println(err)
		return
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        "zone_name",
		AccountName: "account_name",
		Type:        "PRIMARY",
	}
	primaryZone := ultradns.PrimaryZone{
		CreateType: "NEW",
	}
	zone := ultradns.Zone{
		Properties:        &zoneProp,
		PrimaryCreateInfo: &primaryZone,
	}

	res, err := client.CreateZone(zone)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(res.StatusCode)
}
```