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

// RDS An RDS instance
//
// swagger:model RDS
type RDS struct {

	// Availability zone.
	// Required: true
	AvailabilityZone *string `json:"availabilityZone"`

	// RDS DB name.
	// Required: true
	Name *string `json:"name"`

	// Status of RDS DB instance
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this r d s
func (m *RDS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RDS) validateAvailabilityZone(formats strfmt.Registry) error {

	if err := validate.Required("availabilityZone", "body", m.AvailabilityZone); err != nil {
		return err
	}

	return nil
}

func (m *RDS) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RDS) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RDS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RDS) UnmarshalBinary(b []byte) error {
	var res RDS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}