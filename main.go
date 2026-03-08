package main

import (
	"github.com/sidz111/jtw-second/controller"
	dbconfig "github.com/sidz111/jtw-second/dbConfig"
	"github.com/sidz111/jtw-second/models"
	"github.com/sidz111/jtw-second/repository"
	"github.com/sidz111/jtw-second/routers"
	"github.com/sidz111/jtw-second/service"
)

func main() {
	if err := dbconfig.ConnectDatabase(); err != nil {
		panic(err)
	}
	dbconfig.DB.AutoMigrate(&models.Student{})
	// r := gin.Default()
	repo := repository.NewStudentRepository(dbconfig.DB)
	serv := service.NewStudentService(repo)
	controller := controller.NewStudentController(serv)
	r := routers.SetupRoutes(controller)
	r.Run(":8080")
}
