package model

import "time"

type Model interface {
	GetID() int
}

type Base struct {
	ID int `db:"id"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type BaseData struct {
	UserID   int    `db:"user_id"`
	Metadata string `db:"metadata"`
}

// The user's entity that stores the ID, login, and password
type User struct {
	Base
	Login    string `db:"login"`
	Password string `db:"password"`
}

// Stores data about user's accounts
type Account struct {
	Base
	BaseData
	Login    string `db:"login"`
	Password string `db:"password"`
}

// Stores the user's text data
type TextData struct {
	Base
	BaseData
	Data string `db:"data"`
}

// Stores the user's binary data
type BinaryData struct {
	Base
	BaseData
	Data []byte `db:"data"`
}

// Stores the user's bank card information
type CardData struct {
	Base
	BaseData
	Number         string    `db:"number"`
	CardholderName string    `db:"cardholder_name"`
	Expire         time.Time `db:"expire"`
	CVV            string    `db:"cvv"`
}

func (u User) GetID() int       { return u.ID }
func (d Account) GetID() int    { return d.ID }
func (d TextData) GetID() int   { return d.ID }
func (d BinaryData) GetID() int { return d.ID }
func (d CardData) GetID() int   { return d.ID }
