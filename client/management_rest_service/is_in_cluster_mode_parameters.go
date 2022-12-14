// Code generated by go-swagger;

package management_rest_service

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

// NewIsInClusterModeParams creates a new IsInClusterModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIsInClusterModeParams() *IsInClusterModeParams {
	return &IsInClusterModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIsInClusterModeParamsWithTimeout creates a new IsInClusterModeParams object
// with the ability to set a timeout on a request.
func NewIsInClusterModeParamsWithTimeout(timeout time.Duration) *IsInClusterModeParams {
	return &IsInClusterModeParams{
		timeout: timeout,
	}
}

// NewIsInClusterModeParamsWithContext creates a new IsInClusterModeParams object
// with the ability to set a context for a request.
func NewIsInClusterModeParamsWithContext(ctx context.Context) *IsInClusterModeParams {
	return &IsInClusterModeParams{
		Context: ctx,
	}
}

// NewIsInClusterModeParamsWithHTTPClient creates a new IsInClusterModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewIsInClusterModeParamsWithHTTPClient(client *http.Client) *IsInClusterModeParams {
	return &IsInClusterModeParams{
		HTTPClient: client,
	}
}

/* IsInClusterModeParams contains all the parameters to send to the API endpoint
   for the is in cluster mode operation.

   Typically these are written to a http.Request.
*/
type IsInClusterModeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the is in cluster mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IsInClusterModeParams) WithDefaults() *IsInClusterModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the is in cluster mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IsInClusterModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the is in cluster mode params
func (o *IsInClusterModeParams) WithTimeout(timeout time.Duration) *IsInClusterModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is in cluster mode params
func (o *IsInClusterModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is in cluster mode params
func (o *IsInClusterModeParams) WithContext(ctx context.Context) *IsInClusterModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is in cluster mode params
func (o *IsInClusterModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is in cluster mode params
func (o *IsInClusterModeParams) WithHTTPClient(client *http.Client) *IsInClusterModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is in cluster mode params
func (o *IsInClusterModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IsInClusterModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
