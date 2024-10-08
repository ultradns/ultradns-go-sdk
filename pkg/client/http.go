package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
)

const contentType = "application/json"
const throttleSleep = 1 * time.Second

var (
	defaultUserAgent = version.GetSDKVersion()
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

	userAgent := defaultUserAgent + ";" + c.userAgent

	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Accept", contentType)
	req.Header.Add("User-Agent", userAgent)

	c.logHttpRequest(req)
	res, err := c.httpClient.Do(req)
	c.logHttpResponse(res)

	resp := &http.Response{}

	if res != nil {
		resp.Status = res.Status
		resp.StatusCode = res.StatusCode
		resp.Header = res.Header
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		c.Warn("Throttling the request - '%s %s'", method, url)
		time.Sleep(throttleSleep)
		return c.Do(method, path, payload, target)
	}

	if err != nil {
		return resp, err
	}

	defer res.Body.Close()

	er := validateResponse(res, target)

	if er != nil {
		return resp, er
	}

	return resp, nil
}

func validateResponse(res *http.Response, t interface{}) error {
	if t == nil {
		return errors.ResponseTargetError("<nil>")
	}

	// Api Response should be always Response{} struct
	// or else throw ResponseTargetError
	// with current Response struct type
	target, ok := t.(*Response)

	if !ok {
		return errors.ResponseTargetError(fmt.Sprintf("%T", t))
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
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return err
		}

		err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&target.ErrorList)

		if err == nil {
			return errors.APIResponseError(target.ErrorList[0].String())
		}

		err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&target.Error)

		if err == nil {
			return errors.APIResponseError(target.Error.String())
		}

		return err
	}

	return nil
}
