package repository

import (
	"database/sql"
	"log"

	"github.com/golang-database/model"
	"github.com/golang-database/model/dto"
	"github.com/golang-database/utils"
)

type OrderRepository interface {
	BaseRepository[model.Order]
	GetMonthlyOrder() ([]dto.MonthlyOrder, error)
	GetMonthlyOrderCustomer() ([]dto.MonthlyCustomerOrder, error)
	// GetAddressOrderSummary() ([]dto.AddressOrderSummary, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (c *orderRepository) Create(payload *model.Order) error {
	_, err := c.db.Exec(utils.ORDER_INSERT, payload.CustomerId, payload.DriverId, payload.OrderDate, payload.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) GetMonthlyOrder() ([]dto.MonthlyOrder, error) {
	rows, err := r.db.Query(utils.VIEW_ORDER_MONTH)
	if err != nil {
		log.Printf("Gagal eksekusi queri %s", err)
		return nil, err
	}
	defer rows.Close()

	var summaries []dto.MonthlyOrder

	for rows.Next() {
		var summary dto.MonthlyOrder
		if err := rows.Scan(&summary.Month, &summary.TotalOrder); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		summaries = append(summaries, summary)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		return nil, err
	}

	return summaries, nil
}

func (r *orderRepository) GetMonthlyOrderCustomer() ([]dto.MonthlyCustomerOrder, error) {
	rows, err := r.db.Query(utils.VIEW_ORDER_CS_MONTH)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var summaries []dto.MonthlyCustomerOrder
	for rows.Next() {
		var summary dto.MonthlyCustomerOrder
		if err := rows.Scan(&summary.CustomerName, &summary.Month, &summary.TotalOrders); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		summaries = append(summaries, summary)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		return nil, err
	}

	return summaries, nil
}
