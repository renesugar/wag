package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ExtendedError extended error
// swagger:model ExtendedError
type ExtendedError struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this extended error
func (m *ExtendedError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}