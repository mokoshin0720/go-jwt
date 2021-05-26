//go:generate mockgen -package mock -source=invitation.go -destination=../mock/invitation.go

package repository

import (
	"github.com/ispec-inc/sample/pkg/apperror"
	"github.com/ispec-inc/sample/pkg/domain/model"
)

type Invitation interface {
	Find(id int64) (model.Invitation, apperror.Error)
	FindByUserID(uid int64) (model.Invitation, apperror.Error)
	Create(minv model.Invitation) apperror.Error
}
