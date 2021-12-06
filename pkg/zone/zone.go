package zone

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/internal/util"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/task"
)

type ZoneService struct {
	c *client.Client
}

func New(config client.Config) (*ZoneService, error) {
	client, err := client.NewClient(config)

	if err != nil {
		return nil, err
	}
	return &ZoneService{c: client}, nil
}

func Get(client *client.Client) (*ZoneService, error) {
	if client == nil {
		return nil, fmt.Errorf("zone service is not properly configured")
	}
	return &ZoneService{c: client}, nil
}

func (zs *ZoneService) CreateZone(zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if zs.c == nil {
		return nil, fmt.Errorf("zone service is not properly configured")
	}

	res, err := zs.c.Do(http.MethodPost, "zones", zone, target)

	if err != nil {
		return nil, err
	}

	er := zs.checkZoneTask(res)

	if er != nil {
		return nil, er
	}

	return res, nil
}

func (zs *ZoneService) ReadZone(zoneName string) (*http.Response, *ZoneResponse, error) {
	target := client.Target(&ZoneResponse{})
	zoneName = url.PathEscape(zoneName)

	if zs.c == nil {
		return nil, nil, fmt.Errorf("zone service is not properly configured")
	}

	res, err := zs.c.Do(http.MethodGet, "zones/"+zoneName, nil, target)

	if err != nil {
		return nil, nil, err
	}

	zoneResponse := target.Data.(*ZoneResponse)

	return res, zoneResponse, nil
}

func (zs *ZoneService) UpdateZone(zoneName string, zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if zs.c == nil {
		return nil, fmt.Errorf("zone service is not properly configured")
	}

	res, err := zs.c.Do(http.MethodPut, "zones/"+zoneName, zone, target)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (zs *ZoneService) PatchUpdateZone(zoneName string, zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if zs.c == nil {
		return nil, fmt.Errorf("zone service is not properly configured")
	}

	res, err := zs.c.Do(http.MethodPatch, "zones/"+zoneName, zone, target)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (zs *ZoneService) DeleteZone(zoneName string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if zs.c == nil {
		return nil, fmt.Errorf("zone service is not properly configured")
	}

	res, err := zs.c.Do(http.MethodDelete, "zones/"+zoneName, nil, target)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (zs *ZoneService) ListZone(queryInfo *helper.QueryInfo) (*http.Response, *ZoneListResponse, error) {
	target := client.Target(&ZoneListResponse{})

	if zs.c == nil {
		return nil, nil, fmt.Errorf("zone service is not properly configured")
	}

	res, err := zs.c.Do(http.MethodGet, "v3/zones/"+queryInfo.String(), nil, target)

	if err != nil {
		return nil, nil, err
	}

	zoneListResponse := target.Data.(*ZoneListResponse)

	return res, zoneListResponse, nil
}

func (zs *ZoneService) checkZoneTask(res *http.Response) error {
	if res.StatusCode == http.StatusAccepted {
		taskID := res.Header.Get(util.TaskHeader)
		ts, err := task.Get(zs.c)

		if err != nil {
			return err
		}

		er := ts.TaskWait(taskID, 5, 10)

		if er != nil {
			return er
		}
	}
	return nil
}
