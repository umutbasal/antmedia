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

	"antmedia/models"
)

// NewEditConferenceRoomParams creates a new EditConferenceRoomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEditConferenceRoomParams() *EditConferenceRoomParams {
	return &EditConferenceRoomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEditConferenceRoomParamsWithTimeout creates a new EditConferenceRoomParams object
// with the ability to set a timeout on a request.
func NewEditConferenceRoomParamsWithTimeout(timeout time.Duration) *EditConferenceRoomParams {
	return &EditConferenceRoomParams{
		timeout: timeout,
	}
}

// NewEditConferenceRoomParamsWithContext creates a new EditConferenceRoomParams object
// with the ability to set a context for a request.
func NewEditConferenceRoomParamsWithContext(ctx context.Context) *EditConferenceRoomParams {
	return &EditConferenceRoomParams{
		Context: ctx,
	}
}

// NewEditConferenceRoomParamsWithHTTPClient creates a new EditConferenceRoomParams object
// with the ability to set a custom HTTPClient for a request.
func NewEditConferenceRoomParamsWithHTTPClient(client *http.Client) *EditConferenceRoomParams {
	return &EditConferenceRoomParams{
		HTTPClient: client,
	}
}

/* EditConferenceRoomParams contains all the parameters to send to the API endpoint
   for the edit conference room operation.

   Typically these are written to a http.Request.
*/
type EditConferenceRoomParams struct {

	/* Body.

	   Conference Room object with start and end date
	*/
	Body *models.ConferenceRoom

	/* RoomID.

	   Room id
	*/
	RoomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edit conference room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditConferenceRoomParams) WithDefaults() *EditConferenceRoomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edit conference room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditConferenceRoomParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edit conference room params
func (o *EditConferenceRoomParams) WithTimeout(timeout time.Duration) *EditConferenceRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit conference room params
func (o *EditConferenceRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit conference room params
func (o *EditConferenceRoomParams) WithContext(ctx context.Context) *EditConferenceRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit conference room params
func (o *EditConferenceRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit conference room params
func (o *EditConferenceRoomParams) WithHTTPClient(client *http.Client) *EditConferenceRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit conference room params
func (o *EditConferenceRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit conference room params
func (o *EditConferenceRoomParams) WithBody(body *models.ConferenceRoom) *EditConferenceRoomParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit conference room params
func (o *EditConferenceRoomParams) SetBody(body *models.ConferenceRoom) {
	o.Body = body
}

// WithRoomID adds the roomID to the edit conference room params
func (o *EditConferenceRoomParams) WithRoomID(roomID string) *EditConferenceRoomParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the edit conference room params
func (o *EditConferenceRoomParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WriteToRequest writes these params to a swagger request
func (o *EditConferenceRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param room_id
	if err := r.SetPathParam("room_id", o.RoomID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
