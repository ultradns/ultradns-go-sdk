package token

import (
	"context"
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/util"
	"golang.org/x/oauth2"
)

func TestTokenSuccessWithPasswordCredentials(t *testing.T) {
	ts := getTokenSource()
	token, err := ts.Token()

	if err != nil {
		t.Fatal(err)
	}
	validateTokenType(token.TokenType, t)
}

func TestTokenSuccessWithRefreshTokenFailure(t *testing.T) {
	ts := getTokenSource()
	_, err := ts.Token()

	if err != nil {
		t.Fatal(err)
	}

	ts.Ctx = nil

	token, er := ts.Token()

	if er != nil {
		t.Fatal(er)
	}
	validateTokenType(token.TokenType, t)

}

func TestTokenFailureWithPasswordCredentials(t *testing.T) {
	ts := getTokenSource()
	ts.Password = ""
	_, err := ts.Token()

	if !strings.Contains(err.Error(), "invalid_request:password parameter is required for grant_type=password") {
		t.Fatal(err)
	}

}

func TestTokenFailureWithRefreshTokenFailure(t *testing.T) {
	ts := getTokenSource()
	ts.Password = ""
	ts.token = &oauth2.Token{}
	_, err := ts.Token()

	if !strings.Contains(err.Error(), "invalid_request:password parameter is required for grant_type=password") {
		t.Fatal(err)
	}

}

func getTokenSource() *TokenSource {
	return &TokenSource{
		Ctx:      context.TODO(),
		Username: util.TestUsername,
		Password: util.TestPassword,
		BaseURL:  util.TestHost,
	}
}

func validateTokenType(tokenType string, t *testing.T) {
	expected := "Bearer"
	found := tokenType

	if expected != found {
		t.Errorf("token type mismatched : expected - %v : found - %v", expected, found)
	}
}
