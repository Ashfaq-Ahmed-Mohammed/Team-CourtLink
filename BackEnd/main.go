package main

import (
	"BackEnd/Court"
	"BackEnd/Customer"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const publickey = `-----BEGIN CERTIFICATE-----
MIIDHTCCAgWgAwIBAgIJDMeoXNyQNkpkMA0GCSqGSIb3DQEBCwUAMCwxKjAoBgNVBAMTIWRldi03Z3Bwamk4djNiZGJzajZrLnVzLmF1dGgwLmNvbTAeFw0yNTAyMDcwNTI5MDRaFw0zODEwMTcwNTI5MDRaMCwxKjAoBgNVBAMTIWRldi03Z3Bwamk4djNiZGJzajZrLnVzLmF1dGgwLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALGFXxGQB92b/Q1qhwJFDh1m7u/ui0xSfXsLks16KmjSAMkyNm0pn1t2dZcttSxR1tLE2Zrk0HlN3uKXLCQ3GvlPfhc/L51vQ3fHVrubcbwwISgqzcdoG+GOpTJNWE1g1wtmMP1HrF/YEkeLykJiEzKt+X4ZcKiF4PB6zzUdyWwNx81gcTzUzSlQb+WFsKNOxJiGDy1N9RF6khn4ElWEVF+g0DUBSdpIVzSEFpEwLO8zNmarTSrP0p4SjtPF18mc8r/4Ce9niIr3vqWsu2qZg7O0aQQ+AqWhePkqxuhRs4zfZLsQ58nhiEd6EuowjaifLNWPnmlDWJhZq8hBcOrIoRUCAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUMdsDNggQEIJkPlEEkTq5lZJ7HtMwDgYDVR0PAQH/BAQDAgKEMA0GCSqGSIb3DQEBCwUAA4IBAQCq1ZlZt37mkrTcOiH7PgLpljUvFEl1RN6DBrGDStRPif5mWZYA2SGXFlMhycgijbtJ9jbiHyhyoNkF7Sq60K35q6YHJ8i0hhWxFomFFEKQe2RSwOM7RmQM2TwRHQJ2QU4PswQjm/JAI9HMBC6e/goQOSL0+kDYuZstlNvVSZ8/Gq0XUt9JP1dSBZK8tZvX4kQ5AOHar6Kxpm24pOsTSJzI+m8AIEID/ojYtwTfhsszYbWNAXtNz7BKxTDEPUEdrwcuf16iYjZMBY9zuCjrGnC/OqoVze+f1mkpLYGKZDOyHC1AOjHQyz817AE3mfH0O7JpEu2ePi6YAEDduQ3sMr5a
-----END CERTIFICATE-----`

func validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			block, _ := pem.Decode([]byte(publickey))
			if block == nil {
				return nil, fmt.Errorf("failed to parse PEM block containing the public key")
			}

			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse certificate: %v", err)
			}

			return cert.PublicKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/Customer", Customer.CreateCustomer).Methods("POST")
	r.HandleFunc("/GetCourt", Court.GetCourt).Methods("GET")
	newroute := r.PathPrefix("/api").Subrouter()
	newroute.Use(validateToken)
	newroute.HandleFunc("/CreateCustomer", Customer.CreateCustomer).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
