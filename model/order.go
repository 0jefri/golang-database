package model

type Order struct {
	ID         int
	CustomerId int
	DriverId   int
	OrderDate  string
	Price      int
}
