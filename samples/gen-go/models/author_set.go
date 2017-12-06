package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// AuthorSet author set
// swagger:model AuthorSet
type AuthorSet struct {

	// random prop
	RandomProp int64 `json:"randomProp,omitempty"`

	// results
	Results AuthorArray `json:"results"`
}

// Validate validates this author set
func (m *AuthorSet) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
