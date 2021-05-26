package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ispec-inc/sample/pkg/config"
)

func Init() (*gorm.DB, func() error, error) {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.RDS.User, config.RDS.Password,
		config.RDS.Host, config.RDS.Port,
		config.RDS.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(config.RDS.MaxIdle)
	sqlDB.SetMaxOpenConns(config.RDS.MaxOpen)
	sqlDB.SetConnMaxLifetime(config.RDS.MaxLifetime)

	return db, cleanup(db), nil
}

func cleanup(db *gorm.DB) func() error {
	sqlDB, _ := db.DB()
	return sqlDB.Close
}
