package service

import (
	"database/sql"
	"fmt"

	"github.com/golang-database/model"
	"github.com/golang-database/repository"
)

func InputDataAddress(db *sql.DB, street string, city string, province string) error {
	if street == "" || city == "" || province == "" {
		return fmt.Errorf("street, city, province is required")
	}
	addressRepo := repository.NewAddressRepository(db)
	address := model.Address{
		Street:   street,
		City:     city,
		Province: province,
	}

	addressRepo.Create(&address)
	fmt.Println("Success input data dengan id", address.ID)
	return nil
}
