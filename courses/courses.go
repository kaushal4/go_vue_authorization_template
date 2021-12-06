package courses

import (
	"ISA_DA5/teachers"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func checkTeacherName(teacherName string) bool {
	teachers, _ := teachers.ReadTeachers()
	for _, teacher := range teachers {
		if teacher.Name == teacherName {
			return true
		}
	}
	return false

}

func readCourses() ([]Course, error) {
	file, _ := ioutil.ReadFile("./courses/courseDatabase.json")
	var datas []Course
	if err := json.Unmarshal(file, &datas); err != nil {
		return nil, err
	}
	return datas, nil
}

func writeCourse(datas []Course) {
	data, _ := json.MarshalIndent(datas, "", " ")
	ioutil.WriteFile("./courses/courseDatabase.json", data, os.ModePerm)
}

func RegisterCourse(c *gin.Context) {
	var body Course
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "invalid input", "err": err.Error()})
		return
	}
	fmt.Println(body)
	if !(len(body.Name) > 0) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Course needs a name"})
		return
	}
	//below condition is kinda redundent
	if !checkTeacherName(body.Teacher) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No such teacher found"})
		return
	}

	courses, _ := readCourses()
	courses = append(courses, body)
	writeCourse(courses)
	c.JSON(http.StatusOK, gin.H{"message": "successfully added"})
}

func GetCourse(c *gin.Context) {
	var params map[string]string = make(map[string]string)

	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if _, ok := params["name"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No name"})
		return
	}
	courses, _ := readCourses()
	var course Course
	var flag bool = false
	for _, element := range courses {
		if element.Name == params["name"] {
			course = element
			flag = true
			break
		}
	}
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No such course"})
		return
	}
	c.JSON(http.StatusAccepted, course)
}

func checkTeacher(c *gin.Context, name string) bool {
	var cookie string
	var er error
	if cookie, er = c.Cookie("jwt"); er != nil {
		return false
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return false
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.VerifyIssuer(name+":teacher", true)

}

func EditCourse(c *gin.Context) {

	var body map[string]string
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var course Course
	courses, _ := readCourses()
	var index int = -1
	for i, element := range courses {
		if element.Name == body["name"] {
			course = element
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No such course"})
		return
	}
	if !checkTeacher(c, course.Teacher) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not Authorised"})
		return
	}
	length, _ := strconv.Atoi(body["materialNo"])
	if !(len(courses[index].Material) > length) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invlaid material"})
		return
	}
	courses[index].Material[length] = body["material"]
	writeCourse(courses)
	c.JSON(200, gin.H{"message": "material updated"})

}

func handleError(err error, c *gin.Context) {
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func AddFile(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	var body map[string]string = make(map[string]string)
	for key, value := range c.Request.PostForm {
		body[key] = value[0]
		fmt.Println(key, value)
	}
	var course Course
	courses, _ := readCourses()
	var index int = -1
	for i, element := range courses {
		if element.Name == body["name"] {
			course = element
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No such course"})
		return
	}
	if !checkTeacher(c, course.Teacher) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not Authorised"})
		return
	}

	handleError(err, c)
	if err != nil {
		return
	}
	file, header, err := c.Request.FormFile("file")
	handleError(err, c)
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)
	tempFile, err := ioutil.TempFile("courses\\uploadedFiles", "upload-*.pdf")
	handleError(err, c)
	if err != nil {
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	handleError(err, c)
	if err != nil {
		return
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	course.Files = append(course.Files, tempFile.Name())

	courses[index] = course

	writeCourse(courses)
	// return that we have successfully uploaded our file!
	c.JSON(http.StatusAccepted, gin.H{"message": "Uploaded Successfully"})
}

func DownloadFile(c *gin.Context) {
	var body map[string]string
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if _, err := os.Stat(body["path"]); errors.Is(err, os.ErrNotExist) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File does not exits"})
		return
	}
	name := body["path"][14:]
	//c.Header("Content-Disposition", "attachment; filename="+name)
	//c.Header("Content-Type", "application/octet-stream")
	c.FileAttachment(body["path"], name)
}
