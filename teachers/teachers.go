package teachers

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

func RegisterTeacher(c *gin.Context) {
	var body map[string]string
	if err := c.BindJSON(&body); err != nil {
		c.JSON(401, gin.H{"status": "invalid input", "err": err.Error()})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(body["password"]), bcrypt.DefaultCost)
	newTeacher := Teacher{Name: body["name"], Password: password}
	file, _ := ioutil.ReadFile("./teachers/teacherDatabase.json")
	var datas []Teacher
	json.Unmarshal(file, &datas)
	datas = append(datas, newTeacher)
	data, _ := json.MarshalIndent(datas, "", " ")
	ioutil.WriteFile("./teachers/teacherDatabase.json", data, os.ModePerm)
	c.JSON(200, gin.H{"status": "successful"})
}

func TeacherLogin(c *gin.Context) {
	var body map[string]string = make(map[string]string)
	//body = c.Request.URL.Query()
	if err := c.BindQuery(&body); err != nil {
		c.JSON(401, gin.H{"status": "invalid input", "err": err.Error()})
		return
	}
	file, _ := ioutil.ReadFile("./teachers/teacherDatabase.json")
	var datas []Teacher
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
		Issuer:    body["name"] + ":teacher",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte("secret"))

	fmt.Println("token :" + token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error", "message": err.Error()})
		return
	}

	c.SetCookie("jwt", token, int(time.Hour)*24, "/teacher", "", false, true)

	c.JSON(200, gin.H{"message": "success"})
}

func GetUser(c *gin.Context) {
	issuer := strings.Split(utilities.GetUserName(c), ":")[0]
	c.JSON(http.StatusOK, gin.H{"message": "authenticated User Successfully", "user": issuer})
}

func AuthenticateTeacher() gin.HandlerFunc {
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
		if splitIssuer[1] != "teacher" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "you are not a teacher"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ReadTeachers() ([]Teacher, error) {
	file, _ := ioutil.ReadFile("./teachers/teacherDatabase.json")
	var datas []Teacher
	if err := json.Unmarshal(file, &datas); err != nil {
		return nil, err
	}
	return datas, nil
}
func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/teacher", "", false, true)
	c.JSON(200, gin.H{"message": "successful"})
}
