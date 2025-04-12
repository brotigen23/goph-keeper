package model

import "time"

// Stores metadata about record in a certain table
type Metadata struct {
	ID   int
	Data string

	CreatedAt time.Time
	UpdatedAt time.Time
}
