// models.go

package main

type event struct {
	ID 		int		`json:"id"`
	Label 	string 	`json:"label"`
}

func (evt *event) createEvent() error {
	evt.ID = 1

	return nil
}
