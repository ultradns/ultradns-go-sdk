package test

import (
	"crypto/rand"
	"math/big"
	"os"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const (
	randStringLength                  = 5
	randSecondaryZoneCount            = 50
	randStringSet                     = "abcdefghijklmnopqrstuvwxyz012346789"
	randZoneNamePrefix                = "sdk-go-test-"
	randZoneNameSuffix                = ".com."
	randZoneNameWithSpecialCharSuffix = ".in-addr.arpa."
	recordTypeA                       = "A"
)

var (
	TestUsername                         = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	TestPassword                         = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
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

func GetRandomZoneNameWithSpecialChar() string {
	return randZoneNamePrefix + "/" + GetRandomString() + "/" + GetRandomString() + randZoneNameWithSpecialCharSuffix
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
		AccountName: TestUsername,
		Type:        zoneType,
	}
}

func GetRRSetKey(ownerName, zoneName, recordType string) *rrset.RRSetKey {
	return &rrset.RRSetKey{
		Name: ownerName,
		Zone: zoneName,
		Type: recordType,
	}
}

func GetRRSetTypeA(ownerName string) *rrset.RRSet {
	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    recordTypeA,
		RData:     []string{"192.168.1.1"},
	}
}
