package main

import (
	"ISA_DA5/courses"
	"ISA_DA5/students"
	"ISA_DA5/teachers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.String(200, "Hello World 2!!")
}

func main() {
	router := gin.Default()
	router.GET("/student", students.AuthenticateStudent(), students.GetUser)
	router.GET("/student/login", students.StudentLogin)
	router.POST("/student", students.RegisterStudent)
	router.GET("/teacher", teachers.AuthenticateTeacher(), teachers.GetUser)
	router.GET("/teacher/login", teachers.TeacherLogin)
	router.POST("/teacher", teachers.RegisterTeacher)
	router.GET("/teacher/course", teachers.AuthenticateTeacher(), courses.GetCourse)
	router.POST("/teacher/course", teachers.AuthenticateTeacher(), courses.RegisterCourse)
	router.PATCH("/teacher/course", teachers.AuthenticateTeacher(), courses.EditCourse)
	router.GET("/", hello)
	router.Run("localhost:8080")
	fmt.Println("Server Started!!")
}
