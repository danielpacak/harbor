package scanner

import "github.com/goharbor/harbor/src/common/models"

type ScanRequest struct {
	RegistryURL   string `json:"registry_url"`
	RegistryToken string `json:"registry_token"`
	Repository    string `json:"repository"`
	Tag           string `json:"tag"`
	Digest        string `json:"digest"`
}

type ScanResponse struct {
	DetailsKey string `json:"details_key"`
}

type ScanResult struct {
	Digest          string                      `json:"digest"`
	Severity        models.Severity             `json:"severity"`
	Overview        *models.ComponentsOverview  `json:"overview"`
	Vulnerabilities []*models.VulnerabilityItem `json:"vulnerabilities"`
}

// ImageScanner defines methods of a pluggable image scanner.
//
// RequestScan sends image scan request to the actual scanner.
//
// GetResult pulls scan results from the actual scanner.
type ImageScanner interface {
	RequestScan(req ScanRequest) (*ScanResponse, error)
	GetResult(detailsKey string) (*ScanResult, error)
}
