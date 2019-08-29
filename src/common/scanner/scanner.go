package scanner

import (
	"github.com/goharbor/harbor/src/common/models"
)

type ScanRequest struct {
	ID                    string `json:"id"`
	RegistryURL           string `json:"registry_url"`
	RegistryAuthorization string `json:"registry_authorization"`
	ArtifactRepository    string `json:"artifact_repository"`
	ArtifactDigest        string `json:"artifact_digest"`
}

type VulnerabilityReport struct {
	Severity        models.Severity             `json:"severity"`
	Vulnerabilities []*models.VulnerabilityItem `json:"vulnerabilities"`
}

func (r *VulnerabilityReport) ToComponentsOverview() *models.ComponentsOverview {
	overallSev := models.SevNone
	total := 0
	sevToCount := map[models.Severity]int{
		models.SevHigh:    0,
		models.SevMedium:  0,
		models.SevLow:     0,
		models.SevUnknown: 0,
		models.SevNone:    0,
	}

	for _, v := range r.Vulnerabilities {
		sev := v.Severity
		sevToCount[sev]++
		total++
		if sev > overallSev {
			overallSev = sev
		}
	}

	var summary []*models.ComponentsOverviewEntry
	for k, v := range sevToCount {
		summary = append(summary, &models.ComponentsOverviewEntry{
			Sev:   int(k),
			Count: v,
		})
	}

	return &models.ComponentsOverview{
		Total:   total,
		Summary: summary,
	}
}

// ImageScanner defines methods of a pluggable image scanner.
//
// RequestScan sends image scan request to the actual scanner.
//
// GetResult pulls scan results from the actual scanner.
type ImageScanner interface {
	RequestScan(req ScanRequest) error
	GetScanReport(scanRequestID string) (*VulnerabilityReport, error)
}
