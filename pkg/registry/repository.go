package registry

import (
	"gorm.io/gorm"

	"github.com/ispec-inc/sample/pkg/infra/dao"
	"github.com/ispec-inc/sample/pkg/mysql"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() (Repository, func() error) {
	db, cleanup, err := mysql.Init()
	if err != nil {
		panic(err)
	}

	repo := Repository{
		db: db,
	}
	f := func() error {
		return cleanup()
	}
	return repo, f
}

func (repo Repository) NewInvitation() dao.Invitation {
	return dao.NewInvitation(repo.db)
}

// daoを初期化する
func (repo Repository) NewUser() dao.User {
	return dao.NewUser(repo.db)
}