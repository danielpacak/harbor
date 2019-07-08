package adapter

import (
	"github.com/goharbor/harbor/src/common/scanner"
)

const (
	// TODO Read from config
	EndpointURL = "http://harbor-microscanner-adapter:8080/api/v1"
)

type scannerAdapter struct {
	client *Client
}

func NewImageScannerAdapter(endpointURL string) scanner.ImageScanner {
	client := NewClient(endpointURL)

	return &scannerAdapter{
		client: client,
	}
}

func (ms *scannerAdapter) RequestScan(req scanner.ScanRequest) (*scanner.ScanResponse, error) {
	return ms.client.RequestScan(req)
}

func (ms *scannerAdapter) GetResult(detailsKey string) (*scanner.ScanResult, error) {
	return ms.client.GetScanResult(detailsKey)
}
