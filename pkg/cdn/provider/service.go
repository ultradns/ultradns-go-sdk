package provider

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
)

const serviceName = "CDNProvider"

type Service struct {
	c *client.Client
}

func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)

	if err != nil {
		return nil, errors.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) Create(accountName string, payload *Provider) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPost, providerURI(accountName, payload.ClientCdnID), payload, target)
	if err != nil {
		return res, errors.CreateError(serviceName, providerID(accountName, payload.ClientCdnID), err)
	}

	return res, nil
}

func (s *Service) Read(accountName, clientCdnID string) (*http.Response, *Provider, error) {
	target := client.Target(&Provider{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, providerURI(accountName, clientCdnID), nil, target)
	if err != nil {
		return res, nil, errors.ReadError(serviceName, providerID(accountName, clientCdnID), err)
	}

	return res, target.Data.(*Provider), nil
}

func (s *Service) Update(accountName, clientCdnID string, payload *Provider) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPut, providerURI(accountName, clientCdnID), payload, target)
	if err != nil {
		return res, errors.UpdateError(serviceName, providerID(accountName, clientCdnID), err)
	}

	return res, nil
}

func (s *Service) Delete(accountName, clientCdnID string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, providerURI(accountName, clientCdnID), nil, target)
	if err != nil {
		return res, errors.DeleteError(serviceName, providerID(accountName, clientCdnID), err)
	}

	return res, nil
}

func (s *Service) List(accountName string) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	uri := fmt.Sprintf("accounts/%s/cdn_providers", url.PathEscape(accountName))
	res, err := s.c.Do(http.MethodGet, uri, nil, target)
	if err != nil {
		return res, nil, errors.ListError(serviceName, uri, err)
	}

	return res, target.Data.(*ResponseList), nil
}

func providerURI(accountName, clientCdnID string) string {
	return fmt.Sprintf("accounts/%s/cdn_providers/%s", url.PathEscape(accountName), url.PathEscape(clientCdnID))
}

func providerID(accountName, clientCdnID string) string {
	return fmt.Sprintf("%s:%s", accountName, clientCdnID)
}