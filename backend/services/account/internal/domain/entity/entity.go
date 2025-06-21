package entity

import "time"

type Account struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time

	Login    string
	Password string
}

func (a *Account) Validate() error {
	if a.Login == "" || a.Password == "" {
		return ErrInvalidCredentials
	}
	return nil
}
