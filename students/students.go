package students

import (
	"ISA_DA5/utilities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterStudent(c *gin.Context) {
	var body map[string]string
	if err := c.BindJSON(&body); err != nil {
		c.JSON(401, gin.H{"status": "invalid input", "err": err.Error()})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(body["password"]), bcrypt.DefaultCost)
	newStudent := Student{Name: body["name"], Password: password}
	file, _ := ioutil.ReadFile("./students/studentDatabase.json")
	var datas []Student
	json.Unmarshal(file, &datas)
	datas = append(datas, newStudent)
	data, _ := json.MarshalIndent(datas, "", " ")
	ioutil.WriteFile("./students/studentDatabase.json", data, os.ModePerm)
	c.JSON(200, gin.H{"status": "successful"})
}

func StudentLogin(c *gin.Context) {
	var body map[string]string = make(map[string]string)
	//body = c.Request.URL.Query()
	if err := c.ShouldBindQuery(&body); err != nil {
		c.JSON(401, gin.H{"status": "invalid input", "err": err.Error()})
		return
	}

	file, _ := ioutil.ReadFile("./students/studentDatabase.json")
	var datas []Student
	json.Unmarshal(file, &datas)
	var password []byte = nil
	for _, user := range datas {
		if user.Name == body["name"] {
			password = user.Password
			break
		}
	}
	if password == nil {
		c.JSON(401, gin.H{"status": "user not found"})
		return
	}
	if err := bcrypt.CompareHashAndPassword(password, []byte(body["password"])); err != nil {
		c.JSON(401, gin.H{"status": "invalid password"})
		return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    body["name"] + ":student",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error", "message": err.Error()})
		return
	}
	c.SetCookie("jwt", "", -1, "/teacher", "", false, true)
	c.SetCookie("jwt", token, 60*60*24, "/student", "", false, true)
	//c.SetCookie("cookieName", "name", 10, "/", "", false, false)

	c.JSON(200, gin.H{"status": "success"})
}

func GetUser(c *gin.Context) {
	issuer := strings.Split(utilities.GetUserName(c), ":")[0]
	c.JSON(200, gin.H{"message": "authenticated User Successfully", "user": issuer})
}

func AuthenticateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {

		var cookie string
		var er error
		if cookie, er = c.Cookie("jwt"); er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": er.Error()})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}
		issuer := token.Claims.(*jwt.StandardClaims).Issuer
		splitIssuer := strings.Split(issuer, ":")
		if splitIssuer[1] != "student" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "you are not a student"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func Logout(c *gin.Context) {
	fmt.Println("here")
	c.SetCookie("jwt", "", -1, "/student", "", false, true)
	c.JSON(200, gin.H{"message": "successful"})
}
