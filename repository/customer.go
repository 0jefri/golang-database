package repository

import (
	"database/sql"

	"github.com/golang-database/model"
	"github.com/golang-database/utils"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
}

type customerRepository struct {
	db *sql.DB
}

// Create implements CustomerRepository.
func (c *customerRepository) Create(payload *model.Customer) error {
	_, err := c.db.Exec(utils.CUSTOMER_INSERT, payload.Name, payload.Address.ID, payload.Status)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

// func (c *addressRepository) CreateCustomer(payload *model.Customer) error {
// 	_, err := c.db.Exec(utils.DRIVER_INSERT, payload.Name, payload.Address.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
