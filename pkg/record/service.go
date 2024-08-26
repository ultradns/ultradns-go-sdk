package record

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceName = "Record"

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

func (s *Service) Create(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s create started", version.GetSDKVersion(), serviceName)

	if err := validatePoolProfile(rrSet); err != nil {
		s.c.Error("[%s] %s create failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return nil, err
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		s.c.Error("[%s] %s create failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.CreateError(serviceName, rrSetKey.RecordID(), err)
	}

	s.c.Trace("[%s] %s create completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) Read(rrSetKey *rrset.RRSetKey) (*http.Response, *rrset.ResponseList, error) {
	target := client.Target(&rrset.ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s read started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodGet, rrSetKey.RecordURI(), nil, target)

	if err != nil {
		s.c.Error("[%s] %s read failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, errors.ReadError(serviceName, rrSetKey.RecordID(), err)
	}

	rrsetList := target.Data.(*rrset.ResponseList)

	if len(rrsetList.RRSets) != 1 {
		s.c.Error("[%s] %s read failed with error: multiple resource for the filter applied", version.GetSDKVersion(), serviceName)
		return nil, nil, errors.MultipleResourceFoundError(serviceName, rrSetKey.RecordID())
	}

	profile := rrsetList.RRSets[0].Profile

	if profile != nil && getPoolSchema(rrSetKey.PType) != profile.GetContext() {
		s.c.Error("[%s] %s read failed with error: queried pool data not available for the owner name", version.GetSDKVersion(), serviceName)
		return nil, nil, errors.ResourceTypeNotFoundError(serviceName, rrSetKey.PType, rrSetKey.RecordID())
	}

	s.c.Trace("[%s] %s read completed successfully", version.GetSDKVersion(), serviceName)

	return res, rrsetList, nil
}

func (s *Service) Update(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s update started", version.GetSDKVersion(), serviceName)

	if err := validatePoolProfile(rrSet); err != nil {
		s.c.Error("[%s] %s update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return nil, err
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		s.c.Error("[%s] %s update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.UpdateError(serviceName, rrSetKey.RecordID(), err)
	}

	s.c.Trace("[%s] %s update completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) PartialUpdate(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s partial update started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodPatch, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		s.c.Error("[%s] %s partial update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.PartialUpdateError(serviceName, rrSetKey.RecordID(), err)
	}

	s.c.Trace("[%s] %s partial update completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) Delete(rrSetKey *rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s delete started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodDelete, rrSetKey.RecordURI(), nil, target)

	if err != nil {
		s.c.Error("[%s] %s delete failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.DeleteError(serviceName, rrSetKey.RecordID(), err)
	}

	s.c.Trace("[%s] %s delete completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) List(rrSetKey *rrset.RRSetKey, queryInfo *helper.QueryInfo) (*http.Response, *rrset.ResponseList, error) {
	target := client.Target(&rrset.ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s list started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodGet, rrSetKey.RecordURI()+queryInfo.URI(), nil, target)

	if err != nil {
		s.c.Error("[%s] %s list failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, errors.ListError(serviceName, rrSetKey.RecordID()+queryInfo.URI(), err)
	}

	rrsetList := target.Data.(*rrset.ResponseList)

	s.c.Trace("[%s] %s list completed successfully", version.GetSDKVersion(), serviceName)

	return res, rrsetList, nil
}
