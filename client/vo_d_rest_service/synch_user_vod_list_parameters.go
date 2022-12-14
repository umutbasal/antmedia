// Code generated by go-swagger;

package vo_d_rest_service

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

// NewSynchUserVodListParams creates a new SynchUserVodListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSynchUserVodListParams() *SynchUserVodListParams {
	return &SynchUserVodListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSynchUserVodListParamsWithTimeout creates a new SynchUserVodListParams object
// with the ability to set a timeout on a request.
func NewSynchUserVodListParamsWithTimeout(timeout time.Duration) *SynchUserVodListParams {
	return &SynchUserVodListParams{
		timeout: timeout,
	}
}

// NewSynchUserVodListParamsWithContext creates a new SynchUserVodListParams object
// with the ability to set a context for a request.
func NewSynchUserVodListParamsWithContext(ctx context.Context) *SynchUserVodListParams {
	return &SynchUserVodListParams{
		Context: ctx,
	}
}

// NewSynchUserVodListParamsWithHTTPClient creates a new SynchUserVodListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSynchUserVodListParamsWithHTTPClient(client *http.Client) *SynchUserVodListParams {
	return &SynchUserVodListParams{
		HTTPClient: client,
	}
}

/* SynchUserVodListParams contains all the parameters to send to the API endpoint
   for the synch user vod list operation.

   Typically these are written to a http.Request.
*/
type SynchUserVodListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the synch user vod list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SynchUserVodListParams) WithDefaults() *SynchUserVodListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the synch user vod list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SynchUserVodListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the synch user vod list params
func (o *SynchUserVodListParams) WithTimeout(timeout time.Duration) *SynchUserVodListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the synch user vod list params
func (o *SynchUserVodListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the synch user vod list params
func (o *SynchUserVodListParams) WithContext(ctx context.Context) *SynchUserVodListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the synch user vod list params
func (o *SynchUserVodListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the synch user vod list params
func (o *SynchUserVodListParams) WithHTTPClient(client *http.Client) *SynchUserVodListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the synch user vod list params
func (o *SynchUserVodListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SynchUserVodListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
