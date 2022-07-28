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

// NewAddStreamToTheRoomParams creates a new AddStreamToTheRoomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddStreamToTheRoomParams() *AddStreamToTheRoomParams {
	return &AddStreamToTheRoomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddStreamToTheRoomParamsWithTimeout creates a new AddStreamToTheRoomParams object
// with the ability to set a timeout on a request.
func NewAddStreamToTheRoomParamsWithTimeout(timeout time.Duration) *AddStreamToTheRoomParams {
	return &AddStreamToTheRoomParams{
		timeout: timeout,
	}
}

// NewAddStreamToTheRoomParamsWithContext creates a new AddStreamToTheRoomParams object
// with the ability to set a context for a request.
func NewAddStreamToTheRoomParamsWithContext(ctx context.Context) *AddStreamToTheRoomParams {
	return &AddStreamToTheRoomParams{
		Context: ctx,
	}
}

// NewAddStreamToTheRoomParamsWithHTTPClient creates a new AddStreamToTheRoomParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddStreamToTheRoomParamsWithHTTPClient(client *http.Client) *AddStreamToTheRoomParams {
	return &AddStreamToTheRoomParams{
		HTTPClient: client,
	}
}

/* AddStreamToTheRoomParams contains all the parameters to send to the API endpoint
   for the add stream to the room operation.

   Typically these are written to a http.Request.
*/
type AddStreamToTheRoomParams struct {

	/* RoomID.

	   Room id
	*/
	RoomID string

	/* StreamID.

	   Stream id to add to the conference room
	*/
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add stream to the room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddStreamToTheRoomParams) WithDefaults() *AddStreamToTheRoomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add stream to the room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddStreamToTheRoomParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add stream to the room params
func (o *AddStreamToTheRoomParams) WithTimeout(timeout time.Duration) *AddStreamToTheRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add stream to the room params
func (o *AddStreamToTheRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add stream to the room params
func (o *AddStreamToTheRoomParams) WithContext(ctx context.Context) *AddStreamToTheRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add stream to the room params
func (o *AddStreamToTheRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add stream to the room params
func (o *AddStreamToTheRoomParams) WithHTTPClient(client *http.Client) *AddStreamToTheRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add stream to the room params
func (o *AddStreamToTheRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoomID adds the roomID to the add stream to the room params
func (o *AddStreamToTheRoomParams) WithRoomID(roomID string) *AddStreamToTheRoomParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the add stream to the room params
func (o *AddStreamToTheRoomParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WithStreamID adds the streamID to the add stream to the room params
func (o *AddStreamToTheRoomParams) WithStreamID(streamID string) *AddStreamToTheRoomParams {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the add stream to the room params
func (o *AddStreamToTheRoomParams) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *AddStreamToTheRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param room_id
	if err := r.SetPathParam("room_id", o.RoomID); err != nil {
		return err
	}

	// query param streamId
	qrStreamID := o.StreamID
	qStreamID := qrStreamID
	if qStreamID != "" {

		if err := r.SetQueryParam("streamId", qStreamID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}