package ingest

import (
	"fmt"

	"github.com/peterhoward42/fabble/ingest/storer"
	"github.com/peterhoward42/fabble/ingest/usrparser"
)

// Ingestor offers to build *User*s iteratively using its UserParser, and
// for each thus found, storing it via its Storer object.
type Ingestor struct {
	userParser usrparser.UserParser
	storer     storer.Storer
}

// NewIngestor constructs an Ingestor bound to the given usr parser and
// storer.
func NewIngestor(userParser usrparser.UserParser, storer storer.Storer) *Ingestor {
	return &Ingestor{
		userParser: userParser,
		storer:     storer,
	}
}

// ParseInputAndStore is the mandate to perform the parsing and storage.
func (ingestor *Ingestor) ParseInputAndStore() (int, error) {

	count := 0
	for {
		user, err := ingestor.userParser.Next()
		if err != nil {
			return count, fmt.Errorf("ingestoer.UserParser.Next(): %v", err)
		}
		if user == nil {
			break
		}
		count++
		err = ingestor.storer.Store(user)
	}
	return count, nil
}
