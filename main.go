package main

import (
	"dataon/company"
	"dataon/handler"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=secret dbname=organisational_structure port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("koneksi database berhasil!")

	companyRepository := company.NewRepository(db)

	companyService := company.NewService(companyRepository)

	companyHandler := handler.NewCompanyHandler(companyService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.GET("/companies", companyHandler.GetCompanies)

	router.Run()
}
