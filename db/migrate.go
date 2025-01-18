package database

import "github.com/jkroy667/fluxy/models"

// Migrate handles migrations
func Migrate(db *DB) {
	var migrationModels = []interface{}{&models.Client{}}
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
