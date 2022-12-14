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

// NewGetServerTimeParams creates a new GetServerTimeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerTimeParams() *GetServerTimeParams {
	return &GetServerTimeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerTimeParamsWithTimeout creates a new GetServerTimeParams object
// with the ability to set a timeout on a request.
func NewGetServerTimeParamsWithTimeout(timeout time.Duration) *GetServerTimeParams {
	return &GetServerTimeParams{
		timeout: timeout,
	}
}

// NewGetServerTimeParamsWithContext creates a new GetServerTimeParams object
// with the ability to set a context for a request.
func NewGetServerTimeParamsWithContext(ctx context.Context) *GetServerTimeParams {
	return &GetServerTimeParams{
		Context: ctx,
	}
}

// NewGetServerTimeParamsWithHTTPClient creates a new GetServerTimeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerTimeParamsWithHTTPClient(client *http.Client) *GetServerTimeParams {
	return &GetServerTimeParams{
		HTTPClient: client,
	}
}

/* GetServerTimeParams contains all the parameters to send to the API endpoint
   for the get server time operation.

   Typically these are written to a http.Request.
*/
type GetServerTimeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerTimeParams) WithDefaults() *GetServerTimeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerTimeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get server time params
func (o *GetServerTimeParams) WithTimeout(timeout time.Duration) *GetServerTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server time params
func (o *GetServerTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server time params
func (o *GetServerTimeParams) WithContext(ctx context.Context) *GetServerTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server time params
func (o *GetServerTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server time params
func (o *GetServerTimeParams) WithHTTPClient(client *http.Client) *GetServerTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server time params
func (o *GetServerTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
