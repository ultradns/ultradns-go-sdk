package probe

import (
	"net/http"

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

	if err := ValidateProbeDetails(probeData); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.ProbeURI(), probeData, target)

	if err != nil {
		return nil, errors.CreateError(serviceName, rrSetKey.PID(), err)
	}

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

	res, err := s.c.Do(http.MethodGet, rrSetKey.ProbeURI(), nil, target)

	if err != nil {
		return nil, nil, errors.ReadError(serviceName, rrSetKey.PID(), err)
	}

	probeRes := target.Data.(*Probe)

	if rrSetKey.PType != probeRes.Type {
		return nil, nil, errors.ResourceTypeNotFoundError(serviceName, rrSetKey.PType, rrSetKey.PID())
	}

	return res, probeRes, nil
}

func (s *Service) Update(rrSetKey *rrset.RRSetKey, probeData *Probe) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	if err := ValidateProbeDetails(probeData); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.ProbeURI(), probeData, target)

	if err != nil {
		return nil, errors.UpdateError(serviceName, rrSetKey.PID(), err)
	}

	return res, nil
}

func (s *Service) PartialUpdate(rrSetKey *rrset.RRSetKey, probeData *Probe) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, rrSetKey.ProbeURI(), probeData, target)

	if err != nil {
		return nil, errors.PartialUpdateError(serviceName, rrSetKey.PID(), err)
	}

	return res, nil
}

func (s *Service) Delete(rrSetKey *rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, rrSetKey.ProbeURI(), nil, target)

	if err != nil {
		return nil, errors.DeleteError(serviceName, rrSetKey.PID(), err)
	}

	return res, nil
}
