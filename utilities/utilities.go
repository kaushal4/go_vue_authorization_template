package utilities

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUserName(c *gin.Context) string {
	var cookie string
	var er error
	if cookie, er = c.Cookie("jwt"); er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": er.Error()})
		c.Abort()
		return ""
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
		c.Abort()
		return ""
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer
}
