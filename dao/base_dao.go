package dao

import (
	"gorm.io/gorm"
)

var g *gorm.DB

func SetDao(gDB *gorm.DB) {
	g = gDB
}
