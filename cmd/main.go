package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/go-restful-api-example/internal/api"
	"github.com/pancudaniel7/go-restful-api-example/internal/controller"
	repository "github.com/pancudaniel7/go-restful-api-example/internal/repository"
	service "github.com/pancudaniel7/go-restful-api-example/internal/service"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	db     *gorm.DB
	router *gin.Engine

	bookRepository  api.BookRepository
	storeRepository api.StoreRepository

	bookService  api.BookService
	storeService api.StoreService

	healthController api.HealthController
	bookController   api.BookController
	storeController  api.StoreController
)

func main() {
	router = gin.Default()
	readProperties()
	initDatabase()

	initRepositories()
	initServices()
	initControllers()

	registerRoutes(router)
	port := viper.GetInt("server.port")

	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}

func initRepositories() {
	bookRepository = repository.NewBookRepositoryImpl(db)
	storeRepository = repository.NewStoreRepository(db)
}

func readProperties() {
	configName := os.Getenv("CONFIG_NAME")
	if configName == "" {
		configName = "local"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs/")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func registerRoutes(router *gin.Engine) {
	router.GET("/health", healthController.Health)

	router.POST("/books", bookController.AddBook)
	router.PUT("/books", bookController.UpdateBook)
	router.DELETE("/books/:id", bookController.DeleteBook)
	router.GET("/books", bookController.GetBooks)
	router.GET("/books/:id", bookController.GetBook)

	router.POST("/stores", storeController.AddStore)
	router.PUT("/stores", storeController.UpdateStore)
	router.DELETE("/stores/:id", storeController.DeleteStore)
	router.GET("/stores", storeController.GetStores)
	router.GET("/stores/:id", storeController.GetStore)
}

func initDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.dbname"),
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
}

func initServices() {
	storeService = service.NewStoreServiceImpl(storeRepository)
	bookService = service.NewBookServiceImpl(bookRepository)
}

func initControllers() {
	healthController = controller.NewHealthController()
	bookController = controller.NewBookController(bookService)
	storeController = controller.NewStoreController(storeService)
}
