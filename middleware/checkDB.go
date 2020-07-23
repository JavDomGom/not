package middleware

import (
	"net/http"

	"github.com/JavierDominguezGomez/not/db"
)

/*CheckDB middleware to check database status.*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Lost connection to the database.", 500)
		}
		next.ServeHTTP(w, r)
	}
}
