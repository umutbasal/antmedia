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

// NewResetBroadcastParams creates a new ResetBroadcastParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetBroadcastParams() *ResetBroadcastParams {
	return &ResetBroadcastParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetBroadcastParamsWithTimeout creates a new ResetBroadcastParams object
// with the ability to set a timeout on a request.
func NewResetBroadcastParamsWithTimeout(timeout time.Duration) *ResetBroadcastParams {
	return &ResetBroadcastParams{
		timeout: timeout,
	}
}

// NewResetBroadcastParamsWithContext creates a new ResetBroadcastParams object
// with the ability to set a context for a request.
func NewResetBroadcastParamsWithContext(ctx context.Context) *ResetBroadcastParams {
	return &ResetBroadcastParams{
		Context: ctx,
	}
}

// NewResetBroadcastParamsWithHTTPClient creates a new ResetBroadcastParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetBroadcastParamsWithHTTPClient(client *http.Client) *ResetBroadcastParams {
	return &ResetBroadcastParams{
		HTTPClient: client,
	}
}

/* ResetBroadcastParams contains all the parameters to send to the API endpoint
   for the reset broadcast operation.

   Typically these are written to a http.Request.
*/
type ResetBroadcastParams struct {

	/* Appname.

	   Application name
	*/
	Appname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset broadcast params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetBroadcastParams) WithDefaults() *ResetBroadcastParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset broadcast params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetBroadcastParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset broadcast params
func (o *ResetBroadcastParams) WithTimeout(timeout time.Duration) *ResetBroadcastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset broadcast params
func (o *ResetBroadcastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset broadcast params
func (o *ResetBroadcastParams) WithContext(ctx context.Context) *ResetBroadcastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset broadcast params
func (o *ResetBroadcastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset broadcast params
func (o *ResetBroadcastParams) WithHTTPClient(client *http.Client) *ResetBroadcastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset broadcast params
func (o *ResetBroadcastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppname adds the appname to the reset broadcast params
func (o *ResetBroadcastParams) WithAppname(appname string) *ResetBroadcastParams {
	o.SetAppname(appname)
	return o
}

// SetAppname adds the appname to the reset broadcast params
func (o *ResetBroadcastParams) SetAppname(appname string) {
	o.Appname = appname
}

// WriteToRequest writes these params to a swagger request
func (o *ResetBroadcastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appname
	if err := r.SetPathParam("appname", o.Appname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
