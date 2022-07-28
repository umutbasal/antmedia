// Code generated by go-swagger; DO NOT EDIT.

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

// NewSearchOnvifDevicesV2Params creates a new SearchOnvifDevicesV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchOnvifDevicesV2Params() *SearchOnvifDevicesV2Params {
	return &SearchOnvifDevicesV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchOnvifDevicesV2ParamsWithTimeout creates a new SearchOnvifDevicesV2Params object
// with the ability to set a timeout on a request.
func NewSearchOnvifDevicesV2ParamsWithTimeout(timeout time.Duration) *SearchOnvifDevicesV2Params {
	return &SearchOnvifDevicesV2Params{
		timeout: timeout,
	}
}

// NewSearchOnvifDevicesV2ParamsWithContext creates a new SearchOnvifDevicesV2Params object
// with the ability to set a context for a request.
func NewSearchOnvifDevicesV2ParamsWithContext(ctx context.Context) *SearchOnvifDevicesV2Params {
	return &SearchOnvifDevicesV2Params{
		Context: ctx,
	}
}

// NewSearchOnvifDevicesV2ParamsWithHTTPClient creates a new SearchOnvifDevicesV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewSearchOnvifDevicesV2ParamsWithHTTPClient(client *http.Client) *SearchOnvifDevicesV2Params {
	return &SearchOnvifDevicesV2Params{
		HTTPClient: client,
	}
}

/* SearchOnvifDevicesV2Params contains all the parameters to send to the API endpoint
   for the search onvif devices v2 operation.

   Typically these are written to a http.Request.
*/
type SearchOnvifDevicesV2Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search onvif devices v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchOnvifDevicesV2Params) WithDefaults() *SearchOnvifDevicesV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search onvif devices v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchOnvifDevicesV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search onvif devices v2 params
func (o *SearchOnvifDevicesV2Params) WithTimeout(timeout time.Duration) *SearchOnvifDevicesV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search onvif devices v2 params
func (o *SearchOnvifDevicesV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search onvif devices v2 params
func (o *SearchOnvifDevicesV2Params) WithContext(ctx context.Context) *SearchOnvifDevicesV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search onvif devices v2 params
func (o *SearchOnvifDevicesV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search onvif devices v2 params
func (o *SearchOnvifDevicesV2Params) WithHTTPClient(client *http.Client) *SearchOnvifDevicesV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search onvif devices v2 params
func (o *SearchOnvifDevicesV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SearchOnvifDevicesV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}