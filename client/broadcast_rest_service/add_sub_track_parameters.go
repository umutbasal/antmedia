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

// NewAddSubTrackParams creates a new AddSubTrackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddSubTrackParams() *AddSubTrackParams {
	return &AddSubTrackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddSubTrackParamsWithTimeout creates a new AddSubTrackParams object
// with the ability to set a timeout on a request.
func NewAddSubTrackParamsWithTimeout(timeout time.Duration) *AddSubTrackParams {
	return &AddSubTrackParams{
		timeout: timeout,
	}
}

// NewAddSubTrackParamsWithContext creates a new AddSubTrackParams object
// with the ability to set a context for a request.
func NewAddSubTrackParamsWithContext(ctx context.Context) *AddSubTrackParams {
	return &AddSubTrackParams{
		Context: ctx,
	}
}

// NewAddSubTrackParamsWithHTTPClient creates a new AddSubTrackParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddSubTrackParamsWithHTTPClient(client *http.Client) *AddSubTrackParams {
	return &AddSubTrackParams{
		HTTPClient: client,
	}
}

/* AddSubTrackParams contains all the parameters to send to the API endpoint
   for the add sub track operation.

   Typically these are written to a http.Request.
*/
type AddSubTrackParams struct {

	/* ID.

	   Broadcast id
	*/
	PathID string

	/* ID.

	   Subtrack Stream Id
	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add sub track params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSubTrackParams) WithDefaults() *AddSubTrackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add sub track params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSubTrackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add sub track params
func (o *AddSubTrackParams) WithTimeout(timeout time.Duration) *AddSubTrackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add sub track params
func (o *AddSubTrackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add sub track params
func (o *AddSubTrackParams) WithContext(ctx context.Context) *AddSubTrackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add sub track params
func (o *AddSubTrackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add sub track params
func (o *AddSubTrackParams) WithHTTPClient(client *http.Client) *AddSubTrackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add sub track params
func (o *AddSubTrackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPathID adds the id to the add sub track params
func (o *AddSubTrackParams) WithPathID(id string) *AddSubTrackParams {
	o.SetPathID(id)
	return o
}

// SetPathID adds the id to the add sub track params
func (o *AddSubTrackParams) SetPathID(id string) {
	o.PathID = id
}

// WithQueryID adds the id to the add sub track params
func (o *AddSubTrackParams) WithQueryID(id string) *AddSubTrackParams {
	o.SetQueryID(id)
	return o
}

// SetQueryID adds the id to the add sub track params
func (o *AddSubTrackParams) SetQueryID(id string) {
	o.QueryID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddSubTrackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.PathID); err != nil {
		return err
	}

	// query param id
	qrID := o.QueryID
	qID := qrID
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
