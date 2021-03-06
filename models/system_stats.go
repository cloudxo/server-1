// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SystemStats system stats
// swagger:model SystemStats
type SystemStats struct {

	// alloc mem
	AllocMem int64 `json:"allocMem,omitempty"`

	// go version
	GoVersion string `json:"goVersion,omitempty"`

	// num g c
	NumGC int64 `json:"numGC,omitempty"`

	// num goroutines
	NumGoroutines int64 `json:"numGoroutines,omitempty"`

	// num sessions
	NumSessions int64 `json:"numSessions,omitempty"`

	// system mem
	SystemMem int64 `json:"systemMem,omitempty"`

	// total alloc mem
	TotalAllocMem int64 `json:"totalAllocMem,omitempty"`

	// uptime
	Uptime int64 `json:"uptime,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this system stats
func (m *SystemStats) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemStats) UnmarshalBinary(b []byte) error {
	var res SystemStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
