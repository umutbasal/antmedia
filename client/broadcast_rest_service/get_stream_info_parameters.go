// Code generated by go-swagger;

package broadcast_rest_service

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

// NewGetStreamInfoParams creates a new GetStreamInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStreamInfoParams() *GetStreamInfoParams {
	return &GetStreamInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStreamInfoParamsWithTimeout creates a new GetStreamInfoParams object
// with the ability to set a timeout on a request.
func NewGetStreamInfoParamsWithTimeout(timeout time.Duration) *GetStreamInfoParams {
	return &GetStreamInfoParams{
		timeout: timeout,
	}
}

// NewGetStreamInfoParamsWithContext creates a new GetStreamInfoParams object
// with the ability to set a context for a request.
func NewGetStreamInfoParamsWithContext(ctx context.Context) *GetStreamInfoParams {
	return &GetStreamInfoParams{
		Context: ctx,
	}
}

// NewGetStreamInfoParamsWithHTTPClient creates a new GetStreamInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStreamInfoParamsWithHTTPClient(client *http.Client) *GetStreamInfoParams {
	return &GetStreamInfoParams{
		HTTPClient: client,
	}
}

/* GetStreamInfoParams contains all the parameters to send to the API endpoint
   for the get stream info operation.

   Typically these are written to a http.Request.
*/
type GetStreamInfoParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get stream info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStreamInfoParams) WithDefaults() *GetStreamInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get stream info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStreamInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get stream info params
func (o *GetStreamInfoParams) WithTimeout(timeout time.Duration) *GetStreamInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stream info params
func (o *GetStreamInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stream info params
func (o *GetStreamInfoParams) WithContext(ctx context.Context) *GetStreamInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stream info params
func (o *GetStreamInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stream info params
func (o *GetStreamInfoParams) WithHTTPClient(client *http.Client) *GetStreamInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stream info params
func (o *GetStreamInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get stream info params
func (o *GetStreamInfoParams) WithID(id string) *GetStreamInfoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get stream info params
func (o *GetStreamInfoParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStreamInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
