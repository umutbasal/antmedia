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

// NewGetWebRTCClientStatsListV2Params creates a new GetWebRTCClientStatsListV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebRTCClientStatsListV2Params() *GetWebRTCClientStatsListV2Params {
	return &GetWebRTCClientStatsListV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebRTCClientStatsListV2ParamsWithTimeout creates a new GetWebRTCClientStatsListV2Params object
// with the ability to set a timeout on a request.
func NewGetWebRTCClientStatsListV2ParamsWithTimeout(timeout time.Duration) *GetWebRTCClientStatsListV2Params {
	return &GetWebRTCClientStatsListV2Params{
		timeout: timeout,
	}
}

// NewGetWebRTCClientStatsListV2ParamsWithContext creates a new GetWebRTCClientStatsListV2Params object
// with the ability to set a context for a request.
func NewGetWebRTCClientStatsListV2ParamsWithContext(ctx context.Context) *GetWebRTCClientStatsListV2Params {
	return &GetWebRTCClientStatsListV2Params{
		Context: ctx,
	}
}

// NewGetWebRTCClientStatsListV2ParamsWithHTTPClient creates a new GetWebRTCClientStatsListV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebRTCClientStatsListV2ParamsWithHTTPClient(client *http.Client) *GetWebRTCClientStatsListV2Params {
	return &GetWebRTCClientStatsListV2Params{
		HTTPClient: client,
	}
}

/* GetWebRTCClientStatsListV2Params contains all the parameters to send to the API endpoint
   for the get web r t c client stats list v2 operation.

   Typically these are written to a http.Request.
*/
type GetWebRTCClientStatsListV2Params struct {

	/* Offset.

	   offset of the list

	   Format: int32
	*/
	Offset int32

	/* Size.

	   Number of items that will be fetched

	   Format: int32
	*/
	Size int32

	/* StreamID.

	   the id of the stream
	*/
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get web r t c client stats list v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebRTCClientStatsListV2Params) WithDefaults() *GetWebRTCClientStatsListV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get web r t c client stats list v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebRTCClientStatsListV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) WithTimeout(timeout time.Duration) *GetWebRTCClientStatsListV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) WithContext(ctx context.Context) *GetWebRTCClientStatsListV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) WithHTTPClient(client *http.Client) *GetWebRTCClientStatsListV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOffset adds the offset to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) WithOffset(offset int32) *GetWebRTCClientStatsListV2Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) SetOffset(offset int32) {
	o.Offset = offset
}

// WithSize adds the size to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) WithSize(size int32) *GetWebRTCClientStatsListV2Params {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) SetSize(size int32) {
	o.Size = size
}

// WithStreamID adds the streamID to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) WithStreamID(streamID string) *GetWebRTCClientStatsListV2Params {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the get web r t c client stats list v2 params
func (o *GetWebRTCClientStatsListV2Params) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebRTCClientStatsListV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param offset
	if err := r.SetPathParam("offset", swag.FormatInt32(o.Offset)); err != nil {
		return err
	}

	// path param size
	if err := r.SetPathParam("size", swag.FormatInt32(o.Size)); err != nil {
		return err
	}

	// path param stream_id
	if err := r.SetPathParam("stream_id", o.StreamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
