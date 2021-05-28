package entity

import (
	"time"

	"github.com/ispec-inc/sample/pkg/domain/model"
)

// 実際のDB上のデータを表現
type User struct {
	ID        int64      `gorm:"column:id; type:bigint(20) auto_increment; not null; primary_key"`
	Name      string     `gorm:"column:name"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	CreatedAt *time.Time `gorm:"column:created_at; not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at; not null"`
}

func NewUserFromModel(m model.User) User {
	return User{
		ID:       m.ID,
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
	}
}

func (u User) ToModel() model.User {
	return model.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}