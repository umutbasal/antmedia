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
)

// NewChangeLogSettingsParams creates a new ChangeLogSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeLogSettingsParams() *ChangeLogSettingsParams {
	return &ChangeLogSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeLogSettingsParamsWithTimeout creates a new ChangeLogSettingsParams object
// with the ability to set a timeout on a request.
func NewChangeLogSettingsParamsWithTimeout(timeout time.Duration) *ChangeLogSettingsParams {
	return &ChangeLogSettingsParams{
		timeout: timeout,
	}
}

// NewChangeLogSettingsParamsWithContext creates a new ChangeLogSettingsParams object
// with the ability to set a context for a request.
func NewChangeLogSettingsParamsWithContext(ctx context.Context) *ChangeLogSettingsParams {
	return &ChangeLogSettingsParams{
		Context: ctx,
	}
}

// NewChangeLogSettingsParamsWithHTTPClient creates a new ChangeLogSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeLogSettingsParamsWithHTTPClient(client *http.Client) *ChangeLogSettingsParams {
	return &ChangeLogSettingsParams{
		HTTPClient: client,
	}
}

/* ChangeLogSettingsParams contains all the parameters to send to the API endpoint
   for the change log settings operation.

   Typically these are written to a http.Request.
*/
type ChangeLogSettingsParams struct {

	// Level.
	Level string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change log settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeLogSettingsParams) WithDefaults() *ChangeLogSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change log settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeLogSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change log settings params
func (o *ChangeLogSettingsParams) WithTimeout(timeout time.Duration) *ChangeLogSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change log settings params
func (o *ChangeLogSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change log settings params
func (o *ChangeLogSettingsParams) WithContext(ctx context.Context) *ChangeLogSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change log settings params
func (o *ChangeLogSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change log settings params
func (o *ChangeLogSettingsParams) WithHTTPClient(client *http.Client) *ChangeLogSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change log settings params
func (o *ChangeLogSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLevel adds the level to the change log settings params
func (o *ChangeLogSettingsParams) WithLevel(level string) *ChangeLogSettingsParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the change log settings params
func (o *ChangeLogSettingsParams) SetLevel(level string) {
	o.Level = level
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeLogSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param level
	if err := r.SetPathParam("level", o.Level); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
