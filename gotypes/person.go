package gotypes

import "time"

type PhoneType string

const (
	PhoneTypeMobile = 0
	PhoneTypeHome = 0
	PhoneTypeWork = 0
)

// PhoneNumber of a person.
type PhoneNumber struct {
	Number string `json:"number"`
	Type PhoneType `json:"type"`
}

type Person struct {
	// Name of the person.
	Name string `json:"name"`
	Birthday time.Time `json:"birthday"`
	Email string `json:"email"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers"`
}
