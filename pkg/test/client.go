package test

import (
	"log"

	"github.com/ultradns/ultradns-go-sdk/internal/util"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

var TestClient *client.Client

func init() {
	client, err := client.NewClient(GetConfig())
	if err != nil {
		log.Panicf("unable to initialize test client for testing error : %s", err)
	}
	TestClient = client
}

func GetConfig() client.Config {
	return client.Config{
		Username:  util.TestUsername,
		Password:  util.TestPassword,
		HostURL:   util.TestHost,
		UserAgent: util.TestUserAgent,
	}
}
