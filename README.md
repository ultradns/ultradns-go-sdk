# UltraDNS SDK for Go

ultradns-go-sdk is the official UltraDNS SDK for the Go programming language.<br/>
Golang-Version : `1.22`

## SDK Features

ultradns-go-sdk is able to do CRUD operations on UltraDNS resources:

* Zones<br/>
	`Primary Zone`, `Secondary Zone`, `Alias Zone`
* Records<br/>
	`A`, `NS`, `CNAME`, `SOA`, `PTR`, `HINFO`, `MX`, `TXT`, `RP`, `AAAA`, `SRV`, `NAPTR`, `DS`, `SSHFP`, `TLSA`, `CDS`, `CDNSKEY`, `SVCB`, `HTTPS`, `SPF`, `CAA`, `APEXALIAS`
* Pools<br/>
	`Simple Failover(SF) Pool`, `Simple Load Balancing(SLB) Pool`, `Resource Distribution(RD) Pool`, `Directional(Dir) Pool`, `Sitebacker Pool(SB) Pool`, `Traffic Controller(TC) Pool`
* Probes<br/>
	`DNS`, `FTP`, `TCP`, `HTTP`, `PING`, `SMTP`, `SMTPSEND`
* Directional Group Geo
* Directional Group Geo

Also able to do read only operation on:

* Tasks

Jump To:
* [Getting Started](#Getting-Started)
* [Quick Examples](#Quick-Examples)
* [Getting Help](#Getting-Help)
* [More Resources](#Resources)

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
		HostURL:  "https://api.ultradns.com/",
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

### Configuring Credentials

The UltraDNS credentials to authenticate with services must be provided either as a config struct or as environment variable or a combination of both. Config struct values will be taken as priority if same config is provided as struct and environment variable. The required configs are: `USERNAME`, `PASSWORD`, `ENDPOINT`.

* Config Struct - SDK has config struct (`client.Config`) which has Username, Password, HostURL fields.

```go
conf := client.Config{
	Username: "username",
	Password: "password",
	HostURL:  "https://api.ultradns.com/",
}

client, err := client.NewClient(conf)
```

* Environment Variables - The required fields can be passed as environment variables.

```
export ULTRADNS_USERNAME="username"
export ULTRADNS_PASSWORD="password"
export ULTRADNS_HOST_URL="https://api.ultradns.com/"
```

```go
client, err := client.NewClient(client.Config{})
```

* Combination of Struct and Environment Variable - Few configs can be set as environment variable and other can be passed on config struct.

```
export ULTRADNS_PASSWORD="password"
export ULTRADNS_HOST_URL="https://api.ultradns.com/"
```

```go
client, err := client.NewClient(client.Config{Username: "username",})
```

### Configuring Service Clients

Service client can be created with either new config (`client.Config`) or existing client (`client.Client`). All service clients (`ZONE`,`RECORD and POOL`,`PROBE`,`TASK`) will follow common pattern of creation and usage.

Once the service's client is created you can use it to make CRUD operations on UltraDNS resources. These clients are safe to use concurrently.

* Service Client with New Config:

```go
zoneService, err := zone.New(client.Config{Username: "username",})
recordService, err := record.New(client.Config{Username: "username",})
```

* Service Client with Existing Client:

```go
client, err := client.NewClient(client.Config{Username: "username",})

zoneService, err := zone.Get(client)
recordService, err := record.Get(client)
```

### Configuring Client Loggers

Loggers can be enabled and disabled on demand for debugging purposes. By default the loggers will be disabled. Custom logger can be enabled with avialable loglevel (`client.logLevelType`) and flags ([GO Log Flags](https://pkg.go.dev/log#pkg-constants)).

```go
client, err := client.NewClient(client.Config{Username: "username",})

// Enable default debug logger.
client.EnableDefaultDebugLogger()

// Enable default trace logger.
client.EnableDefaultDebugLogger()

// Enable custom logger.
client.EnableLogger(client.LogError, log.LstdFlags)

// Disble logger.
client.DisableLogger()
```

### Service Client Usage

* CREATE:

```go
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

// create zone
res, err := zoneService.CreateZone(zone)
```

```go
rrSetKey := &rrset.RRSetKey{
	Owner:      "owner",
	Zone:       "zone.com.",
	RecordType: "A",
}

rdataInfo := &pool.RDataInfo{
	State:         "NORMAL",
	RunProbes:     true,
	Priority:      1,
	FailoverDelay: 0,
	Threshold:     1,
}

profile := &sbpool.Profile{
	RDataInfo:        []*pool.RDataInfo{rdataInfo},
	RunProbes:        true,
	ActOnProbes:      true,
	FailureThreshold: 0,
	Order:            "FIXED",
	MaxActive:        1,
	MaxServed:        1,
}

rrSet := &rrset.RRSet{
	OwnerName: "",
	RRType:    "A",
	RData:     []string{"1.1.1.1"},
	Profile:   profile,
}

//Create Record or pools
res, err := recordService.Create(rrSetKey, rrSet)
```

* READ:

```go
// Read a zone
// returns *http.Response, *zone.Response, error
res, resZone, err := zoneService.ReadZone("<zonename>")


rrSetKey := &rrset.RRSetKey{
	Owner:      "owner",
	Zone:       "zone.com.",
	RecordType: "A",
}
// Read a record
//returns *http.Response, *rrset.ResponseList, error
res, resList, err := recordService.Read(rrSetKey)
```

* UPDATE:

```go
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

// Update zone
res, err := zoneService.UpdateZone(zone)
```

```go
rrSetKey := &rrset.RRSetKey{
	Owner:      "owner",
	Zone:       "zone.com.",
	RecordType: "A",
}

rdataInfo := &pool.RDataInfo{
	State:         "NORMAL",
	RunProbes:     true,
	Priority:      1,
	FailoverDelay: 0,
	Threshold:     1,
}

profile := &sbpool.Profile{
	RDataInfo:        []*pool.RDataInfo{rdataInfo},
	RunProbes:        true,
	ActOnProbes:      true,
	FailureThreshold: 0,
	Order:            "FIXED",
	MaxActive:        1,
	MaxServed:        1,
}

rrSet := &rrset.RRSet{
	OwnerName: "",
	RRType:    "A",
	RData:     []string{"1.1.1.1"},
	Profile:   profile,
}

//Update Record or pools
res, err := recordService.Update(rrSetKey, rrSet)
```

* DELETE:

```go
// Delete a zone
res, err := zoneService.DeleteZone("<zonename>")

rrSetKey := &rrset.RRSetKey{
	Owner:      "owner",
	Zone:       "zone.com.",
	RecordType: "A",
}
// Delete a record
res, err := recordService.Delete(rrSetKey)
```

* LIST:

```go
query := &helper.QueryInfo{Limit: 1, Offset: 1}

// List a zone
// returns *http.Response, *zone.ResponseList, error
res, resZone, err := zoneService.ListZone(query)


rrSetKey := &rrset.RRSetKey{
	Owner:      "owner",
	Zone:       "zone.com.",
	RecordType: "ANY",
}
// List a record
//returns *http.Response, *rrset.ResponseList, error
res, resList, err := recordService.List(rrSetKey, query)
```

## Getting Help

* Open a support ticket with [Customer Support Ticket]().
* If you think you may have found a bug, please open an [issue](https://github.com/ultradns/ultradns-go-sdk/issues/new).
* Contact customer support [email]().

## Resources

* [SDK Examples](https://github.com/ultradns/ultradns-go-sdk/tree/master/examples) Included in the SDK's repo are several examples using the SDK features.