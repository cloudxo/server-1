// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// User user
// swagger:model User
type User struct {

	// ID
	ID int64 `json:"ID,omitempty" gorm:"primary_key;auto_increment"`

	// created at
	CreatedAt int64 `json:"createdAt,omitempty"`

	// email
	Email string `json:"email,omitempty" gorm:"unique_index"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// has avatar
	HasAvatar bool `json:"hasAvatar,omitempty"`

	// is admin
	IsAdmin bool `json:"isAdmin,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// last session at
	LastSessionAt int64 `json:"lastSessionAt,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updatedAt,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}