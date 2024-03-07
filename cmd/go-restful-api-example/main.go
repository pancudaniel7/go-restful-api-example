package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/go-restful-api-example/api/controller"
	service "github.com/pancudaniel7/go-restful-api-example/internal/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := propsConfig()
	configLogger(err)

	dsn := configDatabase()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	storeController, bookController, healthController := initServices(db)
	router := registerRoutes(storeController, bookController, healthController)

	port := viper.GetInt("server.port")
	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic("failed to start server")
	}
}

func configLogger(err error) {
	logLevel := viper.GetString("logLevel")
	var zapLevel zapcore.Level

	switch logLevel {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "trace":
		zapLevel = zapcore.DebugLevel + 1 // Zap doesn't have a direct trace level, so we use debug
	default:
		zapLevel = zap.InfoLevel
	}

	config := zap.Config{
		Level:         zap.NewAtomicLevelAt(zapLevel),
		Development:   false,
		Encoding:      "json",
		OutputPaths:   []string{"stdout"},
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
}

func propsConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	return err
}

func registerRoutes(storeController *controller.StoreController, bookController *controller.BookController, healthController *controller.HealthController) *gin.Engine {
	router := gin.Default()
	storeController.RegisterRoutes(router)
	bookController.RegisterRoutes(router)
	healthController.RegisterRoutes(router)
	return router
}

func initServices(db *gorm.DB) (*controller.StoreController, *controller.BookController, *controller.HealthController) {

	storeService := service.NewStoreService(db)
	storeController := controller.NewStoreController(storeService)

	bookService := service.NewBookService(db)
	bookController := controller.NewBookController(bookService)

	healthController := controller.NewHealthController()
	return storeController, bookController, healthController
}

func configDatabase() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.dbname"),
	)
	return dsn
}
