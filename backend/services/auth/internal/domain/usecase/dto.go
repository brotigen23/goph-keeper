package usecase

import "github.com/brotigen23/goph-keeper/auth/internal/domain/entity"

type User struct {
	Login    string
	Password string
}

type CreatedUser struct {
	ID int
	User
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		Login:    u.Login,
		Password: u.Password,
	}
}

func (u *User) FromEntity(e entity.User) {
	u.Login = e.Login
	u.Password = e.Password
}

func (u *CreatedUser) ToEntity() *entity.User {
	return &entity.User{
		ID:       u.ID,
		Login:    u.Login,
		Password: u.Password,
	}
}

func (u *CreatedUser) FromEntity(e entity.User) {
	u.ID = e.ID
	u.Login = e.Login
	u.Password = e.Password
}
