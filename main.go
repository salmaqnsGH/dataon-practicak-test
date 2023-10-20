package main

import (
	"dataon/company"
	"dataon/division"
	"dataon/executiveCommittee"
	"dataon/handler"
	subdivision "dataon/subdivision"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=secret dbname=dataon port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("koneksi database berhasil!")

	// company
	companyRepository := company.NewRepository(db)
	companyService := company.NewService(companyRepository)
	companyHandler := handler.NewCompanyHandler(companyService)

	// executive committee
	executiveCommitteeRepository := executiveCommittee.NewRepository(db)
	executiveCommitteeService := executiveCommittee.NewService(executiveCommitteeRepository)
	executiveCommitteeHandler := handler.NewExecutiveCommitteeHandler(executiveCommitteeService)

	// division
	divisionRepository := division.NewRepository(db)
	divisionService := division.NewService(divisionRepository)
	divisionHandler := handler.NewDivisionHandler(divisionService)

	// sub division
	subDivisionRepository := subdivision.NewRepository(db)
	subDivisionService := subdivision.NewService(subDivisionRepository)
	subDivisionHandler := handler.NewSubDivisionHandler(subDivisionService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.GET("/companies", companyHandler.GetCompanies)
	api.POST("/companies", companyHandler.CreateCompany)

	api.GET("/executive-committies", executiveCommitteeHandler.GetExecutiveCommitties)
	api.POST("/executive-committies/:id", executiveCommitteeHandler.CreateExecutiveCommittee)

	api.GET("/divisions", divisionHandler.GetDivisions)
	api.POST("/divisions/:id", divisionHandler.CreateDivision)

	api.GET("/sub-divisions", subDivisionHandler.GetSubDivisions)
	api.POST("/sub-divisions/:id", subDivisionHandler.CreateSubDivision)

	router.Run()
}
