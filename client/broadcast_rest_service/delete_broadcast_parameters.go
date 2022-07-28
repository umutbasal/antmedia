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

// NewDeleteBroadcastParams creates a new DeleteBroadcastParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBroadcastParams() *DeleteBroadcastParams {
	return &DeleteBroadcastParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBroadcastParamsWithTimeout creates a new DeleteBroadcastParams object
// with the ability to set a timeout on a request.
func NewDeleteBroadcastParamsWithTimeout(timeout time.Duration) *DeleteBroadcastParams {
	return &DeleteBroadcastParams{
		timeout: timeout,
	}
}

// NewDeleteBroadcastParamsWithContext creates a new DeleteBroadcastParams object
// with the ability to set a context for a request.
func NewDeleteBroadcastParamsWithContext(ctx context.Context) *DeleteBroadcastParams {
	return &DeleteBroadcastParams{
		Context: ctx,
	}
}

// NewDeleteBroadcastParamsWithHTTPClient creates a new DeleteBroadcastParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBroadcastParamsWithHTTPClient(client *http.Client) *DeleteBroadcastParams {
	return &DeleteBroadcastParams{
		HTTPClient: client,
	}
}

/* DeleteBroadcastParams contains all the parameters to send to the API endpoint
   for the delete broadcast operation.

   Typically these are written to a http.Request.
*/
type DeleteBroadcastParams struct {

	/* ID.

	   Id of the broadcast
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete broadcast params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBroadcastParams) WithDefaults() *DeleteBroadcastParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete broadcast params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBroadcastParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete broadcast params
func (o *DeleteBroadcastParams) WithTimeout(timeout time.Duration) *DeleteBroadcastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete broadcast params
func (o *DeleteBroadcastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete broadcast params
func (o *DeleteBroadcastParams) WithContext(ctx context.Context) *DeleteBroadcastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete broadcast params
func (o *DeleteBroadcastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete broadcast params
func (o *DeleteBroadcastParams) WithHTTPClient(client *http.Client) *DeleteBroadcastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete broadcast params
func (o *DeleteBroadcastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete broadcast params
func (o *DeleteBroadcastParams) WithID(id string) *DeleteBroadcastParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete broadcast params
func (o *DeleteBroadcastParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBroadcastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
