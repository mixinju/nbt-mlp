package dao

import (
	"nbt-mlp/dao/model"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var logger, _ = zap.NewProduction()

func InitDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/nbt_mlp?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logger.Sugar().Fatalf("failed to connect database: %v", err)
	}

	// Perform automatic migration
	err = DB.AutoMigrate(&model.User{}, &model.Host{}, &model.Container{}, &model.AccessGroup{})
	if err != nil {
		logger.Sugar().Fatalf("failed to migrate database: %v", err)
	}

	sqlDb, err := DB.DB()
	if err != nil {
		logger.Sugar().Fatalf("failed to get db: %v", err)
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
}
