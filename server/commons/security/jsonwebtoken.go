package security

import (
	"os"
	"strconv"
	"time"

	"github.com/devnica/EasyStore/exceptions"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userId string, roles []map[string]interface{}) string {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	exceptions.PanicLogging(err)

	claims := jwt.MapClaims{
		"UserId": userId,
		"roles":  roles,
		"exp":    time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exceptions.PanicLogging(err)

	return tokenSigned
}
