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

	s.c.Trace("%s create started", serviceName)

	res, err := s.c.Do(http.MethodPost, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		geoGroupName := dirGroupGeo.Name
		s.c.Error("%s create failed with error: %v", serviceName, err)
		return res, errors.CreateError(serviceName, geoGroupName, err)
	}

	s.c.Trace("%s create completed successfully", serviceName)

	return res, nil
}

func (s *Service) Read(dirGroupID string) (*http.Response, *Response, string, error) {
	target := client.Target(&Response{})

	dirGroupURI := helper.GetDirGroupURI(dirGroupID, DirGroupType)
	if s.c == nil {
		return nil, nil, dirGroupURI, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s read started", serviceName)

	res, err := s.c.Do(http.MethodGet, dirGroupURI, nil, target)
	if err != nil {
		s.c.Error("%s read failed with error: %v", serviceName, err)
		return res, nil, dirGroupURI, errors.ReadError(serviceName, dirGroupID, err)
	}

	dirGroupGeoResponse := target.Data.(*Response)

	s.c.Trace("%s read completed successfully", serviceName)

	return res, dirGroupGeoResponse, dirGroupURI, nil
}

func (s *Service) Update(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s update started", serviceName)

	res, err := s.c.Do(http.MethodPut, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		s.c.Error("%s update failed with error: %v", serviceName, err)
		return res, errors.UpdateError(serviceName, dirGroupGeo.Name, err)
	}

	s.c.Trace("%s update completed successfully", serviceName)
	return res, nil
}

func (s *Service) PartialUpdate(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s partial update started", serviceName)

	res, err := s.c.Do(http.MethodPatch, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		s.c.Error("%s partial update failed with error: %v", serviceName, err)
		return res, errors.PartialUpdateError(serviceName, dirGroupGeo.Name, err)
	}

	s.c.Trace("%s partial update completed successfully", serviceName)

	return res, nil
}

func (s *Service) Delete(dirGroupID string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s delete started", serviceName)

	res, err := s.c.Do(http.MethodDelete, helper.GetDirGroupURI(dirGroupID, DirGroupType), nil, target)

	if err != nil {
		s.c.Error("%s delete failed with error: %v", serviceName, err)
		return res, errors.DeleteError(serviceName, dirGroupID, err)
	}

	s.c.Trace("%s delete completed successfully", serviceName)

	return res, nil
}

func (s *Service) List(queryInfo *helper.QueryInfo, dirGroupGeo *DirGroupGeo) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s list started", serviceName)

	res, err := s.c.Do(http.MethodGet, helper.GetDirGroupListURI(dirGroupGeo.AccountName, DirGroupType), nil, target)

	if err != nil {
		s.c.Error("%s list failed with error: %v", serviceName, err)
		return res, nil, errors.ListError(serviceName, helper.GetDirGroupListURI(dirGroupGeo.AccountName, DirGroupType), err)
	}

	dirGroupGeoListResponse := target.Data.(*ResponseList)

	s.c.Trace("%s list completed successfully", serviceName)
	return res, dirGroupGeoListResponse, nil
}
