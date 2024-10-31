package repository

import (
	"database/sql"

	"github.com/golang-database/model"
	"github.com/golang-database/utils"
)

type AddressRepository interface {
	BaseRepository[model.Address]
}

type addressRepository struct {
	db *sql.DB
}

func NewAddressRepository(db *sql.DB) AddressRepository {
	return &addressRepository{db: db}
}

func (e *addressRepository) Create(payload *model.Address) error {
	_, err := e.db.Exec(utils.ADDRESS_INSERT, payload.Street, payload.City, payload.Province)
	if err != nil {
		return err
	}
	return nil
}
