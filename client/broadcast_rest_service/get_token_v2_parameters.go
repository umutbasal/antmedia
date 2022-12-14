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

// NewGetTokenV2Params creates a new GetTokenV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTokenV2Params() *GetTokenV2Params {
	return &GetTokenV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTokenV2ParamsWithTimeout creates a new GetTokenV2Params object
// with the ability to set a timeout on a request.
func NewGetTokenV2ParamsWithTimeout(timeout time.Duration) *GetTokenV2Params {
	return &GetTokenV2Params{
		timeout: timeout,
	}
}

// NewGetTokenV2ParamsWithContext creates a new GetTokenV2Params object
// with the ability to set a context for a request.
func NewGetTokenV2ParamsWithContext(ctx context.Context) *GetTokenV2Params {
	return &GetTokenV2Params{
		Context: ctx,
	}
}

// NewGetTokenV2ParamsWithHTTPClient creates a new GetTokenV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetTokenV2ParamsWithHTTPClient(client *http.Client) *GetTokenV2Params {
	return &GetTokenV2Params{
		HTTPClient: client,
	}
}

/* GetTokenV2Params contains all the parameters to send to the API endpoint
   for the get token v2 operation.

   Typically these are written to a http.Request.
*/
type GetTokenV2Params struct {

	/* ExpireDate.

	   The expire time of the token. It's in unix timestamp seconds

	   Format: int64
	*/
	ExpireDate int64

	/* ID.

	   The id of the stream
	*/
	ID string

	/* RoomID.

	   Room Id that token belongs to. It's not mandatory
	*/
	RoomID *string

	/* Type.

	   Type of the token. It may be play or publish
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get token v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTokenV2Params) WithDefaults() *GetTokenV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get token v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTokenV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get token v2 params
func (o *GetTokenV2Params) WithTimeout(timeout time.Duration) *GetTokenV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get token v2 params
func (o *GetTokenV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get token v2 params
func (o *GetTokenV2Params) WithContext(ctx context.Context) *GetTokenV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get token v2 params
func (o *GetTokenV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get token v2 params
func (o *GetTokenV2Params) WithHTTPClient(client *http.Client) *GetTokenV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get token v2 params
func (o *GetTokenV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpireDate adds the expireDate to the get token v2 params
func (o *GetTokenV2Params) WithExpireDate(expireDate int64) *GetTokenV2Params {
	o.SetExpireDate(expireDate)
	return o
}

// SetExpireDate adds the expireDate to the get token v2 params
func (o *GetTokenV2Params) SetExpireDate(expireDate int64) {
	o.ExpireDate = expireDate
}

// WithID adds the id to the get token v2 params
func (o *GetTokenV2Params) WithID(id string) *GetTokenV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get token v2 params
func (o *GetTokenV2Params) SetID(id string) {
	o.ID = id
}

// WithRoomID adds the roomID to the get token v2 params
func (o *GetTokenV2Params) WithRoomID(roomID *string) *GetTokenV2Params {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the get token v2 params
func (o *GetTokenV2Params) SetRoomID(roomID *string) {
	o.RoomID = roomID
}

// WithType adds the typeVar to the get token v2 params
func (o *GetTokenV2Params) WithType(typeVar string) *GetTokenV2Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get token v2 params
func (o *GetTokenV2Params) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetTokenV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param expireDate
	qrExpireDate := o.ExpireDate
	qExpireDate := swag.FormatInt64(qrExpireDate)
	if qExpireDate != "" {

		if err := r.SetQueryParam("expireDate", qExpireDate); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.RoomID != nil {

		// query param roomId
		var qrRoomID string

		if o.RoomID != nil {
			qrRoomID = *o.RoomID
		}
		qRoomID := qrRoomID
		if qRoomID != "" {

			if err := r.SetQueryParam("roomId", qRoomID); err != nil {
				return err
			}
		}
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {

		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
