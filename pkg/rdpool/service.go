package rdpool

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const (
	serviceName = "RD-Pool"
	profileType = "*rdpool.Profile"
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

func (s *Service) CreateRDPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	if err := validateRDPoolProfile(rrSet); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.CreateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) UpdateRDPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	if err := validateRDPoolProfile(rrSet); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.UpdateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) PartialUpdateRDPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
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

func (s *Service) ReadRDPool(rrSetKey *rrset.RRSetKey) (*http.Response, *rrset.ResponseList, error) {
	rdPoolRRSet := &rrset.RRSet{
		Profile: &Profile{},
	}
	rdPoolResList := &rrset.ResponseList{}
	rdPoolResList.RRSets = make([]*rrset.RRSet, 1)
	rdPoolResList.RRSets[0] = rdPoolRRSet
	target := client.Target(rdPoolResList)

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

func (s *Service) DeleteRDPool(rrSetKey *rrset.RRSetKey) (*http.Response, error) {
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

func validateRDPoolProfile(rrSet *rrset.RRSet) error {
	if err := pool.ValidatePoolProfile(profileType, rrSet); err != nil {
		return err
	}

	rdProfile := rrSet.Profile.(*Profile)

	if !isRDPoolOrderValid(rdProfile.Order) {
		list := []string{"FIXED", "RANDOM", "ROUND_ROBIN"}

		return helper.UnknownDataError("RD-Pool order", rdProfile.Order, list)
	}

	return nil
}

func isRDPoolOrderValid(order string) bool {
	var rdPoolOrders = map[string]bool{
		"FIXED":       true,
		"RANDOM":      true,
		"ROUND_ROBIN": true,
	}

	_, ok := rdPoolOrders[order]

	return ok
}
