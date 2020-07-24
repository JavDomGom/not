package routers

import (
	"errors"
	"strings"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email value used on all endpoints.*/
var Email string

/*UserID Returned ID from the model, which will be used on all endpoints. */
var UserID string

/*ProcessToken Function to manage the token. */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("NwodSeogNworbYbbob")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		_, found, _ := db.CheckIfUserAlreadyExists(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
