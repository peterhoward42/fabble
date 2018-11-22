package usrparser

import (
	"github.com/peterhoward42/fabble/store/usr"
)

// UserParser is willing to make *User*s iteratively.
type UserParser interface {
	Next() (*usr.User, error)
}
