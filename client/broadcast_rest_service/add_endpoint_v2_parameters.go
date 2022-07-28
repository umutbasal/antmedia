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

// NewAddEndpointV2Params creates a new AddEndpointV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddEndpointV2Params() *AddEndpointV2Params {
	return &AddEndpointV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddEndpointV2ParamsWithTimeout creates a new AddEndpointV2Params object
// with the ability to set a timeout on a request.
func NewAddEndpointV2ParamsWithTimeout(timeout time.Duration) *AddEndpointV2Params {
	return &AddEndpointV2Params{
		timeout: timeout,
	}
}

// NewAddEndpointV2ParamsWithContext creates a new AddEndpointV2Params object
// with the ability to set a context for a request.
func NewAddEndpointV2ParamsWithContext(ctx context.Context) *AddEndpointV2Params {
	return &AddEndpointV2Params{
		Context: ctx,
	}
}

// NewAddEndpointV2ParamsWithHTTPClient creates a new AddEndpointV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewAddEndpointV2ParamsWithHTTPClient(client *http.Client) *AddEndpointV2Params {
	return &AddEndpointV2Params{
		HTTPClient: client,
	}
}

/* AddEndpointV2Params contains all the parameters to send to the API endpoint
   for the add endpoint v2 operation.

   Typically these are written to a http.Request.
*/
type AddEndpointV2Params struct {

	/* ID.

	   Broadcast id
	*/
	ID string

	/* RtmpURL.

	   RTMP url of the endpoint that stream will be republished. If required, please encode the URL
	*/
	RtmpURL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add endpoint v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddEndpointV2Params) WithDefaults() *AddEndpointV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add endpoint v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddEndpointV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add endpoint v2 params
func (o *AddEndpointV2Params) WithTimeout(timeout time.Duration) *AddEndpointV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add endpoint v2 params
func (o *AddEndpointV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add endpoint v2 params
func (o *AddEndpointV2Params) WithContext(ctx context.Context) *AddEndpointV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add endpoint v2 params
func (o *AddEndpointV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add endpoint v2 params
func (o *AddEndpointV2Params) WithHTTPClient(client *http.Client) *AddEndpointV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add endpoint v2 params
func (o *AddEndpointV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the add endpoint v2 params
func (o *AddEndpointV2Params) WithID(id string) *AddEndpointV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the add endpoint v2 params
func (o *AddEndpointV2Params) SetID(id string) {
	o.ID = id
}

// WithRtmpURL adds the rtmpURL to the add endpoint v2 params
func (o *AddEndpointV2Params) WithRtmpURL(rtmpURL string) *AddEndpointV2Params {
	o.SetRtmpURL(rtmpURL)
	return o
}

// SetRtmpURL adds the rtmpUrl to the add endpoint v2 params
func (o *AddEndpointV2Params) SetRtmpURL(rtmpURL string) {
	o.RtmpURL = rtmpURL
}

// WriteToRequest writes these params to a swagger request
func (o *AddEndpointV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param rtmpUrl
	qrRtmpURL := o.RtmpURL
	qRtmpURL := qrRtmpURL
	if qRtmpURL != "" {

		if err := r.SetQueryParam("rtmpUrl", qRtmpURL); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
