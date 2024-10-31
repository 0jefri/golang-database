package service

import (
	"database/sql"
	"fmt"

	"github.com/golang-database/model"
	"github.com/golang-database/model/dto"
	"github.com/golang-database/repository"
)

type orderService struct {
	orderRepo repository.OrderRepository
}

func InputDataOrder(db *sql.DB, customer_id int, driver_id int, order_date string, price int) error {
	if customer_id == 0 || driver_id == 0 || order_date == "" || price == 0 {
		return fmt.Errorf("customer id, driver id, order date, price is required")
	}

	orderRepo := repository.NewOrderRepository(db)
	order := model.Order{
		CustomerId: customer_id,
		DriverId:   driver_id,
		OrderDate:  order_date,
		Price:      price,
	}

	err := orderRepo.Create(&order)
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	fmt.Println("Success Create data orders")
	return nil
}

func (s orderService) GetMonthlyOrder() ([]dto.MonthlyOrder, error) {
	result, err := s.orderRepo.GetMonthlyOrder()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *orderService) GetMonthlyCustomerOrder() ([]dto.MonthlyCustomerOrder, error) {
	return s.orderRepo.GetMonthlyOrderCustomer()
}

func NewOrderService(orderRepo repository.OrderRepository) *orderService {
	return &orderService{orderRepo: orderRepo}
}
