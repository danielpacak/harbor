package models

type ScannerMetadata struct {
	Name         string        `json:"name"`
	Vendor       string        `json:"vendor"`
	Version      string        `json:"version"`
	Capabilities []*Capability `json:"capabilities"`
}

type Capability struct {
	ArtifactMIMETypes []string `json:"artifact_mime_types"`
	ReportMIMETypes   []string `json:"report_mime_types"`
}
