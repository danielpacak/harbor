package adapter

import (
	"fmt"
	"github.com/goharbor/harbor/src/common/dao"
	"github.com/goharbor/harbor/src/common/scanner"
)

func GetImageScanner() (scanner.ImageScanner, error) {
	registration, err := dao.NewScannerRegistrationDAO().FindDefault()
	if err != nil {
		return nil, fmt.Errorf("finding default scanner registration: %v", err)
	}
	if registration == nil {
		return nil, fmt.Errorf("default scanner registration not set")
	}
	fmt.Printf("DEFAULT SCANNER REGISTRATION ENDPOINT URL: %s\n", registration.EndpointURL)

	return NewImageScannerAdapter(registration.EndpointURL), nil
}
