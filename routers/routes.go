package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jtw-second/controller"
)

func SetupRoutes(studentsController *controller.StudentController) *gin.Engine {
	r := gin.Default()
	s := r.Group("students")
	{
		s.POST("/", studentsController.CreateStudent)
		s.GET("/:id", studentsController.GetStudentByID)
		s.PUT("/", studentsController.UpdateStudent)
		s.DELETE("/:id", studentsController.DeleteStudent)
		s.GET("/", studentsController.GetAllStudents)
	}
	return r
}
