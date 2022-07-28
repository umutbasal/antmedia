// Code generated by go-swagger;

package operations

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

// NewGetClusterDeleteNodeIDParams creates a new GetClusterDeleteNodeIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterDeleteNodeIDParams() *GetClusterDeleteNodeIDParams {
	return &GetClusterDeleteNodeIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterDeleteNodeIDParamsWithTimeout creates a new GetClusterDeleteNodeIDParams object
// with the ability to set a timeout on a request.
func NewGetClusterDeleteNodeIDParamsWithTimeout(timeout time.Duration) *GetClusterDeleteNodeIDParams {
	return &GetClusterDeleteNodeIDParams{
		timeout: timeout,
	}
}

// NewGetClusterDeleteNodeIDParamsWithContext creates a new GetClusterDeleteNodeIDParams object
// with the ability to set a context for a request.
func NewGetClusterDeleteNodeIDParamsWithContext(ctx context.Context) *GetClusterDeleteNodeIDParams {
	return &GetClusterDeleteNodeIDParams{
		Context: ctx,
	}
}

// NewGetClusterDeleteNodeIDParamsWithHTTPClient creates a new GetClusterDeleteNodeIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterDeleteNodeIDParamsWithHTTPClient(client *http.Client) *GetClusterDeleteNodeIDParams {
	return &GetClusterDeleteNodeIDParams{
		HTTPClient: client,
	}
}

/* GetClusterDeleteNodeIDParams contains all the parameters to send to the API endpoint
   for the get cluster delete node ID operation.

   Typically these are written to a http.Request.
*/
type GetClusterDeleteNodeIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster delete node ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterDeleteNodeIDParams) WithDefaults() *GetClusterDeleteNodeIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster delete node ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterDeleteNodeIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster delete node ID params
func (o *GetClusterDeleteNodeIDParams) WithTimeout(timeout time.Duration) *GetClusterDeleteNodeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster delete node ID params
func (o *GetClusterDeleteNodeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster delete node ID params
func (o *GetClusterDeleteNodeIDParams) WithContext(ctx context.Context) *GetClusterDeleteNodeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster delete node ID params
func (o *GetClusterDeleteNodeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster delete node ID params
func (o *GetClusterDeleteNodeIDParams) WithHTTPClient(client *http.Client) *GetClusterDeleteNodeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster delete node ID params
func (o *GetClusterDeleteNodeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get cluster delete node ID params
func (o *GetClusterDeleteNodeIDParams) WithID(id string) *GetClusterDeleteNodeIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get cluster delete node ID params
func (o *GetClusterDeleteNodeIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterDeleteNodeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
