package main

import (
	"loan/db"
	"loan/handlers"
	"loan/middlewares"
	"loan/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("loan service")

	mongoClient := db.Client
	var db = mongoClient.Database("LOANS_DB")

	loanRepo := repositories.NewLoanRepo(db)
	handler := handlers.NewHandler(loanRepo)

	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/api/loan/health", handler.HealthHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/loan/apply", middlewares.Auth(handler.Apply)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/loan/status", middlewares.Auth(handler.Status)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/loan/loans", handler.Loans).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/loan/approve", handler.ApproveLoan).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/loan/reject", handler.RejectLoan).Methods("POST", "OPTIONS")

	http.ListenAndServe(":8020", router)
}
