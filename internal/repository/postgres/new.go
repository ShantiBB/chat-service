package postgres

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"chat-service/internal/config"
	"chat-service/internal/repository/models"
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
		slog.Error("failed to connect to database", "error", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("failed to get database instance", "error", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Postgres.Pool.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Postgres.Pool.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.Pool.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.Postgres.Pool.ConnMaxIdleTime)

	return &Repository{
		db: db,
	}
}

func (r *Repository) Migrate() {
	var err error
	if err = r.db.AutoMigrate(&models.Chat{}); err != nil {
		return
	}
	if err = r.db.AutoMigrate(&models.Message{}); err != nil {
		return
	}
}
