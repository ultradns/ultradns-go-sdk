package token_test

import (
	"context"
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/test"
	"github.com/ultradns/ultradns-go-sdk/internal/token"
	"golang.org/x/oauth2"
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

	if _, err := tokenSource.Token(); err != nil {
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

	if _, err := tokenSource.Token(); !strings.Contains(err.Error(), "invalid_request:password parameter is required for grant_type=password") {
		t.Fatal(err)
	}
}

func TestTokenFailureWithRefreshTokenFailure(t *testing.T) {
	tokenSource := getTokenSource()
	tokenSource.Password = ""
	tokenSource.T = &oauth2.Token{}

	if _, err := tokenSource.Token(); !strings.Contains(err.Error(), "invalid_request:password parameter is required for grant_type=password") {
		t.Fatal(err)
	}
}

func getTokenSource() *token.TokenSource {
	return &token.TokenSource{
		Ctx:      context.TODO(),
		Username: test.TestUsername,
		Password: test.TestPassword,
		BaseURL:  test.TestHost,
	}
}
