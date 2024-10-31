package service

import (
	"database/sql"
	"fmt"

	"github.com/golang-database/model"
	"github.com/golang-database/repository"
)

func InputDataCustomer(db *sql.DB, name string, address_id int, status string) error {
	if name == "" || address_id == 0 || status == "" {
		return fmt.Errorf("name, address id is required")
	}

	customerRepo := repository.NewCustomerRepository(db)
	customer := model.Customer{
		Name:    name,
		Address: model.Address{ID: address_id},
		Status:  status,
	}

	customerRepo.Create(&customer)
	fmt.Println("Success create data customer")
	return nil
}
