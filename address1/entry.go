package address1

import (
	"time"
)

// Entry - customer address1 struct
type Entry struct {
	CustomerId string    `json:"customer-id"`
	CreatedTS  time.Time `json:"created-ts"`

	AddressLine1 string `json:"region"`
	AddressLine2 string `json:"zone"`
	City         string `json:"sub-zone"`
	State        string `json:"host"`
	PostalCode   string `json:"instance-id"`
	Email        string `json:"email"`
}
