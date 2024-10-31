package dto

import "time"

type MonthlyCustomerOrder struct {
	CustomerName string
	Month        time.Time
	TotalOrders  int
}
