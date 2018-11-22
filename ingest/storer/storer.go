package storer

import "github.com/peterhoward42/fabble/store/usr"

// Storer offers to Store a *User*.
type Storer interface {
	Store(user *usr.User) error
}
