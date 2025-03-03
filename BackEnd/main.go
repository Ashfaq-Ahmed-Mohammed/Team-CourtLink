package main

// @title Court Booking API
// @version 1.0
// @description API for managing court bookings
// @host localhost:8080
// @BasePath /

import (
	"BackEnd/Bookings"
	"BackEnd/Court"
	"BackEnd/Customer"
	_ "BackEnd/docs"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()

	// Middleware to handle CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all domains temporarily
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Handle OPTIONS requests globally
	r.Use(mux.CORSMethodMiddleware(r))

	// Explicitly handle OPTIONS requests
	r.HandleFunc("/getCourts", Court.GetCourt).Methods("GET", "OPTIONS")

	// Other API routes
	r.HandleFunc("/Customer", Customer.CreateCustomer).Methods("POST", "OPTIONS")
	r.HandleFunc("/UpdateCourtSlot", Court.UpdateCourtSlot).Methods("PUT", "OPTIONS")
	r.HandleFunc("/CreateBooking", Bookings.CreateBooking).Methods("POST", "OPTIONS")

	// Serve Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	handler := corsHandler.Handler(r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
