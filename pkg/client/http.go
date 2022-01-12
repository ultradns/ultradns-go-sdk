package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const contentType = "application/json"

var (
	defaultUserAgent = version.GetSDKVersion()
	errResponse      = errors.New("error response : ")
)

func (c *Client) Do(method, path string, payload, target interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, path)
	body := new(bytes.Buffer)

	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Accept", contentType)
	req.Header.Add("User-Agent", defaultUserAgent)
	req.Header.Add("User-Agent", c.userAgent)

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	er := validateResponse(res, target)

	if er != nil {
		return nil, er
	}

	return res, nil
}

func validateResponse(res *http.Response, t interface{}) error {
	if t == nil {
		return helper.ResponseTargetError("<nil>")
	}

	target, ok := t.(*Response)

	if !ok {
		return helper.ResponseTargetError(fmt.Sprintf("%T", target))
	}

	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
		if res.StatusCode == http.StatusNoContent {
			return nil
		}

		err := json.NewDecoder(res.Body).Decode(&target.Data)

		if err != nil {
			return err
		}
	} else {
		err := json.NewDecoder(res.Body).Decode(&target.Error)

		if err != nil {
			return err
		}

		return fmt.Errorf("%w %s", errResponse, target.Error[0])
	}

	return nil
}
