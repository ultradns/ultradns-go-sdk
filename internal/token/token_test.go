package token

import (
	"context"
	"os"
	"strings"
	"testing"

	"golang.org/x/oauth2"
)

var (
	testUsername = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	testPassword = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	testHost     = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
)

func TestTokenSuccessWithPasswordCredentials(t *testing.T) {
	tokenSource := getTokenSource()
	token, err := tokenSource.Token()

	if err != nil {
		t.Fatal(err)
	}

	if token.TokenType != "Bearer" {
		t.Errorf("token type mismatched : expected - Bearer : found - %v", token.TokenType)
	}
}

func TestTokenSuccessWithRefreshTokenFailure(t *testing.T) {
	tokenSource := getTokenSource()
	_, err := tokenSource.Token()

	if err != nil {
		t.Fatal(err)
	}

	tokenSource.Ctx = nil

	token, er := tokenSource.Token()

	if er != nil {
		t.Fatal(er)
	}

	if token.TokenType != "Bearer" {
		t.Errorf("token type mismatched : expected - Bearer : found - %v", token.TokenType)
	}
}

func TestTokenFailureWithPasswordCredentials(t *testing.T) {
	tokenSource := getTokenSource()
	tokenSource.Password = ""
	_, err := tokenSource.Token()

	if !strings.Contains(err.Error(), "invalid_request:password parameter is required for grant_type=password") {
		t.Fatal(err)
	}
}

func TestTokenFailureWithRefreshTokenFailure(t *testing.T) {
	tokenSource := getTokenSource()
	tokenSource.Password = ""
	tokenSource.token = &oauth2.Token{}
	_, err := tokenSource.Token()

	if !strings.Contains(err.Error(), "invalid_request:password parameter is required for grant_type=password") {
		t.Fatal(err)
	}
}

func getTokenSource() *TokenSource {
	return &TokenSource{
		Ctx:      context.TODO(),
		Username: testUsername,
		Password: testPassword,
		BaseURL:  testHost,
	}
}
