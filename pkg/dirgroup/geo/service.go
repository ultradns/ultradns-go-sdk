package geo

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
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

	s.c.Trace("[%s] %s create started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodPost, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		geoGroupName := dirGroupGeo.Name
		s.c.Error("[%s] %s create failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.CreateError(serviceName, geoGroupName, err)
	}

	s.c.Trace("[%s] %s create completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) Read(dirGroupID string) (*http.Response, *Response, string, error) {
	target := client.Target(&Response{})

	dirGroupURI := helper.GetDirGroupURI(dirGroupID, DirGroupType)
	if s.c == nil {
		return nil, nil, dirGroupURI, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s read started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodGet, dirGroupURI, nil, target)
	if err != nil {
		s.c.Error("[%s] %s read failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, dirGroupURI, errors.ReadError(serviceName, dirGroupID, err)
	}

	dirGroupGeoResponse := target.Data.(*Response)

	s.c.Trace("[%s] %s read completed successfully", version.GetSDKVersion(), serviceName)

	return res, dirGroupGeoResponse, dirGroupURI, nil
}

func (s *Service) Update(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s update started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodPut, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		s.c.Error("[%s] %s update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.UpdateError(serviceName, dirGroupGeo.Name, err)
	}

	s.c.Trace("[%s] %s update completed successfully", version.GetSDKVersion(), serviceName)
	return res, nil
}

func (s *Service) PartialUpdate(dirGroupGeo *DirGroupGeo) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s partial update started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodPatch, helper.GetDirGroupURI(dirGroupGeo.DirGroupGeoID(), DirGroupType), dirGroupGeo, target)

	if err != nil {
		s.c.Error("[%s] %s partial update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.PartialUpdateError(serviceName, dirGroupGeo.Name, err)
	}

	s.c.Trace("[%s] %s partial update completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) Delete(dirGroupID string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s delete started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodDelete, helper.GetDirGroupURI(dirGroupID, DirGroupType), nil, target)

	if err != nil {
		s.c.Error("[%s] %s delete failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.DeleteError(serviceName, dirGroupID, err)
	}

	s.c.Trace("[%s] %s delete completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) List(queryInfo *helper.QueryInfo, dirGroupGeo *DirGroupGeo) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s list started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodGet, helper.GetDirGroupListURI(dirGroupGeo.AccountName, DirGroupType), nil, target)

	if err != nil {
		s.c.Error("[%s] %s list failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, errors.ListError(serviceName, helper.GetDirGroupListURI(dirGroupGeo.AccountName, DirGroupType), err)
	}

	dirGroupGeoListResponse := target.Data.(*ResponseList)

	s.c.Trace("[%s] %s list completed successfully", version.GetSDKVersion(), serviceName)
	return res, dirGroupGeoListResponse, nil
}
