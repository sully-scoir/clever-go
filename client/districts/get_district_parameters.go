// Code generated by go-swagger; DO NOT EDIT.

package districts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetDistrictParams creates a new GetDistrictParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDistrictParams() *GetDistrictParams {
	return &GetDistrictParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistrictParamsWithTimeout creates a new GetDistrictParams object
// with the ability to set a timeout on a request.
func NewGetDistrictParamsWithTimeout(timeout time.Duration) *GetDistrictParams {
	return &GetDistrictParams{
		timeout: timeout,
	}
}

// NewGetDistrictParamsWithContext creates a new GetDistrictParams object
// with the ability to set a context for a request.
func NewGetDistrictParamsWithContext(ctx context.Context) *GetDistrictParams {
	return &GetDistrictParams{
		Context: ctx,
	}
}

// NewGetDistrictParamsWithHTTPClient creates a new GetDistrictParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDistrictParamsWithHTTPClient(client *http.Client) *GetDistrictParams {
	return &GetDistrictParams{
		HTTPClient: client,
	}
}

/* GetDistrictParams contains all the parameters to send to the API endpoint
   for the get district operation.

   Typically these are written to a http.Request.
*/
type GetDistrictParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get district params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictParams) WithDefaults() *GetDistrictParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get district params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get district params
func (o *GetDistrictParams) WithTimeout(timeout time.Duration) *GetDistrictParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get district params
func (o *GetDistrictParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get district params
func (o *GetDistrictParams) WithContext(ctx context.Context) *GetDistrictParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get district params
func (o *GetDistrictParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get district params
func (o *GetDistrictParams) WithHTTPClient(client *http.Client) *GetDistrictParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get district params
func (o *GetDistrictParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get district params
func (o *GetDistrictParams) WithID(id string) *GetDistrictParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get district params
func (o *GetDistrictParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistrictParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
