/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
package ultradns

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

//client struct wraps the http client and ultradns api base url
type Client struct {
	httpClient *http.Client
	baseUrl    *url.URL
	userAgent  string
}

//returns new ultradns client instance
func NewClient(username, password, hostUrl, apiVersion, userAgent string) (*Client, error) {
	if err := validateParameter("User Name", username); err != nil {
		return nil, err
	}
	if err := validateParameter("Password", password); err != nil {
		return nil, err
	}
	if err := validateParameter("Host Url", hostUrl); err != nil {
		return nil, err
	}
	if err := validateParameter("Api Version", apiVersion); err != nil {
		return nil, err
	}
	if err := validateParameter("User Agent", userAgent); err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(hostUrl + apiVersion)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	ts := tokenSource{
		ctx:      ctx,
		baseUrl:  baseUrl.String(),
		username: username,
		password: password,
	}
	client := &Client{
		httpClient: oauth2.NewClient(ctx, oauth2.ReuseTokenSource(nil, &ts)),
		baseUrl:    baseUrl,
		userAgent:  userAgent,
	}
	return client, nil
}

func validateParameter(key, value string) error {
	if value != "" {
		return nil
	}
	return fmt.Errorf("%v is required to create a new http client for UltraDNS rest api", key)
}
