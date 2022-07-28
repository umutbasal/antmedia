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

	"github.com/umutbasal/antmedia/models"
)

// NewUpdateBroadcastParams creates a new UpdateBroadcastParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBroadcastParams() *UpdateBroadcastParams {
	return &UpdateBroadcastParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBroadcastParamsWithTimeout creates a new UpdateBroadcastParams object
// with the ability to set a timeout on a request.
func NewUpdateBroadcastParamsWithTimeout(timeout time.Duration) *UpdateBroadcastParams {
	return &UpdateBroadcastParams{
		timeout: timeout,
	}
}

// NewUpdateBroadcastParamsWithContext creates a new UpdateBroadcastParams object
// with the ability to set a context for a request.
func NewUpdateBroadcastParamsWithContext(ctx context.Context) *UpdateBroadcastParams {
	return &UpdateBroadcastParams{
		Context: ctx,
	}
}

// NewUpdateBroadcastParamsWithHTTPClient creates a new UpdateBroadcastParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBroadcastParamsWithHTTPClient(client *http.Client) *UpdateBroadcastParams {
	return &UpdateBroadcastParams{
		HTTPClient: client,
	}
}

/* UpdateBroadcastParams contains all the parameters to send to the API endpoint
   for the update broadcast operation.

   Typically these are written to a http.Request.
*/
type UpdateBroadcastParams struct {

	/* Body.

	   Broadcast object with the updates
	*/
	Body *models.Broadcast

	/* ID.

	   Broadcast id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update broadcast params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBroadcastParams) WithDefaults() *UpdateBroadcastParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update broadcast params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBroadcastParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update broadcast params
func (o *UpdateBroadcastParams) WithTimeout(timeout time.Duration) *UpdateBroadcastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update broadcast params
func (o *UpdateBroadcastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update broadcast params
func (o *UpdateBroadcastParams) WithContext(ctx context.Context) *UpdateBroadcastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update broadcast params
func (o *UpdateBroadcastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update broadcast params
func (o *UpdateBroadcastParams) WithHTTPClient(client *http.Client) *UpdateBroadcastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update broadcast params
func (o *UpdateBroadcastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update broadcast params
func (o *UpdateBroadcastParams) WithBody(body *models.Broadcast) *UpdateBroadcastParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update broadcast params
func (o *UpdateBroadcastParams) SetBody(body *models.Broadcast) {
	o.Body = body
}

// WithID adds the id to the update broadcast params
func (o *UpdateBroadcastParams) WithID(id string) *UpdateBroadcastParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update broadcast params
func (o *UpdateBroadcastParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBroadcastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
