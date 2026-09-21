package webforward

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const (
	serviceName           = "WebForward"
	basePath              = "zones/"
	webForwardPath        = "/webforwards"
	accountWebForwardPath = "accounts/webforwards"
)

// Service wraps the Ultradns web forward service.
type Service struct {
	c *client.Client
}

// New creates a web forward service with a newly initialized client.
func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)

	if err != nil {
		return nil, errors.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

// Get returns a web forward service backed by the given client.
func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

// Create adds a web forward to a zone. The API returns the system-generated
// guid in the response body; the returned WebForward carries only that field.
func (s *Service) Create(zoneName string, webForwardData *WebForward) (*http.Response, *WebForward, error) {
	target := client.Target(&WebForward{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s create started", serviceName)

	res, err := s.c.Do(http.MethodPost, basePath+zoneName+webForwardPath, webForwardData, target)

	if err != nil {
		s.c.Error("%s create failed with error: %v", serviceName, err)
		return res, nil, errors.CreateError(serviceName, zoneName, err)
	}

	webForwardResponse := target.Data.(*WebForward)

	s.c.Trace("%s create completed successfully", serviceName)

	return res, webForwardResponse, nil
}

// Read returns a web forward by guid. The API exposes no single-item read
// endpoint, so the zone's web forwards are listed and matched by guid.
func (s *Service) Read(zoneName, guid string) (*http.Response, *WebForward, error) {
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s read started", serviceName)

	res, listResponse, err := s.List(zoneName, nil)

	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			s.c.Error("%s read failed with error: web forward not found", serviceName)
			return res, nil, errors.ResourceTypeNotFoundError(serviceName, "webForward", guid)
		}

		s.c.Error("%s read failed with error: %v", serviceName, err)
		return res, nil, errors.ReadError(serviceName, guid, err)
	}

	for _, webForward := range listResponse.WebForwards {
		if strings.EqualFold(webForward.GUID, guid) {
			s.c.Trace("%s read completed successfully", serviceName)
			return res, webForward, nil
		}
	}

	s.c.Error("%s read failed with error: web forward not found", serviceName)
	return res, nil, errors.ResourceTypeNotFoundError(serviceName, "webForward", guid)
}

// Update replaces the configuration of an existing web forward.
func (s *Service) Update(zoneName, guid string, webForwardData *WebForward) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)
	guid = url.PathEscape(guid)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s update started", serviceName)

	res, err := s.c.Do(http.MethodPut, basePath+zoneName+webForwardPath+"/"+guid, webForwardData, target)

	if err != nil {
		s.c.Error("%s update failed with error: %v", serviceName, err)
		return res, errors.UpdateError(serviceName, guid, err)
	}

	s.c.Trace("%s update completed successfully", serviceName)

	return res, nil
}

// Delete removes a web forward by guid.
func (s *Service) Delete(zoneName, guid string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)
	guid = url.PathEscape(guid)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s delete started", serviceName)

	res, err := s.c.Do(http.MethodDelete, basePath+zoneName+webForwardPath+"/"+guid, nil, target)

	if err != nil {
		s.c.Error("%s delete failed with error: %v", serviceName, err)
		return res, errors.DeleteError(serviceName, guid, err)
	}

	s.c.Trace("%s delete completed successfully", serviceName)

	return res, nil
}

// List returns all web forwards for a zone. A zone without web forwards is
// reported by the API as a 404, which is returned as an error.
func (s *Service) List(zoneName string, queryInfo *helper.QueryInfo) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s list started", serviceName)

	path := basePath + zoneName + webForwardPath

	if queryInfo != nil {
		path = path + queryInfo.URI()
	}

	res, err := s.c.Do(http.MethodGet, path, nil, target)

	if err != nil {
		s.c.Error("%s list failed with error: %v", serviceName, err)
		return res, nil, errors.ListError(serviceName, path, err)
	}

	webForwardListResponse := target.Data.(*ResponseList)

	s.c.Trace("%s list completed successfully", serviceName)

	return res, webForwardListResponse, nil
}

// ListAccount returns all web forwards configured at the account level,
// across every zone of the account. Each returned web forward includes
// its zoneName and accountName.
func (s *Service) ListAccount() (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s account list started", serviceName)

	res, err := s.c.Do(http.MethodGet, accountWebForwardPath, nil, target)

	if err != nil {
		s.c.Error("%s account list failed with error: %v", serviceName, err)
		return res, nil, errors.ListError(serviceName, accountWebForwardPath, err)
	}

	webForwardListResponse := target.Data.(*ResponseList)

	s.c.Trace("%s account list completed successfully", serviceName)

	return res, webForwardListResponse, nil
}
