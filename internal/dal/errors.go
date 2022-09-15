package dal

import "errors"

var (
	// ErrNoRedirect when no redirect entity is found
	ErrNoRedirect = errors.New("redirect is not present")
	// ErrDeleteRedirect when the redirect has been deleted
	ErrDeleteRedirect = errors.New("redirect has been deleted")
)
