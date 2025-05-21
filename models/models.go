package models

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DATABASE_URI = func() string {
	if dsn := os.Getenv("DATABASE_URI"); dsn != "" {
		return dsn
	}
	return "host=localhost port=5433 user=yugabyte password= dbname=yugabyte sslmode=disable TimeZone=UTC"
}()

var DB_MAX_CONNECTIONS = func() int {
	value := os.Getenv("DB_MAX_CONNECTIONS")
	if value != "" {
		if v, err := strconv.Atoi(value); err == nil && v > 0 {
			return v
		}
	}
	return 25
}()

var DB_MAX_IDLE_CONNECTIONS = func() int {
	value := os.Getenv("DB_MAX_IDLE_CONNECTIONS")
	if value != "" {
		if v, err := strconv.Atoi(value); err == nil && v > 0 {
			return v
		}
	}
	return 10
}()

var DB_MAX_CONNECTION_LIFETIME = func() time.Duration {
	value := os.Getenv("DB_MAX_CONNECTION_LIFETIME")
	if value != "" {
		if v, err := strconv.Atoi(value); err == nil && v > 0 {
			return time.Duration(v) * time.Minute
		}
	}
	return 5 * time.Minute
}()

var DB_AUTO_INITIALIZE_SCHEMA = func() bool {
	value := os.Getenv("DB_AUTO_INITIALIZE_SCHEMA")
	if value != "" {
		if v, err := strconv.ParseBool(value); err == nil {
			return v
		}
	}
	return false
}()

var db *gorm.DB
var mu sync.Mutex

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	log.Printf("initializing gorm.DB")
	mu.Lock()
	defer mu.Unlock()

	for db == nil {
		var err error

		db, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{})
		if err != nil {
			log.Printf("failed to connect to database: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}
	}

	// Configure connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance: " + err.Error())
	}

	sqlDB.SetMaxOpenConns(DB_MAX_CONNECTIONS)            // Set the maximum number of open connections
	sqlDB.SetMaxIdleConns(DB_MAX_IDLE_CONNECTIONS)       // Set the maximum number of idle connections
	sqlDB.SetConnMaxLifetime(DB_MAX_CONNECTION_LIFETIME) // Set the maximum lifetime of a connection

	if DB_AUTO_INITIALIZE_SCHEMA {
		log.Printf("automatically initializing database schema")
		if err := InitializeModels(db); err != nil {
			log.Printf("failed to initialize models: %v", err)
		}
		log.Printf("database schema initialized successfully")
	}

	return db
}

func InitializeModels(db *gorm.DB) error {
	// Automatically migrate models in this array
	for _, schema := range []interface{}{
		&User{},
	} {
		// Drop the table if it exists before creating it
		db.Migrator().DropTable(schema)

		if err := db.AutoMigrate(schema); err != nil {
			return err
		}
	}

	return nil
}
