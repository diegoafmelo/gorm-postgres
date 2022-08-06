package database

import (
	"fmt"
	"github.com/mercadolibre/dr-gorm/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

type Database struct {
	Username string
	Password string
	Name     string
}

func NewDataBaseConnection(database Database) *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", database.Username, database.Password, database.Name)

	strategy := GetConfigNameStrategy()
	config := GetConfig(strategy)

	db, err := gorm.Open(postgres.Open(dbURL), config)

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Subject{})
	db.AutoMigrate(&model.Participant{})

	return db
}

func GetConfig(strategy schema.NamingStrategy) *gorm.Config {
	return &gorm.Config{
		NamingStrategy: strategy,
		Logger:         logger.Default.LogMode(logger.Error),
		PrepareStmt:    true,
	}
}

func GetConfigNameStrategy() schema.NamingStrategy {
	return schema.NamingStrategy{
		TablePrefix:   "table_",
		SingularTable: true,
		NoLowerCase:   false,
	}
}
