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

// NewRevokeSubscribersParams creates a new RevokeSubscribersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeSubscribersParams() *RevokeSubscribersParams {
	return &RevokeSubscribersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeSubscribersParamsWithTimeout creates a new RevokeSubscribersParams object
// with the ability to set a timeout on a request.
func NewRevokeSubscribersParamsWithTimeout(timeout time.Duration) *RevokeSubscribersParams {
	return &RevokeSubscribersParams{
		timeout: timeout,
	}
}

// NewRevokeSubscribersParamsWithContext creates a new RevokeSubscribersParams object
// with the ability to set a context for a request.
func NewRevokeSubscribersParamsWithContext(ctx context.Context) *RevokeSubscribersParams {
	return &RevokeSubscribersParams{
		Context: ctx,
	}
}

// NewRevokeSubscribersParamsWithHTTPClient creates a new RevokeSubscribersParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeSubscribersParamsWithHTTPClient(client *http.Client) *RevokeSubscribersParams {
	return &RevokeSubscribersParams{
		HTTPClient: client,
	}
}

/* RevokeSubscribersParams contains all the parameters to send to the API endpoint
   for the revoke subscribers operation.

   Typically these are written to a http.Request.
*/
type RevokeSubscribersParams struct {

	/* ID.

	   the id of the stream
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke subscribers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeSubscribersParams) WithDefaults() *RevokeSubscribersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke subscribers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeSubscribersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke subscribers params
func (o *RevokeSubscribersParams) WithTimeout(timeout time.Duration) *RevokeSubscribersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke subscribers params
func (o *RevokeSubscribersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke subscribers params
func (o *RevokeSubscribersParams) WithContext(ctx context.Context) *RevokeSubscribersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke subscribers params
func (o *RevokeSubscribersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke subscribers params
func (o *RevokeSubscribersParams) WithHTTPClient(client *http.Client) *RevokeSubscribersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke subscribers params
func (o *RevokeSubscribersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the revoke subscribers params
func (o *RevokeSubscribersParams) WithID(id string) *RevokeSubscribersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the revoke subscribers params
func (o *RevokeSubscribersParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeSubscribersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
