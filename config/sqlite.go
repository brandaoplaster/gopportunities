package config

import (
	"github.com/brandaoplaster/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	db, erro := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if erro != nil {
		logger.ErrorFormat("Error initializing SQLite: %v", erro)
		return nil, erro
	}

	erro = db.AutoMigrate(&schemas.Opening{})

	if erro != nil {
		logger.ErrorFormat("Error migrating SQLite: %v", erro)
		return nil, erro
	}

	return db, nil
}
