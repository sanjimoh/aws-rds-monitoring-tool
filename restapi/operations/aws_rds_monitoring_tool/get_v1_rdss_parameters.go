// Code generated by go-swagger; DO NOT EDIT.

package aws_rds_monitoring_tool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetV1RdssParams creates a new GetV1RdssParams object
// no default values defined in spec.
func NewGetV1RdssParams() GetV1RdssParams {

	return GetV1RdssParams{}
}

// GetV1RdssParams contains all the bound params for the get v1 rdss operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetV1Rdss
type GetV1RdssParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetV1RdssParams() beforehand.
func (o *GetV1RdssParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
