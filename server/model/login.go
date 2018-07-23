package models

import (
	"time"
)

type User struct {
	ID        int64     `json:"Id" xorm:"pk"`
	Username  string    `json:"username" xorm:"unique"`
	Password  string    `json:"password"`
	CreateAt  time.Time `json:"-" xorm:"created"`
	LastModif time.Time `json:"-" xorm:"updated"`
}

type Profile struct {
	NurseID    string    `json:"NurseID" xorm:"unique 'nurse_id'"`
	OfficeID   string    `json:"officeID" xorm:"'office_id'"`
	Lastname   string    `json:"lastname" xorm:"'lastname'"`
	Firstname  string    `json:"firstname" xorm:"'firstname'"`
	Birth      time.Time `json:"birthdate" xorm:"'birth_date'"`
	Email      string    `json:"email" xorm:"'email'"`
	PostalCode int       `json:"postalCode" xorm:" 'postal_code'"`
	Complement string    `json:"complement" xorm:"'complement'"`
	City       string    `json:"city" xorm:"'city'"`
	Phone      string    `json:"phone" xorm:"'phone_number'"`
}
