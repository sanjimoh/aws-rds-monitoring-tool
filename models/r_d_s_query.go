// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RDSQuery An RDS query
//
// swagger:model RDSQuery
type RDSQuery struct {

	// An RDS query
	// Required: true
	Query *string `json:"query"`
}

// Validate validates this r d s query
func (m *RDSQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RDSQuery) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RDSQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RDSQuery) UnmarshalBinary(b []byte) error {
	var res RDSQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
