// Code generated by go-swagger;

package operations

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

// NewGetNodeListParams creates a new GetNodeListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNodeListParams() *GetNodeListParams {
	return &GetNodeListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeListParamsWithTimeout creates a new GetNodeListParams object
// with the ability to set a timeout on a request.
func NewGetNodeListParamsWithTimeout(timeout time.Duration) *GetNodeListParams {
	return &GetNodeListParams{
		timeout: timeout,
	}
}

// NewGetNodeListParamsWithContext creates a new GetNodeListParams object
// with the ability to set a context for a request.
func NewGetNodeListParamsWithContext(ctx context.Context) *GetNodeListParams {
	return &GetNodeListParams{
		Context: ctx,
	}
}

// NewGetNodeListParamsWithHTTPClient creates a new GetNodeListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNodeListParamsWithHTTPClient(client *http.Client) *GetNodeListParams {
	return &GetNodeListParams{
		HTTPClient: client,
	}
}

/* GetNodeListParams contains all the parameters to send to the API endpoint
   for the get node list operation.

   Typically these are written to a http.Request.
*/
type GetNodeListParams struct {

	// Offset.
	//
	// Format: int32
	Offset int32

	// Size.
	//
	// Format: int32
	Size int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get node list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodeListParams) WithDefaults() *GetNodeListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get node list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodeListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get node list params
func (o *GetNodeListParams) WithTimeout(timeout time.Duration) *GetNodeListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node list params
func (o *GetNodeListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node list params
func (o *GetNodeListParams) WithContext(ctx context.Context) *GetNodeListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node list params
func (o *GetNodeListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node list params
func (o *GetNodeListParams) WithHTTPClient(client *http.Client) *GetNodeListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node list params
func (o *GetNodeListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOffset adds the offset to the get node list params
func (o *GetNodeListParams) WithOffset(offset int32) *GetNodeListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get node list params
func (o *GetNodeListParams) SetOffset(offset int32) {
	o.Offset = offset
}

// WithSize adds the size to the get node list params
func (o *GetNodeListParams) WithSize(size int32) *GetNodeListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get node list params
func (o *GetNodeListParams) SetSize(size int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
