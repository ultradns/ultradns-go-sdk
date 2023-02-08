package integration

import (
	"crypto/rand"
	"math/big"
	"os"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const (
	randStringLength                  = 5
	randSecondaryZoneCount            = 50
	randStringSet                     = "abcdefghijklmnopqrstuvwxyz012346789"
	randZoneNamePrefix                = "golang-sdk-unit-test-"
	randZoneNameSuffix                = ".com."
	randZoneNameWithSpecialCharSuffix = ".in-addr.arpa."
	testRestrictIP                    = "192.168.1.1"
	testNotifyIP                      = "192.168.1.11"
	testPrimaryZoneCreateType         = "NEW"
)

var (
	TestUsername                         = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	TestPassword                         = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	TestAccount                          = os.Getenv("ULTRADNS_UNIT_TEST_ACCOUNT")
	TestHost                             = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
	TestUserAgent                        = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")
	TestPrimaryNameServer                = os.Getenv("ULTRADNS_UNIT_TEST_NAME_SERVER")
	TestClient            *client.Client = initializeTestClient()
)

func initializeTestClient() *client.Client {
	client, _ := client.NewClient(GetConfig())

	return client
}

func GetConfig() client.Config {
	return client.Config{
		Username:  TestUsername,
		Password:  TestPassword,
		HostURL:   TestHost,
		UserAgent: TestUserAgent,
	}
}

func GetRandomZoneName() string {
	return randZoneNamePrefix + GetRandomString() + randZoneNameSuffix
}

func GetRandomSecondaryZoneName() string {
	if num, err := rand.Int(rand.Reader, big.NewInt(randSecondaryZoneCount)); err == nil {
		return randZoneNamePrefix + num.String() + randZoneNameSuffix
	}

	return randZoneNamePrefix + "0" + randZoneNameSuffix
}

func GetRandomZoneNameWithSpecialChar() string {
	return randZoneNamePrefix + "/" + GetRandomString() + "/" + GetRandomString() + randZoneNameWithSpecialCharSuffix
}

func GetRandomString() string {
	result := make([]byte, randStringLength)

	for i := 0; i < randStringLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(randStringSet))))
		if err != nil {
			result[i] = randStringSet[0]

			continue
		}

		result[i] = randStringSet[num.Int64()]
	}

	return string(result)
}

func GetZoneProperties(zoneName, zoneType string) *zone.Properties {
	return &zone.Properties{
		Name:        zoneName,
		AccountName: TestAccount,
		Type:        zoneType,
	}
}

func GetPrimaryZone(zoneName string) *zone.Zone {
	restrictIP := &zone.RestrictIP{
		SingleIP: testRestrictIP,
	}
	notifyAddress := &zone.NotifyAddress{
		NotifyAddress: testNotifyIP,
	}
	primaryZone := &zone.PrimaryZone{
		CreateType:      testPrimaryZoneCreateType,
		RestrictIPList:  []*zone.RestrictIP{restrictIP},
		NotifyAddresses: []*zone.NotifyAddress{notifyAddress},
	}

	return &zone.Zone{
		Properties:        GetZoneProperties(zoneName, zone.Primary),
		PrimaryCreateInfo: primaryZone,
	}
}

func GetSecondaryZone(zoneName string) *zone.Zone {
	nameServerIP := &zone.NameServer{
		IP: TestPrimaryNameServer,
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

	return &zone.Zone{
		Properties:          GetZoneProperties(zoneName, zone.Secondary),
		SecondaryCreateInfo: secondaryZone,
	}
}

func GetAliasZone(alias, primary string) *zone.Zone {
	aliasZone := &zone.AliasZone{
		OriginalZoneName: primary,
	}

	return &zone.Zone{
		Properties:      GetZoneProperties(alias, zone.Alias),
		AliasCreateInfo: aliasZone,
	}
}

func GetRRSetKey(ownerName, zoneName, recordType, pType string) *rrset.RRSetKey {
	return &rrset.RRSetKey{
		Owner:      ownerName,
		Zone:       zoneName,
		RecordType: recordType,
		PType:      pType,
	}
}

func GetTestRRSetKey() *rrset.RRSetKey {
	return &rrset.RRSetKey{
		ID:         "id",
		Owner:      "www",
		Zone:       "non-existing-zone.com.",
		RecordType: "A",
	}
}

func GetDirGroupGeo(geoName, geoDescription string, geoCodes []string) *geo.DirGroupGeo {
	return &geo.DirGroupGeo{
		Name:        geoName,
		Description: geoDescription,
		AccountName: TestAccount,
		Codes:       geoCodes,
	}
}

func GetTestDirGroupGeo() *geo.DirGroupGeo {
	return &geo.DirGroupGeo{
		Name:        "testDirGroupGeo",
		Description: "Test GEO directional group",
		AccountName: TestAccount,
		Codes:       []string{"CA", "US", "MX"},
	}
}

func GetDirGroupIP(IPName, IPDescription string, IPs []*ip.IPAddress) *ip.DirGroupIP {
	return &ip.DirGroupIP{
		Name:        IPName,
		Description: IPDescription,
		AccountName: TestAccount,
		IPs:         IPs,
	}
}

func GetTestDirGroupIP() *ip.DirGroupIP {

	ipProp := &ip.DirGroupIP{
		AccountName: TestAccount,
		Name:        "Test SourceIP directional group",
		Description: "A test DirIP group",
		IPs: []*ip.IPAddress{
			&ip.IPAddress{
				Start: "192.168.1.1",
				End:   "192.168.1.10",
			},
			&ip.IPAddress{
				Cidr: "192.168.2.0/24",
			},
			&ip.IPAddress{
				Address: "192.168.3.3",
			},
		},
	}

	return ipProp
}
