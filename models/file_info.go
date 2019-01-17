// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FileInfo file info
// swagger:model FileInfo
type FileInfo struct {

	// ID
	ID int64 `json:"ID,omitempty" gorm:"primary_key;auto_increment"`

	// is dir
	IsDir bool `json:"isDir,omitempty"`

	// last changed
	LastChanged int64 `json:"lastChanged,omitempty"`

	// mime type
	MimeType string `json:"mimeType,omitempty"`

	// name
	Name string `json:"name,omitempty" gorm:"index:fullPath"`

	// owner ID
	OwnerID int64 `json:"ownerID,omitempty"`

	// parent ID
	ParentID int64 `json:"parentID,omitempty"`

	// path
	Path string `json:"path,omitempty" gorm:"index:fullPath"`

	// share ID
	ShareID int64 `json:"shareID,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// starred
	Starred bool `json:"starred,omitempty"`
}

// Validate validates this file info
func (m *FileInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfo) UnmarshalBinary(b []byte) error {
	var res FileInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}