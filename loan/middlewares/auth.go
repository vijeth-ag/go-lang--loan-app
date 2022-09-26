package middlewares

import (
	"loan/models"
	"loan/services"
	"log"
	"net/http"
	"strings"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	log.Println("......AUTH MIDDLEWARE.......")
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Println("err retrieving cookie from request")
		}
		log.Println("auth middleware ....cookie", cookie)
		tkn := strings.Split(cookie.String(), "=")
		log.Println("tkn", tkn[1])

		authRqst := models.AuthRequest{
			JWTToken: tkn[1],
		}

		valid := services.ValidateJwt(authRqst)
		log.Println("valid", valid)
		if valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden request", http.StatusForbidden)
		}
	}
}
