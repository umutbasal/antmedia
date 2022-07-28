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

	"antmedia/models"
)

// NewCreateConferenceRoomV2Params creates a new CreateConferenceRoomV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateConferenceRoomV2Params() *CreateConferenceRoomV2Params {
	return &CreateConferenceRoomV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateConferenceRoomV2ParamsWithTimeout creates a new CreateConferenceRoomV2Params object
// with the ability to set a timeout on a request.
func NewCreateConferenceRoomV2ParamsWithTimeout(timeout time.Duration) *CreateConferenceRoomV2Params {
	return &CreateConferenceRoomV2Params{
		timeout: timeout,
	}
}

// NewCreateConferenceRoomV2ParamsWithContext creates a new CreateConferenceRoomV2Params object
// with the ability to set a context for a request.
func NewCreateConferenceRoomV2ParamsWithContext(ctx context.Context) *CreateConferenceRoomV2Params {
	return &CreateConferenceRoomV2Params{
		Context: ctx,
	}
}

// NewCreateConferenceRoomV2ParamsWithHTTPClient creates a new CreateConferenceRoomV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateConferenceRoomV2ParamsWithHTTPClient(client *http.Client) *CreateConferenceRoomV2Params {
	return &CreateConferenceRoomV2Params{
		HTTPClient: client,
	}
}

/* CreateConferenceRoomV2Params contains all the parameters to send to the API endpoint
   for the create conference room v2 operation.

   Typically these are written to a http.Request.
*/
type CreateConferenceRoomV2Params struct {

	/* Body.

	   Conference Room object with start and end date
	*/
	Body *models.ConferenceRoom

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create conference room v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateConferenceRoomV2Params) WithDefaults() *CreateConferenceRoomV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create conference room v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateConferenceRoomV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create conference room v2 params
func (o *CreateConferenceRoomV2Params) WithTimeout(timeout time.Duration) *CreateConferenceRoomV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create conference room v2 params
func (o *CreateConferenceRoomV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create conference room v2 params
func (o *CreateConferenceRoomV2Params) WithContext(ctx context.Context) *CreateConferenceRoomV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create conference room v2 params
func (o *CreateConferenceRoomV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create conference room v2 params
func (o *CreateConferenceRoomV2Params) WithHTTPClient(client *http.Client) *CreateConferenceRoomV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create conference room v2 params
func (o *CreateConferenceRoomV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create conference room v2 params
func (o *CreateConferenceRoomV2Params) WithBody(body *models.ConferenceRoom) *CreateConferenceRoomV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create conference room v2 params
func (o *CreateConferenceRoomV2Params) SetBody(body *models.ConferenceRoom) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateConferenceRoomV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
