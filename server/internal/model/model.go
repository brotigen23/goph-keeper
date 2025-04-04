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
}

// An entity that stores data about user accounts
type AccountData struct {
	ID     int
	UserID int

	Login    string
	Password string
}

// An entity that stores the user's text data
type TextData struct {
	ID     int
	UserID int

	Data string
}

// The entity that stores the user's binary data
type BinaryData struct {
	ID     int
	UserID int

	Data []byte
}

// The entity that stores the user's bank card information
type CardData struct {
	ID     int
	UserID int

	Number         string
	CardholderName string
	Expire         time.Time
	cvv            string
}

// An entity that stores metadata about record in a certain table
type Metadata struct {
	ID        int
	TableName string
	RowID     int
	data      string
}
