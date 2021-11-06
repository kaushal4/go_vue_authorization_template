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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/student", students.AuthenticateStudent(), students.GetUser)
	router.GET("/student/logout", students.AuthenticateStudent(), students.Logout)
	router.GET("/student/course", students.AuthenticateStudent(), courses.GetCourse)
	router.GET("/student/login", students.StudentLogin)
	router.POST("/student", students.RegisterStudent)
	router.GET("/teacher", teachers.AuthenticateTeacher(), teachers.GetUser)
	router.GET("/teacher/logout", teachers.AuthenticateTeacher(), teachers.Logout)
	router.GET("/teacher/login", teachers.TeacherLogin)
	router.POST("/teacher", teachers.RegisterTeacher)
	router.GET("/teacher/course", teachers.AuthenticateTeacher(), courses.GetCourse)
	router.POST("/teacher/course", teachers.AuthenticateTeacher(), courses.RegisterCourse)
	router.PATCH("/teacher/course", teachers.AuthenticateTeacher(), courses.EditCourse)
	router.GET("/", hello)
	router.Run("localhost:8000")
	fmt.Println("Server Started!!")
}
