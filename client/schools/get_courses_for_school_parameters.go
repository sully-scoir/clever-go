// Code generated by go-swagger; DO NOT EDIT.

package schools

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

// NewGetCoursesForSchoolParams creates a new GetCoursesForSchoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCoursesForSchoolParams() *GetCoursesForSchoolParams {
	return &GetCoursesForSchoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCoursesForSchoolParamsWithTimeout creates a new GetCoursesForSchoolParams object
// with the ability to set a timeout on a request.
func NewGetCoursesForSchoolParamsWithTimeout(timeout time.Duration) *GetCoursesForSchoolParams {
	return &GetCoursesForSchoolParams{
		timeout: timeout,
	}
}

// NewGetCoursesForSchoolParamsWithContext creates a new GetCoursesForSchoolParams object
// with the ability to set a context for a request.
func NewGetCoursesForSchoolParamsWithContext(ctx context.Context) *GetCoursesForSchoolParams {
	return &GetCoursesForSchoolParams{
		Context: ctx,
	}
}

// NewGetCoursesForSchoolParamsWithHTTPClient creates a new GetCoursesForSchoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCoursesForSchoolParamsWithHTTPClient(client *http.Client) *GetCoursesForSchoolParams {
	return &GetCoursesForSchoolParams{
		HTTPClient: client,
	}
}

/* GetCoursesForSchoolParams contains all the parameters to send to the API endpoint
   for the get courses for school operation.

   Typically these are written to a http.Request.
*/
type GetCoursesForSchoolParams struct {

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

// WithDefaults hydrates default values in the get courses for school params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoursesForSchoolParams) WithDefaults() *GetCoursesForSchoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get courses for school params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoursesForSchoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get courses for school params
func (o *GetCoursesForSchoolParams) WithTimeout(timeout time.Duration) *GetCoursesForSchoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get courses for school params
func (o *GetCoursesForSchoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get courses for school params
func (o *GetCoursesForSchoolParams) WithContext(ctx context.Context) *GetCoursesForSchoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get courses for school params
func (o *GetCoursesForSchoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get courses for school params
func (o *GetCoursesForSchoolParams) WithHTTPClient(client *http.Client) *GetCoursesForSchoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get courses for school params
func (o *GetCoursesForSchoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get courses for school params
func (o *GetCoursesForSchoolParams) WithEndingBefore(endingBefore *string) *GetCoursesForSchoolParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get courses for school params
func (o *GetCoursesForSchoolParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithID adds the id to the get courses for school params
func (o *GetCoursesForSchoolParams) WithID(id string) *GetCoursesForSchoolParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get courses for school params
func (o *GetCoursesForSchoolParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the get courses for school params
func (o *GetCoursesForSchoolParams) WithLimit(limit *int64) *GetCoursesForSchoolParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get courses for school params
func (o *GetCoursesForSchoolParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStartingAfter adds the startingAfter to the get courses for school params
func (o *GetCoursesForSchoolParams) WithStartingAfter(startingAfter *string) *GetCoursesForSchoolParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get courses for school params
func (o *GetCoursesForSchoolParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetCoursesForSchoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
