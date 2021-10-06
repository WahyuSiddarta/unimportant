package helper

import (
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//JWTCustomClaims : JWT custom claim struct
type JWTCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

//GetHMAC :
func GetHMAC(c echo.Context) string {
	Authentication := c.Request().Header.Get("Authentication")
	a := strings.TrimSpace(Authentication)
	auth := strings.Split(a, " ")
	if len(auth) != 2 || auth[0] != "hmac" {
		return ""
	}
	return auth[1]
}

//GetJWTEmail : Get email from JWT
func GetJWTEmail(c echo.Context) string {
	user := c.Get("user")
	if user == nil {
		return ""
	}
	claims := user.(*jwt.Token).Claims.(*JWTCustomClaims)
	return claims.Audience
}

//GetJWTUserID : Get email from JWT
func GetJWTUserID(c echo.Context) string {
	user := c.Get("user")
	if user == nil {
		return ""
	}
	claims := user.(*jwt.Token).Claims.(*JWTCustomClaims)
	return claims.UserID
}

//GetJWT : get JWT structure
func GetJWT(c echo.Context) *JWTCustomClaims {
	user := c.Get("user")
	if user == nil {
		return nil
	}
	claims := user.(*jwt.Token).Claims.(*JWTCustomClaims)
	return claims
}

//CreateJWT :
func CreateJWT(LoginID string, UserID string, subject, secret string) (int64, string, error) {
	atTime := time.Now().Add(time.Hour * 72).UTC().Unix()
	claims := &JWTCustomClaims{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: atTime,
			Issuer:    "Ex-API",
			Subject:   subject,
			Audience:  LoginID,
			NotBefore: time.Now().UTC().Unix() - 1,
			IssuedAt:  time.Now().UTC().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response.
	tok, err := token.SignedString([]byte(secret))
	if err != nil {
		return 0, "", err
	}
	return atTime, tok, nil
}
