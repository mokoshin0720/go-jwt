package repository

import (
	"github.com/ispec-inc/sample/pkg/apperror"
	"github.com/ispec-inc/sample/pkg/domain/model"
)

type User interface {
	Find(id int64) (model.User, apperror.Error)
	Create(mu model.User) apperror.Error
	FindByEmail(email string) (model.User, apperror.Error)
}