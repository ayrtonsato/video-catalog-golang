package repositories

import (
	"errors"
	"fmt"
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
	index := 0
	for _, value := range fields {
		if index == 0 {
			stmt = stmt + fmt.Sprintf("%s=$%v", value, index+1)
		} else {
			stmt = stmt + fmt.Sprintf(", %s=$%v", value, index+1)
		}
		index++
	}
	stmt = fmt.Sprintf("%v, updated_at=(NOW()) WHERE id=$%v", stmt, index+1)
	return stmt, nil
}
