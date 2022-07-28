// Code generated by go-swagger; DO NOT EDIT.

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

// NewListSubscriberV2Params creates a new ListSubscriberV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSubscriberV2Params() *ListSubscriberV2Params {
	return &ListSubscriberV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSubscriberV2ParamsWithTimeout creates a new ListSubscriberV2Params object
// with the ability to set a timeout on a request.
func NewListSubscriberV2ParamsWithTimeout(timeout time.Duration) *ListSubscriberV2Params {
	return &ListSubscriberV2Params{
		timeout: timeout,
	}
}

// NewListSubscriberV2ParamsWithContext creates a new ListSubscriberV2Params object
// with the ability to set a context for a request.
func NewListSubscriberV2ParamsWithContext(ctx context.Context) *ListSubscriberV2Params {
	return &ListSubscriberV2Params{
		Context: ctx,
	}
}

// NewListSubscriberV2ParamsWithHTTPClient creates a new ListSubscriberV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListSubscriberV2ParamsWithHTTPClient(client *http.Client) *ListSubscriberV2Params {
	return &ListSubscriberV2Params{
		HTTPClient: client,
	}
}

/* ListSubscriberV2Params contains all the parameters to send to the API endpoint
   for the list subscriber v2 operation.

   Typically these are written to a http.Request.
*/
type ListSubscriberV2Params struct {

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

// WithDefaults hydrates default values in the list subscriber v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSubscriberV2Params) WithDefaults() *ListSubscriberV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list subscriber v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSubscriberV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list subscriber v2 params
func (o *ListSubscriberV2Params) WithTimeout(timeout time.Duration) *ListSubscriberV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list subscriber v2 params
func (o *ListSubscriberV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list subscriber v2 params
func (o *ListSubscriberV2Params) WithContext(ctx context.Context) *ListSubscriberV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list subscriber v2 params
func (o *ListSubscriberV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list subscriber v2 params
func (o *ListSubscriberV2Params) WithHTTPClient(client *http.Client) *ListSubscriberV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list subscriber v2 params
func (o *ListSubscriberV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list subscriber v2 params
func (o *ListSubscriberV2Params) WithID(id string) *ListSubscriberV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the list subscriber v2 params
func (o *ListSubscriberV2Params) SetID(id string) {
	o.ID = id
}

// WithOffset adds the offset to the list subscriber v2 params
func (o *ListSubscriberV2Params) WithOffset(offset int32) *ListSubscriberV2Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list subscriber v2 params
func (o *ListSubscriberV2Params) SetOffset(offset int32) {
	o.Offset = offset
}

// WithSize adds the size to the list subscriber v2 params
func (o *ListSubscriberV2Params) WithSize(size int32) *ListSubscriberV2Params {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the list subscriber v2 params
func (o *ListSubscriberV2Params) SetSize(size int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *ListSubscriberV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
