// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TelnetResult telnet result
//
// swagger:model TelnetResult
type TelnetResult struct {

	// error
	Error string `json:"error,omitempty"`

	// ping
	Ping int64 `json:"ping,omitempty"`

	// response time ms
	ResponseTimeMs int64 `json:"response-time-ms,omitempty"`
}

// Validate validates this telnet result
func (m *TelnetResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this telnet result based on context it is used
func (m *TelnetResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TelnetResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TelnetResult) UnmarshalBinary(b []byte) error {
	var res TelnetResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
