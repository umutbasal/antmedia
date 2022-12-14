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

// NewGetLicenceStatusParams creates a new GetLicenceStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLicenceStatusParams() *GetLicenceStatusParams {
	return &GetLicenceStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenceStatusParamsWithTimeout creates a new GetLicenceStatusParams object
// with the ability to set a timeout on a request.
func NewGetLicenceStatusParamsWithTimeout(timeout time.Duration) *GetLicenceStatusParams {
	return &GetLicenceStatusParams{
		timeout: timeout,
	}
}

// NewGetLicenceStatusParamsWithContext creates a new GetLicenceStatusParams object
// with the ability to set a context for a request.
func NewGetLicenceStatusParamsWithContext(ctx context.Context) *GetLicenceStatusParams {
	return &GetLicenceStatusParams{
		Context: ctx,
	}
}

// NewGetLicenceStatusParamsWithHTTPClient creates a new GetLicenceStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLicenceStatusParamsWithHTTPClient(client *http.Client) *GetLicenceStatusParams {
	return &GetLicenceStatusParams{
		HTTPClient: client,
	}
}

/* GetLicenceStatusParams contains all the parameters to send to the API endpoint
   for the get licence status operation.

   Typically these are written to a http.Request.
*/
type GetLicenceStatusParams struct {

	/* Key.

	   License key
	*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get licence status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLicenceStatusParams) WithDefaults() *GetLicenceStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get licence status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLicenceStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get licence status params
func (o *GetLicenceStatusParams) WithTimeout(timeout time.Duration) *GetLicenceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get licence status params
func (o *GetLicenceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get licence status params
func (o *GetLicenceStatusParams) WithContext(ctx context.Context) *GetLicenceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get licence status params
func (o *GetLicenceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get licence status params
func (o *GetLicenceStatusParams) WithHTTPClient(client *http.Client) *GetLicenceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get licence status params
func (o *GetLicenceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the get licence status params
func (o *GetLicenceStatusParams) WithKey(key string) *GetLicenceStatusParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the get licence status params
func (o *GetLicenceStatusParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param key
	qrKey := o.Key
	qKey := qrKey
	if qKey != "" {

		if err := r.SetQueryParam("key", qKey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
