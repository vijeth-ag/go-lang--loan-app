package main

import (
	"log"
	"net"
	"net/http"

	"auth/db"
	"auth/handlers"
	"auth/repositories"
	"auth/services"

	"auth/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
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

	go startGrpcServer()

	mongoClient := db.Client

	var db = mongoClient.Database("USERS_DB")

	userRepo := repositories.NewUserRepo(db)
	handler := handlers.NewHandler(userRepo)

	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/api/health", handler.HealthHandler)
	router.HandleFunc("/api/register", handler.RegisterHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/verify", handler.ActiavteAccountHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/login", handler.LoginHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/test", handler.Protected).Methods("GET", "OPTIONS")

	log.Println("starting")
	http.ListenAndServe(":9000", router)

}

func startGrpcServer() {
	log.Println("starting grpc")
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Println("err listening on :9001")
		return
	}
	var s = services.MyGrpcServer{}
	grpcServer := grpc.NewServer()

	proto.RegisterAuthServiceServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("err starting grpc server")
		return
	}
	log.Println("grpc up")
}
