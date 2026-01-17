package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"chat-service/internal/config"
)

type Repository struct {
	db *gorm.DB
}

func New(cfg *config.Config) *Repository {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DB,
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode,
		cfg.Postgres.TimeZone,
	)

	db, err := gorm.Open(
		postgres.Open(dsn), &gorm.Config{
			Logger: gormlogger.Default.LogMode(gormlogger.Silent),
		},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get database instance: %v", err))
	}

	sqlDB.SetMaxOpenConns(cfg.Postgres.Pool.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Postgres.Pool.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.Pool.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.Postgres.Pool.ConnMaxIdleTime)

	return &Repository{
		db: db,
	}
}
