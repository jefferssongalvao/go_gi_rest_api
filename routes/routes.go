package routes

import (
	"go-gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

const (
	BaseUrlStudent = "/students"
)

func HandleRequests() {
	router := gin.Default()
	router.GET(BaseUrlStudent, controllers.AllStudents)
	router.GET(BaseUrlStudent+"/:id", controllers.GetStudent)
	router.POST(BaseUrlStudent, controllers.CreateStudent)
	router.DELETE(BaseUrlStudent+"/:id", controllers.DeleteStudent)
	router.PATCH(BaseUrlStudent+"/:id", controllers.UpdateStudent)
	router.Run()
}
