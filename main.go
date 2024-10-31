package main

import (
	"fmt"
	"log"

	"github.com/golang-database/database"
	"github.com/golang-database/service"
)

func main() {
	db, err := database.ConnetDatabase()
	if err != nil {
		log.Fatal("Gagal connet database!!!!", err)
	} else {
		log.Println("Success connet to database !!!")
	}
	defer db.Close()

	// service.InputDataAddress(db, "Rajeg", "Tangerang", "Banten")
	// service.InputDataCustomer(db, "dewi", 2, "active")
	// err = service.InputDataDriver(db, "julpi", 2)
	// if err != nil {
	// 	fmt.Println("error :", err)
	// }

	err = service.InputDataOrder(db, 2, 3, "2024-10-19 01:00:00", 20000000)
	if err != nil {
		fmt.Println("err :", err)
	}

	// orderRepo := repository.NewOrderRepository(db)
	// orderService := service.NewOrderService(orderRepo)

	// orderan terbanyak setiap bulan
	// summaries, err := orderService.GetMonthlyOrder()
	// if err != nil {
	// 	log.Fatalf("Error getting monthly order summary: %v", err)
	// }

	// for _, summary := range summaries {
	// 	fmt.Printf("Bulan: %s, Total Pesanan: %d\n", summary.Month, summary.TotalOrder)
	// }

	// customer paling banyak order
	// sum, err := orderService.GetMonthlyCustomerOrder()
	// if err != nil {
	// 	log.Fatalf("Error getting monthly customer order summary: %v", err)
	// }

	// for _, summary := range sum {
	// 	fmt.Printf("Pelanggan: %s, Bulan: %s, Total Pesanan: %d\n", summary.CustomerName, summary.Month.Format("2006-01"), summary.TotalOrders)
	// }

	//
}
