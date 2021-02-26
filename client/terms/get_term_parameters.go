// Code generated by go-swagger; DO NOT EDIT.

package terms

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

// NewGetTermParams creates a new GetTermParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTermParams() *GetTermParams {
	return &GetTermParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTermParamsWithTimeout creates a new GetTermParams object
// with the ability to set a timeout on a request.
func NewGetTermParamsWithTimeout(timeout time.Duration) *GetTermParams {
	return &GetTermParams{
		timeout: timeout,
	}
}

// NewGetTermParamsWithContext creates a new GetTermParams object
// with the ability to set a context for a request.
func NewGetTermParamsWithContext(ctx context.Context) *GetTermParams {
	return &GetTermParams{
		Context: ctx,
	}
}

// NewGetTermParamsWithHTTPClient creates a new GetTermParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTermParamsWithHTTPClient(client *http.Client) *GetTermParams {
	return &GetTermParams{
		HTTPClient: client,
	}
}

/* GetTermParams contains all the parameters to send to the API endpoint
   for the get term operation.

   Typically these are written to a http.Request.
*/
type GetTermParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get term params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTermParams) WithDefaults() *GetTermParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get term params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTermParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get term params
func (o *GetTermParams) WithTimeout(timeout time.Duration) *GetTermParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get term params
func (o *GetTermParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get term params
func (o *GetTermParams) WithContext(ctx context.Context) *GetTermParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get term params
func (o *GetTermParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get term params
func (o *GetTermParams) WithHTTPClient(client *http.Client) *GetTermParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get term params
func (o *GetTermParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get term params
func (o *GetTermParams) WithID(id string) *GetTermParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get term params
func (o *GetTermParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTermParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
