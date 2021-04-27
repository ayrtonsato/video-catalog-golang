package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ayrtonsato/video-catalog-golang/pkg/logger"
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

func TransactionRollback(tx *sql.Tx, log logger.Logger, err error) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		log.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
	}
}

func TransactionCommit(tx *sql.Tx, log logger.Logger) error {
	if commitErr := tx.Commit(); commitErr != nil {
		log.Errorf("Fail to commit transaction: %v", commitErr)
		return commitErr
	}
	return nil
}
