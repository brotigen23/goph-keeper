package domain

import "time"

type BaseModel struct {
	ID int

	CreatedAt time.Time
	UpdatedAt time.Time
}
