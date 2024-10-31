package repository

import (
	"database/sql"

	"github.com/golang-database/model"
	"github.com/golang-database/utils"
)

type DriverRepository interface {
	BaseRepository[model.Driver]
}

type driverRepository struct {
	db *sql.DB
}

func (d *driverRepository) Create(payload *model.Driver) error {
	_, err := d.db.Exec(utils.DRIVER_INSERT, payload.Name, payload.Address.ID)
	if err != nil {
		return err
	}
	return nil
}

func NewDriverRepository(db *sql.DB) DriverRepository {
	return &driverRepository{db: db}
}
