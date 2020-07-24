package db

import (
	"github.com/JavierDominguezGomez/not/models"
	"golang.org/x/crypto/bcrypt"
)

/*LoginAttemp Perform the login check in the database. */
func LoginAttemp(email string, password string) (models.User, bool) {
	user, found, _ := CheckIfUserAlreadyExists(email)
	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
