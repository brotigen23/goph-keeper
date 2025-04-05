package pgErrors

import (
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

func CheckIfUniqueViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == pgerrcode.UniqueViolation {
			return true
		} else {
			return false
		}
	}
	return false
}
