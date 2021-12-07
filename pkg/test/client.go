package test

import (
	"os"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

var (
	testUsername  = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	testPassword  = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	testHost      = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
	testUserAgent = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")
)

var TestClient *client.Client

// func init() {
// 	client, err := client.NewClient(GetConfig())
// 	if err != nil {
// 		log.Panicf("unable to initialize test client for testing error : %s", err)
// 	}
// 	TestClient = client
// }

func GetConfig() client.Config {
	return client.Config{
		Username:  testUsername,
		Password:  testPassword,
		HostURL:   testHost,
		UserAgent: testUserAgent,
	}
}
