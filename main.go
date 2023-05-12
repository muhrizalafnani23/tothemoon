package main

import (
	"flight-example-api/flight"
	"flight-example-api/login"
	"flight-example-api/middleware"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/flight", middleware.Authenticate(http.HandlerFunc(flight.FlightHandler)))
	http.HandleFunc("/login", login.Login)

	fmt.Println("server running on port :8088")
	http.ListenAndServe(":8088", nil)
}
