package zone

import (
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/task"
)

const (
	serviceName     = "Zone"
	basePath        = "zones/"
	basePathForList = "v3/zones/"
	zoneTaskRetries = 5
	zoneTaskTimeGap = 10
	taskHeader      = "X-Task-Id"
	migratePrefix   = "v2/accounts/"
	migrateSuffix   = "/zones/move"
)

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

func (s *Service) CreateZone(zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s create started", serviceName)

	res, err := s.c.Do(http.MethodPost, basePath, zone, target)

	if err != nil {
		zoneName := ""
		if zone.Properties != nil {
			zoneName = zone.Properties.Name
		}

		s.c.Error("%s create failed with error: %v", serviceName, err)
		return res, errors.CreateError(serviceName, zoneName, err)
	}

	if er := s.checkZoneTask(res); er != nil {
		s.c.Error("%s create failed with error: %v", serviceName, err)
		return nil, er
	}

	s.c.Trace("%s create completed successfully", serviceName)

	return res, nil
}

func (s *Service) ReadZone(zoneName string) (*http.Response, *Response, error) {
	target := client.Target(&Response{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s read started", serviceName)

	res, err := s.c.Do(http.MethodGet, basePath+zoneName, nil, target)

	if err != nil {
		s.c.Error("%s read failed with error: %v", serviceName, err)
		return res, nil, errors.ReadError(serviceName, zoneName, err)
	}

	zoneResponse := target.Data.(*Response)

	s.c.Trace("%s read completed successfully", serviceName)

	return res, zoneResponse, nil
}

func (s *Service) UpdateZone(zoneName string, zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s update started", serviceName)

	res, err := s.c.Do(http.MethodPut, basePath+zoneName, zone, target)

	if err != nil {
		s.c.Error("%s update failed with error: %v", serviceName, err)
		return res, errors.UpdateError(serviceName, zoneName, err)
	}

	s.c.Trace("%s update completed successfully", serviceName)

	return res, nil
}

func (s *Service) PartialUpdateZone(zoneName string, zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s partial update started", serviceName)

	res, err := s.c.Do(http.MethodPatch, basePath+zoneName, zone, target)

	if err != nil {
		s.c.Error("%s partial update failed with error: %v", serviceName, err)
		return res, errors.PartialUpdateError(serviceName, zoneName, err)
	}

	s.c.Trace("%s partial update completed successfully", serviceName)

	return res, nil
}

func (s *Service) DeleteZone(zoneName string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s delete started", serviceName)

	res, err := s.c.Do(http.MethodDelete, basePath+zoneName, nil, target)

	if err != nil {
		s.c.Error("%s delete failed with error: %v", serviceName, err)
		return res, errors.DeleteError(serviceName, zoneName, err)
	}

	s.c.Trace("%s delete completed successfully", serviceName)

	return res, nil
}

func (s *Service) ListZone(queryInfo *helper.QueryInfo) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s list started", serviceName)

	res, err := s.c.Do(http.MethodGet, basePathForList+queryInfo.URI(), nil, target)

	if err != nil {
		s.c.Error("%s list failed with error: %v", serviceName, err)
		return res, nil, errors.ListError(serviceName, basePathForList+queryInfo.URI(), err)
	}

	zoneListResponse := target.Data.(*ResponseList)

	s.c.Trace("%s list completed successfully", serviceName)

	return res, zoneListResponse, nil
}

func (s *Service) checkZoneTask(res *http.Response) error {
	s.c.Trace("check zone %s started", serviceName)
	if res.StatusCode == http.StatusAccepted {
		taskID := res.Header.Get(taskHeader)
		taskService, err := task.Get(s.c)

		if err != nil {
			return err
		}

		er := taskService.TaskWait(taskID, zoneTaskRetries, zoneTaskTimeGap)

		if er != nil {
			return er
		}
	}

	s.c.Trace("check zone %s completed successfully", serviceName)

	return nil
}

func (s *Service) MigrateZoneAccount(zones []string, old, new string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	old = url.PathEscape(old)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s account migration started", serviceName)

	data := &ZoneAccountChange{Zones: zones, TargetAccount: new}
	res, err := s.c.Do(http.MethodPut, migratePrefix+old+migrateSuffix, data, target)

	if err != nil {
		s.c.Error("%s account migration failed with error: %v", serviceName, err)
		return res, errors.MigrateError(serviceName, "", err)
	}

	s.c.Trace("%s account migration completed successfully", serviceName)

	return res, nil
}
