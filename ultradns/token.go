package ultradns

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

//tokenSource struct holds the credentials, url and context for generating token
type tokenSource struct {
	ctx      context.Context
	baseUrl  string
	username string
	password string
}

//generates access token if expires
//uses username, password to generate token
//return access token
func (ts *tokenSource) Token() (*oauth2.Token, error) {
	conf := oauth2.Config{Endpoint: ts.getTokenEndpoint()}
	return conf.PasswordCredentialsToken(ts.ctx, ts.username, ts.password)
}

//generates oauth2.Endpoint using token api url
//return oauth2.Endpoint
func (ts *tokenSource) getTokenEndpoint() oauth2.Endpoint {
	return oauth2.Endpoint{
		TokenURL: getTokenUrl(ts.baseUrl),
	}
}

//append token api path to the base url
//return token api url
func getTokenUrl(baseUrl string) string {
	return fmt.Sprintf("%s/authorization/token", baseUrl)
}
