package scanner

import (
	"github.com/goharbor/harbor/src/common/models"
)

type Registry struct {
	URL           string `json:"url"`
	Authorization string `json:"authorization"`
}

type Artifact struct {
	Repository string `json:"repository"`
	Digest     string `json:"digest"`
}

type ScanRequest struct {
	Registry Registry `json:"registry"`
	Artifact Artifact `json:"artifact"`
}

type ScanResponse struct {
	ID string `json:"id"`
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
	RequestScan(req ScanRequest) (*ScanResponse, error)
	GetScanReport(scanRequestID string) (*VulnerabilityReport, error)
}
