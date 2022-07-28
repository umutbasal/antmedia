// Code generated by go-swagger; DO NOT EDIT.

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

// NewGetTotalVodNumberParams creates a new GetTotalVodNumberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTotalVodNumberParams() *GetTotalVodNumberParams {
	return &GetTotalVodNumberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTotalVodNumberParamsWithTimeout creates a new GetTotalVodNumberParams object
// with the ability to set a timeout on a request.
func NewGetTotalVodNumberParamsWithTimeout(timeout time.Duration) *GetTotalVodNumberParams {
	return &GetTotalVodNumberParams{
		timeout: timeout,
	}
}

// NewGetTotalVodNumberParamsWithContext creates a new GetTotalVodNumberParams object
// with the ability to set a context for a request.
func NewGetTotalVodNumberParamsWithContext(ctx context.Context) *GetTotalVodNumberParams {
	return &GetTotalVodNumberParams{
		Context: ctx,
	}
}

// NewGetTotalVodNumberParamsWithHTTPClient creates a new GetTotalVodNumberParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTotalVodNumberParamsWithHTTPClient(client *http.Client) *GetTotalVodNumberParams {
	return &GetTotalVodNumberParams{
		HTTPClient: client,
	}
}

/* GetTotalVodNumberParams contains all the parameters to send to the API endpoint
   for the get total vod number operation.

   Typically these are written to a http.Request.
*/
type GetTotalVodNumberParams struct {

	/* Search.

	   Search parameter to get the number of items including it
	*/
	Search string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get total vod number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTotalVodNumberParams) WithDefaults() *GetTotalVodNumberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get total vod number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTotalVodNumberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get total vod number params
func (o *GetTotalVodNumberParams) WithTimeout(timeout time.Duration) *GetTotalVodNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get total vod number params
func (o *GetTotalVodNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get total vod number params
func (o *GetTotalVodNumberParams) WithContext(ctx context.Context) *GetTotalVodNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get total vod number params
func (o *GetTotalVodNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get total vod number params
func (o *GetTotalVodNumberParams) WithHTTPClient(client *http.Client) *GetTotalVodNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get total vod number params
func (o *GetTotalVodNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearch adds the search to the get total vod number params
func (o *GetTotalVodNumberParams) WithSearch(search string) *GetTotalVodNumberParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get total vod number params
func (o *GetTotalVodNumberParams) SetSearch(search string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetTotalVodNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param search
	if err := r.SetPathParam("search", o.Search); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}