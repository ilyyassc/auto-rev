package service

import (
	"gorm.io/gorm"
)

var g *gorm.DB

func SetService(gDB *gorm.DB) {
	g = gDB
}