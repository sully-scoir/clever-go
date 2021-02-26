// Code generated by go-swagger; DO NOT EDIT.

package sections

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
	"github.com/go-openapi/swag"
)

// NewGetResourcesForSectionParams creates a new GetResourcesForSectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourcesForSectionParams() *GetResourcesForSectionParams {
	return &GetResourcesForSectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourcesForSectionParamsWithTimeout creates a new GetResourcesForSectionParams object
// with the ability to set a timeout on a request.
func NewGetResourcesForSectionParamsWithTimeout(timeout time.Duration) *GetResourcesForSectionParams {
	return &GetResourcesForSectionParams{
		timeout: timeout,
	}
}

// NewGetResourcesForSectionParamsWithContext creates a new GetResourcesForSectionParams object
// with the ability to set a context for a request.
func NewGetResourcesForSectionParamsWithContext(ctx context.Context) *GetResourcesForSectionParams {
	return &GetResourcesForSectionParams{
		Context: ctx,
	}
}

// NewGetResourcesForSectionParamsWithHTTPClient creates a new GetResourcesForSectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourcesForSectionParamsWithHTTPClient(client *http.Client) *GetResourcesForSectionParams {
	return &GetResourcesForSectionParams{
		HTTPClient: client,
	}
}

/* GetResourcesForSectionParams contains all the parameters to send to the API endpoint
   for the get resources for section operation.

   Typically these are written to a http.Request.
*/
type GetResourcesForSectionParams struct {

	// EndingBefore.
	EndingBefore *string

	// ID.
	ID string

	// Limit.
	Limit *int64

	// StartingAfter.
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resources for section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourcesForSectionParams) WithDefaults() *GetResourcesForSectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resources for section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourcesForSectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resources for section params
func (o *GetResourcesForSectionParams) WithTimeout(timeout time.Duration) *GetResourcesForSectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resources for section params
func (o *GetResourcesForSectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resources for section params
func (o *GetResourcesForSectionParams) WithContext(ctx context.Context) *GetResourcesForSectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resources for section params
func (o *GetResourcesForSectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resources for section params
func (o *GetResourcesForSectionParams) WithHTTPClient(client *http.Client) *GetResourcesForSectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resources for section params
func (o *GetResourcesForSectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get resources for section params
func (o *GetResourcesForSectionParams) WithEndingBefore(endingBefore *string) *GetResourcesForSectionParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get resources for section params
func (o *GetResourcesForSectionParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithID adds the id to the get resources for section params
func (o *GetResourcesForSectionParams) WithID(id string) *GetResourcesForSectionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get resources for section params
func (o *GetResourcesForSectionParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the get resources for section params
func (o *GetResourcesForSectionParams) WithLimit(limit *int64) *GetResourcesForSectionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get resources for section params
func (o *GetResourcesForSectionParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStartingAfter adds the startingAfter to the get resources for section params
func (o *GetResourcesForSectionParams) WithStartingAfter(startingAfter *string) *GetResourcesForSectionParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get resources for section params
func (o *GetResourcesForSectionParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourcesForSectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingBefore != nil {

		// query param ending_before
		var qrEndingBefore string

		if o.EndingBefore != nil {
			qrEndingBefore = *o.EndingBefore
		}
		qEndingBefore := qrEndingBefore
		if qEndingBefore != "" {

			if err := r.SetQueryParam("ending_before", qEndingBefore); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.StartingAfter != nil {

		// query param starting_after
		var qrStartingAfter string

		if o.StartingAfter != nil {
			qrStartingAfter = *o.StartingAfter
		}
		qStartingAfter := qrStartingAfter
		if qStartingAfter != "" {

			if err := r.SetQueryParam("starting_after", qStartingAfter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
