package probe

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceName = "Probe"

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

func (s *Service) Create(rrSetKey *rrset.RRSetKey, probeData *Probe) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s create started", version.GetSDKVersion(), serviceName)

	if err := ValidateProbeDetails(probeData); err != nil {
		s.c.Error("[%s] %s create failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return nil, err
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.ProbeURI(), probeData, target)

	if err != nil {
		s.c.Error("[%s] %s create failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.CreateError(serviceName, rrSetKey.PID(), err)
	}

	s.c.Trace("[%s] %s create completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) Read(rrSetKey *rrset.RRSetKey) (*http.Response, *Probe, error) {
	probeTarget := &Probe{
		Details: getProbeDetails(rrSetKey.PType),
	}
	target := client.Target(probeTarget)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s read started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodGet, rrSetKey.ProbeURI(), nil, target)

	if err != nil {
		s.c.Error("[%s] %s read failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, errors.ReadError(serviceName, rrSetKey.PID(), err)
	}

	probeRes := target.Data.(*Probe)

	if rrSetKey.PType != probeRes.Type {
		return nil, nil, errors.ResourceTypeNotFoundError(serviceName, rrSetKey.PType, rrSetKey.PID())
	}

	s.c.Trace("[%s] %s read completed successfully", version.GetSDKVersion(), serviceName)

	return res, probeRes, nil
}

func (s *Service) Update(rrSetKey *rrset.RRSetKey, probeData *Probe) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s update started", version.GetSDKVersion(), serviceName)

	if err := ValidateProbeDetails(probeData); err != nil {
		s.c.Error("[%s] %s update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return nil, err
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.ProbeURI(), probeData, target)

	if err != nil {
		s.c.Error("[%s] %s update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.UpdateError(serviceName, rrSetKey.PID(), err)
	}

	s.c.Trace("[%s] %s update completed successfully", version.GetSDKVersion(), serviceName)
	return res, nil
}

func (s *Service) PartialUpdate(rrSetKey *rrset.RRSetKey, probeData *Probe) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s partial update started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodPatch, rrSetKey.ProbeURI(), probeData, target)

	if err != nil {
		s.c.Error("[%s] %s partial update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.PartialUpdateError(serviceName, rrSetKey.PID(), err)
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

	res, err := s.c.Do(http.MethodDelete, rrSetKey.ProbeURI(), nil, target)

	if err != nil {
		s.c.Error("[%s] %s delete failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.DeleteError(serviceName, rrSetKey.PID(), err)
	}

	s.c.Trace("[%s] %s delete completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) List(rrSetKey *rrset.RRSetKey, query *Query) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s list started", version.GetSDKVersion(), serviceName)

	rrSetKey.ID = ""

	res, err := s.c.Do(http.MethodGet, rrSetKey.ProbeListURI(query.String()), nil, target)

	if err != nil {
		s.c.Error("[%s] %s list failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, errors.ListError(serviceName, rrSetKey.ProbeURI(), err)
	}

	probesList := target.Data.(*ResponseList)

	s.c.Trace("[%s] %s list completed successfully", version.GetSDKVersion(), serviceName)
	return res, probesList, nil
}
