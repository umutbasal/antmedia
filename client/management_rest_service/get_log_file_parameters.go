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
	"github.com/go-openapi/swag"
)

// NewGetLogFileParams creates a new GetLogFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogFileParams() *GetLogFileParams {
	return &GetLogFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogFileParamsWithTimeout creates a new GetLogFileParams object
// with the ability to set a timeout on a request.
func NewGetLogFileParamsWithTimeout(timeout time.Duration) *GetLogFileParams {
	return &GetLogFileParams{
		timeout: timeout,
	}
}

// NewGetLogFileParamsWithContext creates a new GetLogFileParams object
// with the ability to set a context for a request.
func NewGetLogFileParamsWithContext(ctx context.Context) *GetLogFileParams {
	return &GetLogFileParams{
		Context: ctx,
	}
}

// NewGetLogFileParamsWithHTTPClient creates a new GetLogFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogFileParamsWithHTTPClient(client *http.Client) *GetLogFileParams {
	return &GetLogFileParams{
		HTTPClient: client,
	}
}

/* GetLogFileParams contains all the parameters to send to the API endpoint
   for the get log file operation.

   Typically these are written to a http.Request.
*/
type GetLogFileParams struct {

	/* CharSize.

	   Char size of the log

	   Format: int32
	*/
	CharSize int32

	/* LogType.

	   Log type. ERROR can be used to get only error logs
	*/
	LogType string

	/* OffsetSize.

	   Offset of the retrieved log

	   Format: int64
	*/
	OffsetSize int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogFileParams) WithDefaults() *GetLogFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log file params
func (o *GetLogFileParams) WithTimeout(timeout time.Duration) *GetLogFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log file params
func (o *GetLogFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log file params
func (o *GetLogFileParams) WithContext(ctx context.Context) *GetLogFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log file params
func (o *GetLogFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log file params
func (o *GetLogFileParams) WithHTTPClient(client *http.Client) *GetLogFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log file params
func (o *GetLogFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharSize adds the charSize to the get log file params
func (o *GetLogFileParams) WithCharSize(charSize int32) *GetLogFileParams {
	o.SetCharSize(charSize)
	return o
}

// SetCharSize adds the charSize to the get log file params
func (o *GetLogFileParams) SetCharSize(charSize int32) {
	o.CharSize = charSize
}

// WithLogType adds the logType to the get log file params
func (o *GetLogFileParams) WithLogType(logType string) *GetLogFileParams {
	o.SetLogType(logType)
	return o
}

// SetLogType adds the logType to the get log file params
func (o *GetLogFileParams) SetLogType(logType string) {
	o.LogType = logType
}

// WithOffsetSize adds the offsetSize to the get log file params
func (o *GetLogFileParams) WithOffsetSize(offsetSize int64) *GetLogFileParams {
	o.SetOffsetSize(offsetSize)
	return o
}

// SetOffsetSize adds the offsetSize to the get log file params
func (o *GetLogFileParams) SetOffsetSize(offsetSize int64) {
	o.OffsetSize = offsetSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param charSize
	if err := r.SetPathParam("charSize", swag.FormatInt32(o.CharSize)); err != nil {
		return err
	}

	// query param logType
	qrLogType := o.LogType
	qLogType := qrLogType
	if qLogType != "" {

		if err := r.SetQueryParam("logType", qLogType); err != nil {
			return err
		}
	}

	// path param offsetSize
	if err := r.SetPathParam("offsetSize", swag.FormatInt64(o.OffsetSize)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
