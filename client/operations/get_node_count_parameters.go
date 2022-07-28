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

// NewGetNodeCountParams creates a new GetNodeCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNodeCountParams() *GetNodeCountParams {
	return &GetNodeCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeCountParamsWithTimeout creates a new GetNodeCountParams object
// with the ability to set a timeout on a request.
func NewGetNodeCountParamsWithTimeout(timeout time.Duration) *GetNodeCountParams {
	return &GetNodeCountParams{
		timeout: timeout,
	}
}

// NewGetNodeCountParamsWithContext creates a new GetNodeCountParams object
// with the ability to set a context for a request.
func NewGetNodeCountParamsWithContext(ctx context.Context) *GetNodeCountParams {
	return &GetNodeCountParams{
		Context: ctx,
	}
}

// NewGetNodeCountParamsWithHTTPClient creates a new GetNodeCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNodeCountParamsWithHTTPClient(client *http.Client) *GetNodeCountParams {
	return &GetNodeCountParams{
		HTTPClient: client,
	}
}

/* GetNodeCountParams contains all the parameters to send to the API endpoint
   for the get node count operation.

   Typically these are written to a http.Request.
*/
type GetNodeCountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get node count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodeCountParams) WithDefaults() *GetNodeCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get node count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodeCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get node count params
func (o *GetNodeCountParams) WithTimeout(timeout time.Duration) *GetNodeCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node count params
func (o *GetNodeCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node count params
func (o *GetNodeCountParams) WithContext(ctx context.Context) *GetNodeCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node count params
func (o *GetNodeCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node count params
func (o *GetNodeCountParams) WithHTTPClient(client *http.Client) *GetNodeCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node count params
func (o *GetNodeCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
