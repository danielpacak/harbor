package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/goharbor/harbor/src/common/scanner"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	endpointURL string
}

func NewClient(endpointURL string) *Client {
	return &Client{
		endpointURL: endpointURL,
	}
}

func (c *Client) RequestScan(request scanner.ScanRequest) error {
	url := c.endpointURL + "/scan"
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("invalid response status: %v %v", resp.StatusCode, resp.Status)
	}

	return nil
}

func (c *Client) GetScanReport(scanRequestID string) (*scanner.VulnerabilityReport, error) {
	res, err := c.doGetScanReport(scanRequestID)
	for err == nil && res.StatusCode == http.StatusFound {
		time.Sleep(10 * time.Second)
		res, err = c.doGetScanReport(scanRequestID)
	}
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response statu %s", res.Status)
	}

	b, _ := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()

	var report scanner.VulnerabilityReport
	err = json.Unmarshal(b, &report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (c *Client) doGetScanReport(scanRequestID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/scan/%s/report", c.endpointURL, scanRequestID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0")

	return http.DefaultTransport.RoundTrip(req)
}
