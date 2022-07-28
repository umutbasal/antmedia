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

// NewGetBroadcastStatisticsParams creates a new GetBroadcastStatisticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBroadcastStatisticsParams() *GetBroadcastStatisticsParams {
	return &GetBroadcastStatisticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBroadcastStatisticsParamsWithTimeout creates a new GetBroadcastStatisticsParams object
// with the ability to set a timeout on a request.
func NewGetBroadcastStatisticsParamsWithTimeout(timeout time.Duration) *GetBroadcastStatisticsParams {
	return &GetBroadcastStatisticsParams{
		timeout: timeout,
	}
}

// NewGetBroadcastStatisticsParamsWithContext creates a new GetBroadcastStatisticsParams object
// with the ability to set a context for a request.
func NewGetBroadcastStatisticsParamsWithContext(ctx context.Context) *GetBroadcastStatisticsParams {
	return &GetBroadcastStatisticsParams{
		Context: ctx,
	}
}

// NewGetBroadcastStatisticsParamsWithHTTPClient creates a new GetBroadcastStatisticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBroadcastStatisticsParamsWithHTTPClient(client *http.Client) *GetBroadcastStatisticsParams {
	return &GetBroadcastStatisticsParams{
		HTTPClient: client,
	}
}

/* GetBroadcastStatisticsParams contains all the parameters to send to the API endpoint
   for the get broadcast statistics operation.

   Typically these are written to a http.Request.
*/
type GetBroadcastStatisticsParams struct {

	/* ID.

	   the id of the stream
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get broadcast statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBroadcastStatisticsParams) WithDefaults() *GetBroadcastStatisticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get broadcast statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBroadcastStatisticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get broadcast statistics params
func (o *GetBroadcastStatisticsParams) WithTimeout(timeout time.Duration) *GetBroadcastStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get broadcast statistics params
func (o *GetBroadcastStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get broadcast statistics params
func (o *GetBroadcastStatisticsParams) WithContext(ctx context.Context) *GetBroadcastStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get broadcast statistics params
func (o *GetBroadcastStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get broadcast statistics params
func (o *GetBroadcastStatisticsParams) WithHTTPClient(client *http.Client) *GetBroadcastStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get broadcast statistics params
func (o *GetBroadcastStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get broadcast statistics params
func (o *GetBroadcastStatisticsParams) WithID(id string) *GetBroadcastStatisticsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get broadcast statistics params
func (o *GetBroadcastStatisticsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBroadcastStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
