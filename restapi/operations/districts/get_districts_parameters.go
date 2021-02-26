// Code generated by go-swagger; DO NOT EDIT.

package districts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetDistrictsParams creates a new GetDistrictsParams object
//
// There are no default values defined in the spec.
func NewGetDistrictsParams() GetDistrictsParams {

	return GetDistrictsParams{}
}

// GetDistrictsParams contains all the bound params for the get districts operation
// typically these are obtained from a http.Request
//
// swagger:parameters getDistricts
type GetDistrictsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Count *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDistrictsParams() beforehand.
func (o *GetDistrictsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCount, qhkCount, _ := qs.GetOK("count")
	if err := o.bindCount(qCount, qhkCount, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCount binds and validates parameter Count from query.
func (o *GetDistrictsParams) bindCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Count = &raw

	if err := o.validateCount(formats); err != nil {
		return err
	}

	return nil
}

// validateCount carries on validations for parameter Count
func (o *GetDistrictsParams) validateCount(formats strfmt.Registry) error {

	if err := validate.EnumCase("count", "query", *o.Count, []interface{}{"", "true"}, true); err != nil {
		return err
	}

	return nil
}