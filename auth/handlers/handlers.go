package handlers

import (
	"auth/db"
	"auth/model"
	"auth/repositories"
	"auth/services"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	userRepo *repositories.UserRepo
}

func NewHandler(userRepo *repositories.UserRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {

	healthStatusStr := string("server OK, ")

	dbHealthOk := db.Health()
	log.Println("healthStatus", dbHealthOk)
	if dbHealthOk {
		healthStatusStr += " mongo OK"
	} else {
		healthStatusStr += " mongo not OK"
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(healthStatusStr))
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("registerHandler logic")

	userData := model.User{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		log.Println("err decoding request body")
	}

	userCreateError := h.userRepo.CreateUser(db.Ctx, &userData)
	if userCreateError != nil {
		log.Println("userCreateError", userCreateError)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User exists"))
		return
	}

	services.SendActivationMail(userData.Username)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created"))
}

func (h *Handler) ActiavteAccountHandler(w http.ResponseWriter, r *http.Request) {

	var requestBody map[string]string

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	email := requestBody["email"]
	code := requestBody["code"]

	if err != nil {
		log.Println("err decoding reques body")
	}

	log.Println("email:", email)
	log.Println("code:", code)

	storedCode := db.Get(email)
	if storedCode == code {
		updateJsonString := `{"activated":true}`

		err := h.userRepo.UpdateUser(email, updateJsonString)
		if err != nil {
			log.Println("err updating user")
			return
		}
		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]string)
		resp["message"] = "User verified"

		tokenStr, expires := services.GenerateJWTToken(email)
		log.Println("GENERATED TOKEN:::::::::", tokenStr)

		_ = tokenStr
		_ = expires

		http.SetCookie(
			w, &http.Cookie{
				Domain:  "127.0.0.1",
				Name:    "token",
				Value:   tokenStr,
				Expires: expires,
			})

		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)

	}
}

func (Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("login Handler")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("je"))
}

func (h *Handler) Protected(w http.ResponseWriter, r *http.Request) {
	log.Println("protected endpoint")
	cookie, err := r.Cookie("token")

	log.Println("cookie is ......", cookie)
	log.Println("cookie err", err)
}
