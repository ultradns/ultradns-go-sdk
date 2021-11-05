package ultradns

import "testing"

var (
	testUserName  = "user"
	testPassword  = "password"
	testUrl       = "url"
	testVersion   = "version"
	testUserAgent = "test"
)

func TestNewClientWithAllParams(t *testing.T) {
	testClient, err := NewClient(testUserName, testPassword, testUrl, testVersion, testUserAgent)
	if testClient == nil {
		t.Fatalf("Client object is nil - Error : %v", err)
	}
	

}
