package adapter

import (
	"github.com/goharbor/harbor/src/common/scanner"
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

func (sa *scannerAdapter) RequestScan(req scanner.ScanRequest) error {
	return sa.client.RequestScan(req)
}

func (sa *scannerAdapter) GetScanReport(scanRequestID string) (*scanner.VulnerabilityReport, error) {
	return sa.client.GetScanReport(scanRequestID)
}
