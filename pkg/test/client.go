package test

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

const (
	randStringLength   = 5
	randStringSet      = "abcdefghijklmnopqrstuvwxyz012346789"
	randZoneNamePrefix = "sdk-go-test-"
	randZoneNameSuffix = ".com."
)

var (
	testPassword  = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	testHost      = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
	testUserAgent = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")

	TestUsername = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	TestClient   *client.Client
)

func init() {
	client, err := client.NewClient(GetConfig())
	if err != nil {
		log.Panicf("unable to initialize test client for testing error : %s", err)
	}
	TestClient = client
}

func GetConfig() client.Config {
	return client.Config{
		Username:  TestUsername,
		Password:  testPassword,
		HostURL:   testHost,
		UserAgent: testUserAgent,
	}
}

func GetRandomZoneName() string {
	return randZoneNamePrefix + GetRandomString() + randZoneNameSuffix
}

func GetRandomString() string {
	result := make([]byte, randStringLength)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < randStringLength; i++ {
		result[i] = randStringSet[random.Intn(len(randStringSet))]
	}

	return string(result)
}
