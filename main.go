package main

import (
	_database "assigment-2-usamah/database"
	_routes "assigment-2-usamah/delivery/routes"
	_orderRepository "assigment-2-usamah/repository/order"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db := _database.GetConnection()
	defer db.Close()
	fmt.Println("Successfully connected to database")

	orderRepository := _orderRepository.NewOrderRepository(db)

	router := _routes.Routes(
		orderRepository,
	)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
