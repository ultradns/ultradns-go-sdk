package geo

import (
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const (
	serviceName = "DirGroupGeo"
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

func (s *Service) CreateDirGroupGeo(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPost, dirGroupGeo.DirGroupGeoURI(), dirGroupGeo, target)

	if err != nil {
		geoGroupName := ""
		geoGroupName = dirGroupGeo.Name

		return nil, errors.CreateError(serviceName, geoGroupName, err)
	}

	return res, nil
}

func (s *Service) ReadDirGroupGeo(dirGroupGeo *DirGroupGeo) (*http.Response, *Response, error) {
	target := client.Target(&Response{})
	dirGroupGeoName := url.PathEscape(dirGroupGeo.Name)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, dirGroupGeo.DirGroupGeoURI(), nil, target)
	if err != nil {
		return nil, nil, errors.ReadError(serviceName, dirGroupGeoName, err)
	}

	dirGroupGeoResponse := target.Data.(*Response)

	return res, dirGroupGeoResponse, nil
}

func (s *Service) UpdateDirGroupGeo(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupGeoName := url.PathEscape(dirGroupGeo.Name)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPut, dirGroupGeo.DirGroupGeoURI(), dirGroupGeo, target)

	if err != nil {
		return nil, errors.UpdateError(serviceName, dirGroupGeoName, err)
	}

	return res, nil
}

func (s *Service) PartialUpdateDirGroupGeo(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupGeoName := url.PathEscape(dirGroupGeo.Name)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, dirGroupGeo.DirGroupGeoURI(), dirGroupGeo, target)

	if err != nil {
		return nil, errors.PartialUpdateError(serviceName, dirGroupGeoName, err)
	}

	return res, nil
}

func (s *Service) DeleteDirGroupGeo(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupGeoName := url.PathEscape(dirGroupGeo.Name)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, dirGroupGeo.DirGroupGeoURI(), nil, target)

	if err != nil {
		return nil, errors.DeleteError(serviceName, dirGroupGeoName, err)
	}

	return res, nil
}

func (s *Service) ListDirGroupGeo(queryInfo *helper.QueryInfo, dirGroupGeo *DirGroupGeo) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, dirGroupGeo.DirGroupGeoListURI(), nil, target)

	if err != nil {
		return nil, nil, errors.ListError(serviceName, dirGroupGeo.DirGroupGeoListURI(), err)
	}

	dirGroupGeoListResponse := target.Data.(*ResponseList)

	return res, dirGroupGeoListResponse, nil
}
