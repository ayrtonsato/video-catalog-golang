package repositories

import (
	"errors"
	"fmt"
	"time"
)

// RepoReader is an interface to override sql.Rows.Scan and sql.Row.Scan
type RepoReader interface {
	Scan(dest ...interface{}) error
}

func DynamicUpdateQuery(table string, fields []string) (string, error) {
	if len(fields) <= 0 {
		return "", errors.New("validation: fields must be greater than 0")
	}
	stmt := fmt.Sprintf("UPDATE %s SET ", table)
	for index, value := range fields {
		if index == 0 {
			stmt = stmt + fmt.Sprintf("%s=?", value)
		} else {
			stmt = stmt + fmt.Sprintf(", %s=?", value)
		}
	}
	s := time.Now().UTC().String()
	stmt = fmt.Sprintf("%v, updated_at='%v' WHERE id=?", stmt, s)
	return stmt, nil
}
