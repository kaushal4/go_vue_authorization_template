package teachers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	var body map[string]string
	if err := c.BindJSON(&body); err != nil {
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
		Issuer:    body["name"],
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte("secret"))

	fmt.Println("token :" + token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error", "message": err.Error()})
		return
	}

	c.SetCookie("jwt", token, int(time.Hour)*24, "/teacher", "localhost", true, true)

	c.JSON(200, gin.H{"token": token})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"message": "authenticated User Successfully"})
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
		_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
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
