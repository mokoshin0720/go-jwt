package dao

import (
	"errors"

	"github.com/ispec-inc/sample/pkg/apperror"
	"github.com/ispec-inc/sample/pkg/domain/model"
	"github.com/ispec-inc/sample/pkg/infra/entity"
	"github.com/ispec-inc/sample/pkg/transaction"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

// Userをdb情報で初期化？
func NewUser(db *gorm.DB) User {
	return User{db}
}

func (repo User) Find(id int64) (model.User, apperror.Error) {
	var u entity.User
	
	// 引数のidを持つUserが存在しなければエラーを返す
	if err := repo.db.First(&u, id).Error; err != nil {
		return model.User{}, newGormError(
			err, "error searching user in database",
		)
	}
	return u.ToModel(), nil
}

func (repo User) Create(mu model.User) apperror.Error {
	f := func(tx *gorm.DB) apperror.Error {
		var us []entity.User
		err := tx.
			Set("gorm:query_option", "for update").
			Find(&us, "id = ?", mu.ID).
			Error
		
		if err != nil {
			return newGormError(err, "error searching user in database")
		}

		if len(us) > 0 {
			return apperror.New(apperror.CodeInvalid, errors.New("error: user name is already exists"))
		}

		u := entity.NewUserFromModel(mu)
		if err := tx.Create(&u).Error; err != nil {
			return newGormError(err, "error inserting user in database")
		}

		return nil
	}

	if aerr := transaction.Run(repo.db, f); aerr != nil {
		return aerr
	}

	return nil
}

func (repo User) Login(mu model.User) apperror.Error {

}