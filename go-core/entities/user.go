package entity

type User struct {
	FirstName string `json:"first-name"`
	LastName string `json:"last-name"`
	Email string `json:"email"`
	AddressLine1 string `json:"address-line-1"`
	AddressLine2 string `json:"address-line-2"`
	Postcode string `json:"postcode"`
	PhoneNumber string `json:"phone"`