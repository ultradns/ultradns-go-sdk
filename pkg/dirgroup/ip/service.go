package ip

import (
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const (
	serviceName = "DirGroupIP"
)

type Service struct {
	c *client.Client
}

func New(cfn client.Config) (*Service, error) {
	c, err := client.NewClient(cfn)
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

func (s *Service) CreateDirGroupIP(dirGroupIP *DirGroupIP) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPost, dirGroupIP.DirGroupIPURI(), dirGroupIP, target)

	if err != nil {
		ipGroupName := ""
		ipGroupName = dirGroupIP.Name

		return nil, errors.CreateError(serviceName, ipGroupName, err)
	}

	return res, nil
}

func (s *Service) ReadDirGroupIP(dirGroupIP *DirGroupIP) (*http.Response, *Response, error) {
	target := client.Target(&Response{})
	dirGroupIPName := url.PathEscape(dirGroupIP.Name)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, dirGroupIP.DirGroupIPURI(), nil, target)
	if err != nil {
		return nil, nil, errors.ReadError(serviceName, dirGroupIPName, err)
	}

	dirGroupIPResponse := target.Data.(*Response)

	return res, dirGroupIPResponse, nil
}

func (s *Service) UpdateDirGroupIP(dirGroupIP *DirGroupIP) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupIPName := url.PathEscape(dirGroupIP.Name)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPut, dirGroupIP.DirGroupIPURI(), dirGroupIP, target)

	if err != nil {
		return nil, errors.UpdateError(serviceName, dirGroupIPName, err)
	}

	return res, nil
}

func (s *Service) PartialUpdateDirGroupIP(dirGroupIP *DirGroupIP) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupIPName := url.PathEscape(dirGroupIP.Name)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, dirGroupIP.DirGroupIPURI(), dirGroupIP, target)

	if err != nil {
		return nil, errors.PartialUpdateError(serviceName, dirGroupIPName, err)
	}

	return res, nil
}

func (s *Service) DeleteDirGroupIP(dirGroupIP *DirGroupIP) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupIPName := url.PathEscape(dirGroupIP.Name)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, dirGroupIP.DirGroupIPURI(), nil, target)

	if err != nil {
		return nil, errors.DeleteError(serviceName, dirGroupIPName, err)
	}

	return res, nil
}

func (s *Service) ListDirGroupIP(queryInfo *helper.QueryInfo, dirGroupIP *DirGroupIP) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, dirGroupIP.DirGroupIPListURI(), nil, target)

	if err != nil {
		return nil, nil, errors.ListError(serviceName, dirGroupIP.DirGroupIPListURI(), err)
	}

	dirGroupIPListResponse := target.Data.(*ResponseList)

	return res, dirGroupIPListResponse, nil
}
