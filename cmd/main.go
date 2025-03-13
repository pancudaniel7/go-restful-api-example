package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/go-restful-api-example/internal/controller"
	service "github.com/pancudaniel7/go-restful-api-example/internal/service"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	databaseClient *gorm.DB
	router         *gin.Engine

	bookService  service.BookService
	storeService service.StoreService

	healthController *controller.HealthController
	bookController   *controller.BookController
	storeController  *controller.StoreController
)

func main() {
	readProperties()
	initDatabase()
	initRouter()
	initServices()
	initControllers()

	registerRoutes()
	port := viper.GetInt("server.port")

	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
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

func registerRoutes() {
	storeController.RegisterRoutes(router)
	bookController.RegisterRoutes(router)
	healthController.RegisterRoutes(router)
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
	databaseClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
}

func initRouter() {
	router = gin.Default()
}

func initServices() {
	storeService = service.NewStoreService(databaseClient)
	bookService = service.NewBookService(databaseClient)
}

func initControllers() {
	healthController = controller.NewHealthController()
	bookController = controller.NewBookController(bookService)
	storeController = controller.NewStoreController(storeService)
}
