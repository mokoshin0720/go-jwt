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
	if err := repo.db.First(&u, "id = ?", id).Error; err != nil {
		return model.User{}, newGormError(
			err, "error searching user in database（dao/findだよ）",
		)
	}
	return u.ToModel(), nil
}

func (repo User) FindByEmail(email string) (model.User, apperror.Error) {
	var u entity.User

	if err := repo.db.First(&u, "email = ?", email).Error; err != nil {
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
			Set("gorm:query_option", "for update"). // select時に行をロックする
			Find(&us, "email = ?", mu.Email).
			Error
		if err != nil {
			return newGormError(err, "error searching user in database")
		}

		// 指定したEmailをもつUserがすでにいる場合
		if len(us) > 0 {
			return apperror.New(apperror.CodeInvalid, errors.New("error: user email is already exists"))
		}

		// 指定したEmailをもつUserを作成
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