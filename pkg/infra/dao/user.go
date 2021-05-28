package dao

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
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
	if err := repo.db.First(&u, "id = ?",id).Error; err != nil {
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
			Set("gorm:query_option", "for update").
			Find(&us, "email = ?", mu.Email).
			Error
		if err != nil {
			return newGormError(err, "error searching user in database")
		}

		if len(us) > 0 {
			return apperror.New(apperror.CodeInvalid, errors.New("error: user email is already exists"))
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


func (repo User) EmailPass(email string) (model.User, apperror.Error) {
	var u entity.User

	// 引数のEmailを持つUserが存在しなければエラーを返す
	if err := repo.db.Where("email = ?", email).First(&u).Error; err != nil {
		return model.User{}, newGormError(
			err, "error searching user with email in database",
		)
	}
	return u.ToModel(), nil
}

func (repo User) Publish(email string, password string) (model.User, apperror.Error) {
	var u entity.User
	
	// Emailを持つUserが存在しているかどうかのエラーハンドリング
	if err := repo.db.Where("email = ?", email).First(&u).Error; err != nil {
		return model.User{}, newGormError(
			err, "指定したEmailをもつユーザーが存在しません。",
		)
	}

	// EmailとPasswordが一致するかどうかのハンドリング
	if err := repo.db.Where("email = ? AND password = ?", email, password).First(&u).Error; err != nil {
		return model.User{}, newGormError(
			err, "メールアドレスとパスワードが一致しません。",
		)
	}

	return u.ToModel(), nil
}

func (repo User) ParseToken(tokenString string) (model.User, apperror.Error) {
	var u entity.User

	// tokenの解析
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINKEY")), nil
	})

	// tokenに関するエラーハンドリング
	if !token.Valid {
		return model.User{}, newGormError(
			err, "tokenが正しくありません。",
		)
	}

	// tokenのUserが存在するかどうか
	id := claims["id"]
	if err := repo.db.First(&u, id).Error; err != nil {
		return model.User{}, newGormError(
			err, "tokenのユーザーが存在しません。",
		)
	}

	return u.ToModel(), nil
}