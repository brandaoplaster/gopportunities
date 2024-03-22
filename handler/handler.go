package handler

import (
	"github.com/brandaoplaster/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Initialize() {
	logger = config.GetLogger("Handler")
	db = config.GetDB()
}
