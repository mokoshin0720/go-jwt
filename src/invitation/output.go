package invitation

import (
	"github.com/ispec-inc/sample/pkg/domain/model"
)

type FindCodeOutput struct {
	Invitation model.Invitation
}

type AddCodeOutput struct {
	Invitation model.Invitation
}
