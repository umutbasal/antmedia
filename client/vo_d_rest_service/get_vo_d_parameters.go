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

// NewGetVoDParams creates a new GetVoDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVoDParams() *GetVoDParams {
	return &GetVoDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoDParamsWithTimeout creates a new GetVoDParams object
// with the ability to set a timeout on a request.
func NewGetVoDParamsWithTimeout(timeout time.Duration) *GetVoDParams {
	return &GetVoDParams{
		timeout: timeout,
	}
}

// NewGetVoDParamsWithContext creates a new GetVoDParams object
// with the ability to set a context for a request.
func NewGetVoDParamsWithContext(ctx context.Context) *GetVoDParams {
	return &GetVoDParams{
		Context: ctx,
	}
}

// NewGetVoDParamsWithHTTPClient creates a new GetVoDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVoDParamsWithHTTPClient(client *http.Client) *GetVoDParams {
	return &GetVoDParams{
		HTTPClient: client,
	}
}

/* GetVoDParams contains all the parameters to send to the API endpoint
   for the get vo d operation.

   Typically these are written to a http.Request.
*/
type GetVoDParams struct {

	/* ID.

	   id of the VoD
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vo d params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVoDParams) WithDefaults() *GetVoDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vo d params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVoDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vo d params
func (o *GetVoDParams) WithTimeout(timeout time.Duration) *GetVoDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vo d params
func (o *GetVoDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vo d params
func (o *GetVoDParams) WithContext(ctx context.Context) *GetVoDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vo d params
func (o *GetVoDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vo d params
func (o *GetVoDParams) WithHTTPClient(client *http.Client) *GetVoDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vo d params
func (o *GetVoDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vo d params
func (o *GetVoDParams) WithID(id string) *GetVoDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vo d params
func (o *GetVoDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
