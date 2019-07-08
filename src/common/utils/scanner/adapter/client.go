package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/goharbor/harbor/src/common/scanner"
	"io/ioutil"
	"net/http"
)

type Client struct {
	endpointURL string
}

func NewClient(endpointURL string) *Client {
	return &Client{
		endpointURL: endpointURL,
	}
}

func (c *Client) RequestScan(request scanner.ScanRequest) (*scanner.ScanResponse, error) {
	url := c.endpointURL + "/scan"
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("invalid response status: %v %v", resp.StatusCode, resp.Status)
	}

	b, _ = ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()

	var result scanner.ScanResponse
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetScanResult(detailsKey string) (*scanner.ScanResult, error) {
	url := c.endpointURL + "/scan/" + detailsKey
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()

	var result scanner.ScanResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
