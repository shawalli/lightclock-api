// models.go

package main

import (
	"github.com/google/uuid"
)

type event struct {
	ID		uuid.UUID	`json:"id"`
	Label	string		`json:"label"`
}

func (evt *event) createEvent() error {
	evt.ID, _ = uuid.NewUUID()

	return nil
}
