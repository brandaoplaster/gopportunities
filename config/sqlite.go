package config

import (
	"os"

	"github.com/brandaoplaster/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	fileDB := "./db/main.db"
	_, erro := os.Stat(fileDB)

	if os.IsNotExist(erro) {
		logger.Info("Database file does not exist, creating it...")
		erro = os.Mkdir("./db", os.ModePerm)
		if erro != nil {
			return nil, erro
		}
		file, erro := os.Create(fileDB)
		if erro != nil {
			return nil, erro
		}
		file.Close()
	}

	db, erro := gorm.Open(sqlite.Open(fileDB), &gorm.Config{})

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
