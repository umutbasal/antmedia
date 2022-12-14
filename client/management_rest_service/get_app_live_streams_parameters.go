// Code generated by go-swagger;

package management_rest_service

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

// NewGetAppLiveStreamsParams creates a new GetAppLiveStreamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppLiveStreamsParams() *GetAppLiveStreamsParams {
	return &GetAppLiveStreamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppLiveStreamsParamsWithTimeout creates a new GetAppLiveStreamsParams object
// with the ability to set a timeout on a request.
func NewGetAppLiveStreamsParamsWithTimeout(timeout time.Duration) *GetAppLiveStreamsParams {
	return &GetAppLiveStreamsParams{
		timeout: timeout,
	}
}

// NewGetAppLiveStreamsParamsWithContext creates a new GetAppLiveStreamsParams object
// with the ability to set a context for a request.
func NewGetAppLiveStreamsParamsWithContext(ctx context.Context) *GetAppLiveStreamsParams {
	return &GetAppLiveStreamsParams{
		Context: ctx,
	}
}

// NewGetAppLiveStreamsParamsWithHTTPClient creates a new GetAppLiveStreamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppLiveStreamsParamsWithHTTPClient(client *http.Client) *GetAppLiveStreamsParams {
	return &GetAppLiveStreamsParams{
		HTTPClient: client,
	}
}

/* GetAppLiveStreamsParams contains all the parameters to send to the API endpoint
   for the get app live streams operation.

   Typically these are written to a http.Request.
*/
type GetAppLiveStreamsParams struct {

	/* Appname.

	   Application name
	*/
	Appname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app live streams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppLiveStreamsParams) WithDefaults() *GetAppLiveStreamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app live streams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppLiveStreamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app live streams params
func (o *GetAppLiveStreamsParams) WithTimeout(timeout time.Duration) *GetAppLiveStreamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app live streams params
func (o *GetAppLiveStreamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app live streams params
func (o *GetAppLiveStreamsParams) WithContext(ctx context.Context) *GetAppLiveStreamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app live streams params
func (o *GetAppLiveStreamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app live streams params
func (o *GetAppLiveStreamsParams) WithHTTPClient(client *http.Client) *GetAppLiveStreamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app live streams params
func (o *GetAppLiveStreamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppname adds the appname to the get app live streams params
func (o *GetAppLiveStreamsParams) WithAppname(appname string) *GetAppLiveStreamsParams {
	o.SetAppname(appname)
	return o
}

// SetAppname adds the appname to the get app live streams params
func (o *GetAppLiveStreamsParams) SetAppname(appname string) {
	o.Appname = appname
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppLiveStreamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appname
	if err := r.SetPathParam("appname", o.Appname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
