/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
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
	token    *oauth2.Token
}

//generates access token if expires
//uses refresh token to refresh access token
//uses username, password to generate token if refresh token is nil
//return access token
func (ts *tokenSource) Token() (*oauth2.Token, error) {
	conf := &oauth2.Config{Endpoint: ts.getTokenEndpoint()}
	if ts.token == nil {
		return ts.PasswordCredentialsToken(conf)
	}
	token, err := conf.TokenSource(ts.ctx, ts.token).Token()
	if err != nil {
		return ts.PasswordCredentialsToken(conf)
	}
	ts.token = token
	return token, err
}

//generates access token using credentials
//set token to tokensource
func (ts *tokenSource) PasswordCredentialsToken(conf *oauth2.Config) (token *oauth2.Token, err error) {
	token, err = conf.PasswordCredentialsToken(ts.ctx, ts.username, ts.password)
	ts.token = token
	return token, err
}

//generates oauth2.Endpoint using token api url
//return oauth2.Endpoint
func (ts *tokenSource) getTokenEndpoint() oauth2.Endpoint {
	return oauth2.Endpoint{
		TokenURL:  getTokenUrl(ts.baseUrl),
		AuthURL:   getTokenUrl(ts.baseUrl),
		AuthStyle: 2,
	}
}

//append token api path to the base url
//return token api url
func getTokenUrl(baseUrl string) string {
	return fmt.Sprintf("%s/authorization/token", baseUrl)
}
