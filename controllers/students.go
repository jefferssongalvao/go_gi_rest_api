package controllers

import (
	"go-gin-rest-api/database"
	"go-gin-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllStudents(ctx *gin.Context) {
	var students []models.Student
	if name := ctx.Query("name"); name != "" {
		students = searchStudentByName(name)
		if len(students) == 0 {
			ctx.JSON(http.StatusNoContent, nil)
			return
		}
		ctx.JSON(http.StatusOK, students)
		return
	}

	database.DB.Find(&students)
	ctx.JSON(http.StatusOK, students)
}

func GetStudent(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Student Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func CreateStudent(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	ctx.JSON(http.StatusCreated, student)
}

func DeleteStudent(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")

	database.DB.Delete(&student, id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "student has been deleted",
	})
}

func UpdateStudent(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	var student models.Student
	database.DB.First(&student, id)

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Save(&student)

	ctx.JSON(http.StatusOK, student)
}

func searchStudentByName(name string) []models.Student {
	var students []models.Student
	database.DB.Where("name ilike ?", "%"+name+"%").Find(&students)
	return students
}
