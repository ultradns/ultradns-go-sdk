package sfpool

import (
	"fmt"
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const (
	serviceName = "SF-Pool"
	profileType = "*sfpool.Profile"
)

type Service struct {
	c *client.Client
}

func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)

	if err != nil {
		return nil, helper.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) CreateSFPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	if err := validateSFPoolProfile(rrSet); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.CreateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) UpdateSFPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	if err := validateSFPoolProfile(rrSet); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.UpdateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) PartialUpdateSFPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.PartialUpdateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) ReadSFPool(rrSetKey *rrset.RRSetKey) (*http.Response, *rrset.ResponseList, error) {
	sfPoolRRSet := &rrset.RRSet{
		Profile: &Profile{},
	}
	sfPoolResList := &rrset.ResponseList{}
	sfPoolResList.RRSets = make([]*rrset.RRSet, 1)
	sfPoolResList.RRSets[0] = sfPoolRRSet
	target := client.Target(sfPoolResList)

	if s.c == nil {
		return nil, nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, rrSetKey.URI(), nil, target)

	if err != nil {
		return nil, nil, helper.ReadError(serviceName, rrSetKey.ID(), err)
	}

	rrsetList := target.Data.(*rrset.ResponseList)

	return res, rrsetList, nil
}

func (s *Service) DeleteSFPool(rrSetKey *rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, rrSetKey.URI(), nil, target)

	if err != nil {
		return nil, helper.DeleteError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func validateSFPoolProfile(rrSet *rrset.RRSet) error {
	ptrProfileType := fmt.Sprintf("%T", rrSet.Profile)

	if ptrProfileType != profileType {
		return helper.TypeMismatchError(profileType, ptrProfileType)
	}

	rrSet.Profile.SetContext()

	sfProfile := rrSet.Profile.(*Profile)

	if !isSFRegionFailureSensitivityValid(sfProfile.RegionFailureSensitivity) {
		list := []string{"HIGH", "LOW"}

		return helper.UnknownDataError("SF-Pool Region Failure Sensitivity", sfProfile.RegionFailureSensitivity, list)
	}

	if !isSFMonitorMethodValid(sfProfile.Monitor.Method) {
		list := []string{"GET", "POST"}

		return helper.UnknownDataError("SF-Pool Monitor Method", sfProfile.Monitor.Method, list)
	}

	return nil
}

func isSFRegionFailureSensitivityValid(val string) bool {
	var sfRegionFailureSensitivity = map[string]bool{
		"HIGH": true,
		"LOW":  true,
	}

	_, ok := sfRegionFailureSensitivity[val]

	return ok
}

func isSFMonitorMethodValid(val string) bool {
	var sfMonitorMethod = map[string]bool{
		"GET":  true,
		"POST": true,
	}

	_, ok := sfMonitorMethod[val]

	return ok
}
