package test

import (
	internal "github.com/pancudaniel7/go-restful-api-example/internal/entity"
	services "github.com/pancudaniel7/go-restful-api-example/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func setupDatabase() *gorm.DB {
	dsn := "appuser:password@tcp(localhost:3306)/restful?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// AutoMigrate for test schema
	err = db.AutoMigrate(&internal.Store{})
	if err != nil {
		return nil
	}

	return db
}

func TestAddStore(t *testing.T) {
	db := setupDatabase()
	storeService := services.NewStoreService(db)

	// Test AddStore
	store, err := storeService.AddStore("Test Store", "123 Test St")
	if err != nil {
		t.Errorf("Error creating store: %v", err)
	}

	if store.Name != "Test Store" || store.Location != "123 Test St" {
		t.Errorf("Store details do not match")
	}
}
