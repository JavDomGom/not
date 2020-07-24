package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/jwt"
	"github.com/JavierDominguezGomez/not/models"
)

/*Login make login. */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid username or password: "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "The email is required.", 400)
		return
	}

	document, exist := db.LoginAttemp(t.Email, t.Password)
	if !exist {
		http.Error(w, "Invalid username or password.", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error occurred while trying to generate the token.: "+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	// Make cookie from backend.
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
