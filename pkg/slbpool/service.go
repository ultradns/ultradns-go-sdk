package slbpool

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const (
	serviceName = "SLB-Pool"
	profileType = "*slbpool.Profile"
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

func (s *Service) CreateSLBPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	if err := validateSLBPoolProfile(rrSet); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.CreateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) UpdateSLBPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	if err := validateSLBPoolProfile(rrSet); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.UpdateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) PartialUpdateSLBPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
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

func (s *Service) ReadSLBPool(rrSetKey *rrset.RRSetKey) (*http.Response, *rrset.ResponseList, error) {
	slbPoolRRSet := &rrset.RRSet{
		Profile: &Profile{},
	}
	slbPoolResList := &rrset.ResponseList{}
	slbPoolResList.RRSets = make([]*rrset.RRSet, 1)
	slbPoolResList.RRSets[0] = slbPoolRRSet
	target := client.Target(slbPoolResList)

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

func (s *Service) DeleteSLBPool(rrSetKey *rrset.RRSetKey) (*http.Response, error) {
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

func validateSLBPoolProfile(rrSet *rrset.RRSet) error {
	if err := pool.ValidatePoolProfile(profileType, rrSet); err != nil {
		return err
	}

	slbProfile := rrSet.Profile.(*Profile)

	if !pool.IsRegionFailureSensitivityValid(slbProfile.RegionFailureSensitivity) {
		list := []string{"HIGH", "LOW"}

		return helper.UnknownDataError("SLB-Pool Region Failure Sensitivity", slbProfile.RegionFailureSensitivity, list)
	}

	if !pool.IsMonitorMethodValid(slbProfile.Monitor.Method) {
		list := []string{"GET", "POST"}

		return helper.UnknownDataError("SLB-Pool Monitor Method", slbProfile.Monitor.Method, list)
	}

	if !isSLBResponseMethodValid(slbProfile.ResponseMethod) {
		list := []string{"PRIORITY_HUNT", "RANDOM", "ROUND_ROBIN"}

		return helper.UnknownDataError("SLB-Pool Response Method", slbProfile.ResponseMethod, list)
	}

	if !isSLBServingPreferenceValid(slbProfile.ServingPreference) {
		list := []string{"AUTO_SELECT", "SERVE_PRIMARY", "SERVE_ALL_FAIL"}

		return helper.UnknownDataError("SLB-Pool Serving Preference", slbProfile.ServingPreference, list)
	}

	return nil
}

func isSLBResponseMethodValid(val string) bool {
	var slbResponseMethod = map[string]bool{
		"PRIORITY_HUNT": true,
		"RANDOM":        true,
		"ROUND_ROBIN":   true,
	}

	return pool.IsValidField(val, slbResponseMethod)
}

func isSLBServingPreferenceValid(val string) bool {
	var slbServingPreference = map[string]bool{
		"AUTO_SELECT":    true,
		"SERVE_PRIMARY":  true,
		"SERVE_ALL_FAIL": true,
	}

	return pool.IsValidField(val, slbServingPreference)
}
