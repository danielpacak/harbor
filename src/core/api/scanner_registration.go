package api

import (
	"errors"
	"fmt"
	"github.com/goharbor/harbor/src/common/dao"
	"github.com/goharbor/harbor/src/common/models"
	"github.com/goharbor/harbor/src/common/utils/scanner/adapter"
)

type ScannerRegistrationAPI struct {
	BaseController
}

func (sra *ScannerRegistrationAPI) DAO() dao.ScannerRegistrationDAO {
	return dao.NewScannerRegistrationDAO()
}

func (sra *ScannerRegistrationAPI) List() {
	registrations, err := sra.DAO().FindAll()
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}

	sra.serveJSONData(registrations)
}

func (sra *ScannerRegistrationAPI) Post() {
	var registration models.ScannerRegistration
	err := sra.DecodeJSONReq(&registration)
	if err != nil {
		sra.SendInternalServerError(fmt.Errorf("decoding request: %v", err))
		return
	}
	err = sra.DAO().Create(&registration)
	if err != nil {
		sra.SendInternalServerError(fmt.Errorf("creating scanner registration: %v", err))
		return
	}
}

func (sra *ScannerRegistrationAPI) Put() {
	var registration models.ScannerRegistration
	err := sra.DecodeJSONReq(&registration)
	if err != nil {
		sra.SendInternalServerError(fmt.Errorf("decoding request: %v", err))
		return
	}
	err = sra.DAO().Update(&registration)
	if err != nil {
		sra.SendInternalServerError(fmt.Errorf("updating scanner registration: %v", err))
		return
	}
}

func (sra *ScannerRegistrationAPI) Get() {
	id, err := sra.GetInt64FromPath(":id")
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}
	registration, err := sra.DAO().FindByID(id)
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}
	if registration == nil {
		sra.SendNotFoundError(fmt.Errorf("scanner registration %d not found", id))
		return
	}
	sra.serveJSONData(registration)
}

func (sra *ScannerRegistrationAPI) Delete() {
	id, err := sra.GetInt64FromPath(":id")
	if err != nil {
		sra.SendInternalServerError(err)
	}
	err = sra.DAO().Delete(id)
	if err != nil {
		sra.SendInternalServerError(fmt.Errorf("deleting scanner registration: %v", err))
	}
}

func (sra *ScannerRegistrationAPI) SetDefault() {
	id, err := sra.GetInt64FromPath(":id")
	if err != nil {
		sra.SendInternalServerError(fmt.Errorf("getting id from path: %v", err))
	}

	err = sra.DAO().SetAsDefault(id)
	if err != nil {
		sra.SendInternalServerError(fmt.Errorf("setting scanner registration %d as default: %v", id, err))
	}
}

func (sra *ScannerRegistrationAPI) GetDefault() {
	registration, err := sra.DAO().FindDefault()
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}
	if registration == nil {
		sra.SendNotFoundError(errors.New("default scanner registration not found"))
		return
	}
	sra.serveJSONData(registration)
}

func (sra *ScannerRegistrationAPI) GetMetadata() {
	id, err := sra.GetInt64FromPath(":id")
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}
	registration, err := sra.DAO().FindByID(id)
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}

	client := adapter.NewClient(registration.EndpointURL)
	metadata, err := client.GetMetadata()
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}

	sra.serveJSONData(metadata)
}

func (sra *ScannerRegistrationAPI) Ping() {
	req := struct {
		EndpointURL   string `json:"endpoint_url"`
		Authorization string `json:"authorization"`
	}{}
	err := sra.DecodeJSONReq(&req)
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}
	client := adapter.NewClient(req.EndpointURL)
	metadata, err := client.GetMetadata()
	if err != nil {
		sra.SendInternalServerError(err)
		return
	}

	sra.serveJSONData(metadata)
}

func (sra *ScannerRegistrationAPI) serveJSONData(data interface{}) {
	sra.Data["json"] = data
	sra.ServeJSON()
}
