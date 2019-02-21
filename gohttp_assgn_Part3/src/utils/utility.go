package utils

import (
	"github.com/satori/go.uuid"
)

//NewUUID creates new uuid
func NewUUID() uuid.UUID {
	id, _ := uuid.NewV4()
	return id
}
