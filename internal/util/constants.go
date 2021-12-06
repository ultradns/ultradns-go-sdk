package util

import "os"

const (
	TaskHeader  = "X-Task-Id"
	UserAgent   = "golang-sdk-v1"
	ContentType = "application/json"
)

var (
	TestUsername          = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	TestPassword          = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	TestHost              = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
	TestVersion           = os.Getenv("ULTRADNS_UNIT_TEST_API_VERSION")
	TestUserAgent         = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")
	TestZoneName          = os.Getenv("ULTRADNS_UNIT_TEST_ZONE_NAME")
	TestZoneNameSecondary = "d100-permission.com."
)
