package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/goharbor/harbor/src/common/models"
)

type ScannerRegistrationDAO interface {
	Create(registration *models.ScannerRegistration) error
	Update(registration *models.ScannerRegistration) error
	FindByID(id int64) (*models.ScannerRegistration, error)
	FindAll() ([]*models.ScannerRegistration, error)
	Delete(id int64) error
	SetAsDefault(id int64) error
	FindDefault() (*models.ScannerRegistration, error)
}

func NewScannerRegistrationDAO() ScannerRegistrationDAO {
	return &scannerRegistrationDAO{}
}

type scannerRegistrationDAO struct {
}

func (s *scannerRegistrationDAO) Create(registration *models.ScannerRegistration) error {
	_, err := GetOrmer().Insert(registration)
	return err
}

func (s *scannerRegistrationDAO) Update(registration *models.ScannerRegistration) error {
	_, err := GetOrmer().Update(registration)
	return err
}

func (s *scannerRegistrationDAO) FindByID(id int64) (*models.ScannerRegistration, error) {
	registration := &models.ScannerRegistration{
		ID: id,
	}
	if err := GetOrmer().Read(registration); err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return registration, nil
}

func (s *scannerRegistrationDAO) FindAll() ([]*models.ScannerRegistration, error) {
	var registrations []*models.ScannerRegistration
	qs := GetOrmer().
		QueryTable(&models.ScannerRegistration{}).
		Filter("deleted", 0).
		OrderBy("Name")
	_, err := qs.All(&registrations)
	return registrations, err
}

func (s *scannerRegistrationDAO) Delete(id int64) error {
	registration, err := s.FindByID(id)
	if err != nil {
		return err
	}
	registration.Name = fmt.Sprintf("%s#%d", registration.Name, registration.ID)
	registration.Deleted = true
	_, err = GetOrmer().Update(registration, "Name", "Deleted")
	return err
}

func (s *scannerRegistrationDAO) SetAsDefault(id int64) error {
	registration, err := s.FindByID(id)
	if err != nil {
		return fmt.Errorf("finidng scanner registraiton %d: %v", id, err)
	}
	registration.Default = true
	o := GetOrmer()
	_, err = o.Update(registration, "Default")
	if err != nil {
		return fmt.Errorf("setting scanner registration %d as default: %v", id, err)
	}

	queryParams := make([]interface{}, 2)
	sql := `update scanner_registration set default_flag = ? where id <> ?`
	queryParams = append(queryParams, false)
	queryParams = append(queryParams, id)
	r, err := o.Raw(sql, queryParams).Exec()
	if err != nil {
		return fmt.Errorf("unsetting scanner registrations as non-default: %v", err)
	}

	if _, err := r.RowsAffected(); err != nil {
		return err
	}

	return nil
}

func (s *scannerRegistrationDAO) FindDefault() (*models.ScannerRegistration, error) {
	registration := &models.ScannerRegistration{}
	qs := GetOrmer().
		QueryTable(&models.ScannerRegistration{}).
		Filter("Deleted", 0).
		Filter("Enabled", 1).
		Filter("Default", 1).
		OrderBy("Name")
	err := qs.One(registration)
	return registration, err
}
