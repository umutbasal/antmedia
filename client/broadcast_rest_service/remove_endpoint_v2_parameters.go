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

// NewRemoveEndpointV2Params creates a new RemoveEndpointV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveEndpointV2Params() *RemoveEndpointV2Params {
	return &RemoveEndpointV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveEndpointV2ParamsWithTimeout creates a new RemoveEndpointV2Params object
// with the ability to set a timeout on a request.
func NewRemoveEndpointV2ParamsWithTimeout(timeout time.Duration) *RemoveEndpointV2Params {
	return &RemoveEndpointV2Params{
		timeout: timeout,
	}
}

// NewRemoveEndpointV2ParamsWithContext creates a new RemoveEndpointV2Params object
// with the ability to set a context for a request.
func NewRemoveEndpointV2ParamsWithContext(ctx context.Context) *RemoveEndpointV2Params {
	return &RemoveEndpointV2Params{
		Context: ctx,
	}
}

// NewRemoveEndpointV2ParamsWithHTTPClient creates a new RemoveEndpointV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveEndpointV2ParamsWithHTTPClient(client *http.Client) *RemoveEndpointV2Params {
	return &RemoveEndpointV2Params{
		HTTPClient: client,
	}
}

/* RemoveEndpointV2Params contains all the parameters to send to the API endpoint
   for the remove endpoint v2 operation.

   Typically these are written to a http.Request.
*/
type RemoveEndpointV2Params struct {

	/* EndpointServiceID.

	   RTMP url of the endpoint that will be stopped.
	*/
	EndpointServiceID string

	/* ID.

	   Broadcast id
	*/
	ID string

	/* ResolutionHeight.

	   Resolution specifier if endpoint has been added with resolution. Only applicable if user added RTMP endpoint with a resolution speficier. Otherwise won't work and won't remove the endpoint.

	   Format: int32
	*/
	ResolutionHeight int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove endpoint v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveEndpointV2Params) WithDefaults() *RemoveEndpointV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove endpoint v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveEndpointV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) WithTimeout(timeout time.Duration) *RemoveEndpointV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) WithContext(ctx context.Context) *RemoveEndpointV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) WithHTTPClient(client *http.Client) *RemoveEndpointV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointServiceID adds the endpointServiceID to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) WithEndpointServiceID(endpointServiceID string) *RemoveEndpointV2Params {
	o.SetEndpointServiceID(endpointServiceID)
	return o
}

// SetEndpointServiceID adds the endpointServiceId to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) SetEndpointServiceID(endpointServiceID string) {
	o.EndpointServiceID = endpointServiceID
}

// WithID adds the id to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) WithID(id string) *RemoveEndpointV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) SetID(id string) {
	o.ID = id
}

// WithResolutionHeight adds the resolutionHeight to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) WithResolutionHeight(resolutionHeight int32) *RemoveEndpointV2Params {
	o.SetResolutionHeight(resolutionHeight)
	return o
}

// SetResolutionHeight adds the resolutionHeight to the remove endpoint v2 params
func (o *RemoveEndpointV2Params) SetResolutionHeight(resolutionHeight int32) {
	o.ResolutionHeight = resolutionHeight
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveEndpointV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endpointServiceId
	qrEndpointServiceID := o.EndpointServiceID
	qEndpointServiceID := qrEndpointServiceID
	if qEndpointServiceID != "" {

		if err := r.SetQueryParam("endpointServiceId", qEndpointServiceID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param resolutionHeight
	qrResolutionHeight := o.ResolutionHeight
	qResolutionHeight := swag.FormatInt32(qrResolutionHeight)
	if qResolutionHeight != "" {

		if err := r.SetQueryParam("resolutionHeight", qResolutionHeight); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
