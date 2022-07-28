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
)

// NewSendMessageParams creates a new SendMessageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSendMessageParams() *SendMessageParams {
	return &SendMessageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSendMessageParamsWithTimeout creates a new SendMessageParams object
// with the ability to set a timeout on a request.
func NewSendMessageParamsWithTimeout(timeout time.Duration) *SendMessageParams {
	return &SendMessageParams{
		timeout: timeout,
	}
}

// NewSendMessageParamsWithContext creates a new SendMessageParams object
// with the ability to set a context for a request.
func NewSendMessageParamsWithContext(ctx context.Context) *SendMessageParams {
	return &SendMessageParams{
		Context: ctx,
	}
}

// NewSendMessageParamsWithHTTPClient creates a new SendMessageParams object
// with the ability to set a custom HTTPClient for a request.
func NewSendMessageParamsWithHTTPClient(client *http.Client) *SendMessageParams {
	return &SendMessageParams{
		HTTPClient: client,
	}
}

/* SendMessageParams contains all the parameters to send to the API endpoint
   for the send message operation.

   Typically these are written to a http.Request.
*/
type SendMessageParams struct {

	/* Body.

	   Message through Data Channel which will be sent to all WebRTC stream participants
	*/
	Body string

	/* ID.

	   Broadcast id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the send message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendMessageParams) WithDefaults() *SendMessageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the send message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendMessageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the send message params
func (o *SendMessageParams) WithTimeout(timeout time.Duration) *SendMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send message params
func (o *SendMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send message params
func (o *SendMessageParams) WithContext(ctx context.Context) *SendMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send message params
func (o *SendMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send message params
func (o *SendMessageParams) WithHTTPClient(client *http.Client) *SendMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send message params
func (o *SendMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send message params
func (o *SendMessageParams) WithBody(body string) *SendMessageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send message params
func (o *SendMessageParams) SetBody(body string) {
	o.Body = body
}

// WithID adds the id to the send message params
func (o *SendMessageParams) WithID(id string) *SendMessageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the send message params
func (o *SendMessageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SendMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
