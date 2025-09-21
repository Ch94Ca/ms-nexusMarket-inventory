package main

import (
	"log"
	"os"

	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/app/handler"
	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/infra/postgresrepository"
	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/Ch94Ca/ms-nexusMarket-inventory/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ms-nexus-inventory API
// @version 1.0
// @description This is a API for a inventory management using Gin framework.
// @host localhost:8090
// @BasePath /
func main() {
	db := setupDatabase()

	categoryRepo := postgresrepository.NewCategoryRepositoryPostgres(db)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)

	r := gin.Default()
	categoryHandler := handler.NewCategoryHandler(categoryUC)

	setupRoutes(r, categoryHandler)

	if err := r.Run(":8090"); err != nil {
		log.Panic(err)
	}
}

func setupRoutes(r *gin.Engine, categoryHandler *handler.CategoryHandler) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", handler.HealthCheckHandler)

	setupCategoryRoutes(r, categoryHandler)
}

func setupCategoryRoutes(r *gin.Engine, categoryHandler *handler.CategoryHandler) {
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories", categoryHandler.ListCategories)
	r.GET("/categories/:id", categoryHandler.GetCategoryByID)
	r.PATCH("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)
}

func setupDatabase() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to the database: ", err)
	}

	return db
}
