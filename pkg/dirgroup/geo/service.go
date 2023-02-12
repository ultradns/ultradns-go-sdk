package geo

import (
	"net/http"

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

func (s *Service) Create(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPost, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		geoGroupName := ""
		geoGroupName = dirGroupGeo.Name

		return nil, errors.CreateError(serviceName, geoGroupName, err)
	}

	return res, nil
}

func (s *Service) Read(dirGroupID string) (*http.Response, *Response, string, error) {
	target := client.Target(&Response{})

	dirGroupURI := helper.GetDirGroupURI(dirGroupID, DirGroupType)
	if s.c == nil {
		return nil, nil, dirGroupURI, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, dirGroupURI, nil, target)
	if err != nil {
		return nil, nil, dirGroupURI, errors.ReadError(serviceName, dirGroupID, err)
	}

	dirGroupGeoResponse := target.Data.(*Response)

	return res, dirGroupGeoResponse, dirGroupURI, nil
}

func (s *Service) Update(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPut, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		return nil, errors.UpdateError(serviceName, dirGroupGeo.Name, err)
	}

	return res, nil
}

func (s *Service) PartialUpdate(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		return nil, errors.PartialUpdateError(serviceName, dirGroupGeo.Name, err)
	}

	return res, nil
}

func (s *Service) Delete(dirGroupID string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, helper.GetDirGroupURI(dirGroupID, DirGroupType), nil, target)

	if err != nil {
		return nil, errors.DeleteError(serviceName, dirGroupID, err)
	}

	return res, nil
}

func (s *Service) List(queryInfo *helper.QueryInfo, dirGroupGeo *DirGroupGeo) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, helper.GetDirGroupListURI(dirGroupGeo.AccountName, DirGroupType), nil, target)

	if err != nil {
		return nil, nil, errors.ListError(serviceName, helper.GetDirGroupListURI(dirGroupGeo.AccountName, DirGroupType), err)
	}

	dirGroupGeoListResponse := target.Data.(*ResponseList)

	return res, dirGroupGeoListResponse, nil
}
