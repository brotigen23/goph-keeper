package account

import "errors"

var (
	ErrNotImplemented = errors.New("Not implemented")
)

func translateDBErr(err error) error {
	switch err {
	default:
		return err
	}
}
