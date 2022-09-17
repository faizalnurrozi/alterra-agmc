package config

import "gorm.io/gorm"

type Contract struct {
	DB *gorm.DB
	HTTPResponse
}
