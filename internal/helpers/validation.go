package helpers

import (
	"errors"
	"github.com/gofrs/uuid"
)

func UUIDIsRequired(value interface{}) error {
	err := errors.New("cannot be blank")
	newUUID := value.(uuid.UUID)
	if newUUID == uuid.Nil || newUUID.String() == "" {
		return err
	}
	return nil
}
