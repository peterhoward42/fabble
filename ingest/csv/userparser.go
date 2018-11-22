package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/peterhoward42/fabble/ingest/telephone"
	"github.com/peterhoward42/fabble/store/usr"
)

// UserParser is able to parse and construct *User*s by parsing the CSV
// formatted data provided by an io.Reader.
type UserParser struct {
	csvReader *csv.Reader
}

// NewUserParser constructs an UserParser bound to the given io.Reader.
func NewUserParser(inputSource io.Reader) *UserParser {
	r := csv.NewReader(inputSource)
	return &UserParser{r}
}

// Next provides one *User* by consuming one csv record from its reader.
// It signals EOF by returning a nil User.
func (parser UserParser) Next() (*usr.User, error) {

	record, err := parser.csvReader.Read()
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("csv.UserParser.Next(): %v", err)
	}

	id, _ := strconv.Atoi(record[0]) // todo check err
	telephone := telephone.Sanitize(record[3])
	user := usr.User{
		ID:        id,
		Name:      record[1],
		Email:     record[2],
		Telephone: telephone,
	}
	return &user, nil
}
