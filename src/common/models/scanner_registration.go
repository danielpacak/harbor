package models

type ScannerRegistration struct {
	ID          int64  `orm:"pk;auto;column(id)" json:"id"`
	Name        string `orm:"column(name)" json:"name"`
	EndpointURL string `orm:"column(endpoint_url)" json:"endpoint_url"`
	Default     bool   `orm:"column(default_flag)" json:"default"`
	Enabled     bool   `orm:"column(enabled_flag)" json:"enabled"`
	Deleted     bool   `orm:"column(deleted_flag)" json:"deleted"`
}

func (p *ScannerRegistration) TableName() string {
	return "scanner_registration"
}
