// Code generated by go-swagger; DO NOT EDIT.

package aws_rds_monitoring_tool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"aws-rds-monitoring-tool/models"
)

// NewPostV1RdsQueriesParams creates a new PostV1RdsQueriesParams object
// no default values defined in spec.
func NewPostV1RdsQueriesParams() PostV1RdsQueriesParams {

	return PostV1RdsQueriesParams{}
}

// PostV1RdsQueriesParams contains all the bound params for the post v1 rds queries operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostV1RdsQueries
type PostV1RdsQueriesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*RDS query parameter.
	  In: body
	*/
	RdsQueriesExecAttr *models.RdsQueriesExecAttr
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostV1RdsQueriesParams() beforehand.
func (o *PostV1RdsQueriesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.RdsQueriesExecAttr
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("rdsQueriesExecAttr", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.RdsQueriesExecAttr = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}