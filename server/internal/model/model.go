package model

import "time"

/*

Типы хранимой информации
- пары логин/пароль;
- произвольные текстовые данные;
- произвольные бинарные данные;
- данные банковских карт.

Для любых данных должна быть возможность хранения произвольной текстовой метаинформации
(принадлежность данных к веб-сайту, личности или банку, списки одноразовых кодов активации и прочее).

*/

// The user's entity that stores the ID, login, and password
type User struct {
	ID int

	Login    string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Stores data about user accounts
type AccountData struct {
	ID     int
	UserID int

	Login    string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Stores the user's text data
type TextData struct {
	ID     int
	UserID int

	Data string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Stores the user's binary data
type BinaryData struct {
	ID     int
	UserID int

	Data []byte

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Stores the user's bank card information
type CardData struct {
	ID     int
	UserID int

	Number         string
	CardholderName string
	Expire         time.Time
	CVV            string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Stores metadata about record in a certain table
type Metadata struct {
	ID        int
	TableName string
	RowID     int
	data      string

	CreatedAt time.Time
	UpdatedAt time.Time
}
