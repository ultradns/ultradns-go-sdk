package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	defaultUserAgent = "golang-sdk-v1"
	contentType      = "application/json"
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
		return fmt.Errorf("response target should not be nil")
	}

	target, ok := t.(*Response)
	if !ok {
		return fmt.Errorf("response target mismatched : returned target - %T", target)
	}

	if res.StatusCode >= http.StatusOK && res.StatusCode <= http.StatusIMUsed {
		err := json.NewDecoder(res.Body).Decode(&target.Data)
		if err != nil && err.Error() == "EOF" {
			return nil
		} else if err != nil {
			return err
		}
	} else {
		err := json.NewDecoder(res.Body).Decode(&target.Error)
		if err != nil {
			return err
		}
	}

	return nil
}
