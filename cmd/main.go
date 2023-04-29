package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meziaris/devstore/internal/app/controller"
	"github.com/meziaris/devstore/internal/app/repository"
	"github.com/meziaris/devstore/internal/app/service"
	"github.com/meziaris/devstore/internal/pkg/config"
	"github.com/meziaris/devstore/internal/pkg/db"
	"github.com/meziaris/devstore/internal/pkg/middleware"
	"github.com/sirupsen/logrus"
)

var cfg config.Config
var DBConn *sqlx.DB

func init() {
	// read config
	loadConfig, err := config.LoadConfig("./")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = loadConfig

	// connect database
	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic(err)
	}

	DBConn = db
}

func main() {
	r := gin.New()

	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	// repository
	categoryRepository := repository.NewCategoryRepository(DBConn)
	productRepository := repository.NewProducRepository(DBConn)

	// service
	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository, categoryRepository)

	// controller
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)

	r.GET("/category", categoryController.BrowseCategory)
	r.POST("/category", categoryController.CreateCategory)
	r.POST("/category/:id", categoryController.DetailCategory)
	r.PATCH("/category/:id", categoryController.UpdateCategory)
	r.DELETE("/category/:id", categoryController.DeleteCategory)

	r.GET("/product", productController.BrowseProduct)
	r.POST("/product", productController.CreateProduct)
	r.GET("/product/:id", productController.DetailProduct)
	r.PATCH("/product/:id", productController.UpdateProduct)
	r.DELETE("/product/:id", productController.DeleteProduct)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	if err := r.Run(appPort); err != nil {
		logrus.Fatal(fmt.Errorf("cann't start app: %w", err))
	}
}
