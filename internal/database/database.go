package database

import (
	"fmt"

	"github.com/Dbone29/dflow/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type DflowDatabase struct {
	logger   *zap.Logger
	Database *gorm.DB
}

func InitDatabase(logger *zap.Logger, config *config.DatabaseConfig) *DflowDatabase {
	zlog := zapgorm2.New(logger)
	zlog.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	var db *gorm.DB
	var err error

	if len(config.Host+config.User+config.Password+config.DBName) == 0 {
		logger.Info("using sqllite database")

		db, err = gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{Logger: zlog})
	} else {
		logger.Info("using postgres database")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Europe/Berlin", config.Host, config.User, config.Password, config.DBName)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: zlog})
	}

	if err != nil {
		logger.Panic("failed to connect database", zap.Error(err))
	}

	// Migrate the schema
	db.AutoMigrate()

	return &DflowDatabase{
		logger:   logger,
		Database: db,
	}
}
