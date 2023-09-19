package migrations

import (
	"api-gin-gorm/models"
	"gorm.io/gorm"
)

type PostPostgreSQLDB struct {
	DB *gorm.DB
}

func (p PostPostgreSQLDB) MigrateTable() {
	err := p.DB.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}
}
