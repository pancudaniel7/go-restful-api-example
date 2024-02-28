package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/go-restful-api-example/internal/controller"
	service "github.com/pancudaniel7/go-restful-api-example/internal/service"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("internal/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.dbname"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	storeService := service.NewStoreService(db)
	storeController := controller.NewStoreController(storeService)

	// Instantiate the BookService and BookController
	bookService := service.NewBookService(db)
	bookController := controller.NewBookController(bookService)

	router := gin.Default()
	storeController.RegisterRoutes(router)
	bookController.RegisterRoutes(router)

	port := viper.GetInt("server.port")
	router.Run(fmt.Sprintf(":%d", port))
}
