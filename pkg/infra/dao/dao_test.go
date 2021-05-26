package dao

import (
	"os"
	"testing"

	"gorm.io/gorm"
	"github.com/tanimutomo/sqlfile"

	"github.com/ispec-inc/sample/pkg/config"
	"github.com/ispec-inc/sample/pkg/mysql"
)

var (
	db *gorm.DB
)

func TestMain(m *testing.M) {
	config.Init()

	d, cleanup, err := mysql.Init()
	if err != nil {
		panic(err)
	}

	db = d
	defer cleanup()

	os.Exit(m.Run())
}

func prepareTestData(filepath string) error {
	s := sqlfile.New()
	if err := s.Files("./testdata/delete.sql", filepath); err != nil {
		return err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return err
	}
	if _, err := s.Exec(sqlDB); err != nil {
		return err
	}
	return nil
}
