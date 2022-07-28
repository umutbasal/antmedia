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
	"github.com/go-openapi/swag"
)

// NewListTokensV2Params creates a new ListTokensV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTokensV2Params() *ListTokensV2Params {
	return &ListTokensV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTokensV2ParamsWithTimeout creates a new ListTokensV2Params object
// with the ability to set a timeout on a request.
func NewListTokensV2ParamsWithTimeout(timeout time.Duration) *ListTokensV2Params {
	return &ListTokensV2Params{
		timeout: timeout,
	}
}

// NewListTokensV2ParamsWithContext creates a new ListTokensV2Params object
// with the ability to set a context for a request.
func NewListTokensV2ParamsWithContext(ctx context.Context) *ListTokensV2Params {
	return &ListTokensV2Params{
		Context: ctx,
	}
}

// NewListTokensV2ParamsWithHTTPClient creates a new ListTokensV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListTokensV2ParamsWithHTTPClient(client *http.Client) *ListTokensV2Params {
	return &ListTokensV2Params{
		HTTPClient: client,
	}
}

/* ListTokensV2Params contains all the parameters to send to the API endpoint
   for the list tokens v2 operation.

   Typically these are written to a http.Request.
*/
type ListTokensV2Params struct {

	/* ID.

	   the id of the stream
	*/
	ID string

	/* Offset.

	   the starting point of the list

	   Format: int32
	*/
	Offset int32

	/* Size.

	   size of the return list (max:50 )

	   Format: int32
	*/
	Size int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list tokens v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTokensV2Params) WithDefaults() *ListTokensV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list tokens v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTokensV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list tokens v2 params
func (o *ListTokensV2Params) WithTimeout(timeout time.Duration) *ListTokensV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tokens v2 params
func (o *ListTokensV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tokens v2 params
func (o *ListTokensV2Params) WithContext(ctx context.Context) *ListTokensV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tokens v2 params
func (o *ListTokensV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tokens v2 params
func (o *ListTokensV2Params) WithHTTPClient(client *http.Client) *ListTokensV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tokens v2 params
func (o *ListTokensV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list tokens v2 params
func (o *ListTokensV2Params) WithID(id string) *ListTokensV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the list tokens v2 params
func (o *ListTokensV2Params) SetID(id string) {
	o.ID = id
}

// WithOffset adds the offset to the list tokens v2 params
func (o *ListTokensV2Params) WithOffset(offset int32) *ListTokensV2Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list tokens v2 params
func (o *ListTokensV2Params) SetOffset(offset int32) {
	o.Offset = offset
}

// WithSize adds the size to the list tokens v2 params
func (o *ListTokensV2Params) WithSize(size int32) *ListTokensV2Params {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the list tokens v2 params
func (o *ListTokensV2Params) SetSize(size int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *ListTokensV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param offset
	if err := r.SetPathParam("offset", swag.FormatInt32(o.Offset)); err != nil {
		return err
	}

	// path param size
	if err := r.SetPathParam("size", swag.FormatInt32(o.Size)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
