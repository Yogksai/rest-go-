package storage

import (
	"log"
	"rest-api/internal/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

var (
	once sync.Once
	db   *gorm.DB
)

func Connect() *Storage {
	once.Do(func() {
		cfg := config.MustLoad()

		dsn := "host=" + cfg.DB.Host +
			" port=" + cfg.DB.Port +
			" user=" + cfg.DB.User +
			" dbname=" + cfg.DB.Name +
			" password=" + cfg.DB.Pass +
			" sslmode=disable"

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database")
		}

	})
	return &Storage{DB: db}
}
