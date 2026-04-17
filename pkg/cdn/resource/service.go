package resource

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
)

const serviceName = "CDNResource"

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

func (s *Service) Create(accountName, fqdn string, payload *Resource) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPost, resourceURI(accountName, fqdn), payload, target)
	if err != nil {
		return res, errors.CreateError(serviceName, resourceID(accountName, fqdn), err)
	}

	return res, nil
}

func (s *Service) Read(accountName, fqdn string) (*http.Response, *Resource, error) {
	target := client.Target(&Resource{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, resourceURI(accountName, fqdn), nil, target)
	if err != nil {
		return res, nil, errors.ReadError(serviceName, resourceID(accountName, fqdn), err)
	}

	// The API response body does not include fqdn (it is a URL path param in
	// MultiCdnResource Java DTO). Inject it from the request so callers always
	// have it populated.
	result := target.Data.(*Resource)
	if result.FQDN == "" {
		result.FQDN = fqdn
	}

	return res, result, nil
}

func (s *Service) Update(accountName, fqdn string, payload *Resource) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPut, resourceURI(accountName, fqdn), payload, target)
	if err != nil {
		return res, errors.UpdateError(serviceName, resourceID(accountName, fqdn), err)
	}

	return res, nil
}

func (s *Service) Delete(accountName, fqdn string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, resourceURI(accountName, fqdn), nil, target)
	if err != nil {
		return res, errors.DeleteError(serviceName, resourceID(accountName, fqdn), err)
	}

	return res, nil
}

func (s *Service) List(accountName string, opts *ListOptions) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	if opts == nil {
		opts = &ListOptions{}
	}
	if opts.Size == 0 {
		opts.Size = 100
	}
	if opts.Page == 0 {
		opts.Page = 1
	}

	uri := fmt.Sprintf("accounts/%s/zones/cdn/config?size=%d&page=%d", url.PathEscape(accountName), opts.Size, opts.Page)
	res, err := s.c.Do(http.MethodGet, uri, nil, target)
	if err != nil {
		return res, nil, errors.ListError(serviceName, uri, err)
	}

	return res, target.Data.(*ResponseList), nil
}

func resourceURI(accountName, fqdn string) string {
	return fmt.Sprintf("accounts/%s/zones/cdn/%s", url.PathEscape(accountName), url.PathEscape(fqdn))
}

func resourceID(accountName, fqdn string) string {
	return fmt.Sprintf("%s:%s", accountName, fqdn)
}