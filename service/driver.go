package service

import (
	"database/sql"
	"fmt"

	"github.com/golang-database/model"
	"github.com/golang-database/repository"
)

func InputDataDriver(db *sql.DB, name string, address_id int) error {
	if name == "" || address_id == 0 {
		return fmt.Errorf("name, address id is required")
	}

	driverRepo := repository.NewDriverRepository(db)
	driver := model.Driver{
		Name:    name,
		Address: model.Address{ID: address_id},
	}

	err := driverRepo.Create(&driver)
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	fmt.Println("Success create data driver")
	return nil
}
