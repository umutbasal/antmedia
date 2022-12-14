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
	"github.com/go-openapi/swag"
)

// NewEnableRecordingParams creates a new EnableRecordingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableRecordingParams() *EnableRecordingParams {
	return &EnableRecordingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableRecordingParamsWithTimeout creates a new EnableRecordingParams object
// with the ability to set a timeout on a request.
func NewEnableRecordingParamsWithTimeout(timeout time.Duration) *EnableRecordingParams {
	return &EnableRecordingParams{
		timeout: timeout,
	}
}

// NewEnableRecordingParamsWithContext creates a new EnableRecordingParams object
// with the ability to set a context for a request.
func NewEnableRecordingParamsWithContext(ctx context.Context) *EnableRecordingParams {
	return &EnableRecordingParams{
		Context: ctx,
	}
}

// NewEnableRecordingParamsWithHTTPClient creates a new EnableRecordingParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableRecordingParamsWithHTTPClient(client *http.Client) *EnableRecordingParams {
	return &EnableRecordingParams{
		HTTPClient: client,
	}
}

/* EnableRecordingParams contains all the parameters to send to the API endpoint
   for the enable recording operation.

   Typically these are written to a http.Request.
*/
type EnableRecordingParams struct {

	/* ID.

	   the id of the stream
	*/
	ID string

	/* RecordType.

	   Record type: 'mp4' or 'webm'. It's optional parameter.
	*/
	RecordType *string

	/* RecordingStatus.

	   Change recording status. If true, starts recording. If false stop recording
	*/
	RecordingStatus bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable recording params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableRecordingParams) WithDefaults() *EnableRecordingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable recording params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableRecordingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable recording params
func (o *EnableRecordingParams) WithTimeout(timeout time.Duration) *EnableRecordingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable recording params
func (o *EnableRecordingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable recording params
func (o *EnableRecordingParams) WithContext(ctx context.Context) *EnableRecordingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable recording params
func (o *EnableRecordingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable recording params
func (o *EnableRecordingParams) WithHTTPClient(client *http.Client) *EnableRecordingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable recording params
func (o *EnableRecordingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the enable recording params
func (o *EnableRecordingParams) WithID(id string) *EnableRecordingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enable recording params
func (o *EnableRecordingParams) SetID(id string) {
	o.ID = id
}

// WithRecordType adds the recordType to the enable recording params
func (o *EnableRecordingParams) WithRecordType(recordType *string) *EnableRecordingParams {
	o.SetRecordType(recordType)
	return o
}

// SetRecordType adds the recordType to the enable recording params
func (o *EnableRecordingParams) SetRecordType(recordType *string) {
	o.RecordType = recordType
}

// WithRecordingStatus adds the recordingStatus to the enable recording params
func (o *EnableRecordingParams) WithRecordingStatus(recordingStatus bool) *EnableRecordingParams {
	o.SetRecordingStatus(recordingStatus)
	return o
}

// SetRecordingStatus adds the recordingStatus to the enable recording params
func (o *EnableRecordingParams) SetRecordingStatus(recordingStatus bool) {
	o.RecordingStatus = recordingStatus
}

// WriteToRequest writes these params to a swagger request
func (o *EnableRecordingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.RecordType != nil {

		// query param recordType
		var qrRecordType string

		if o.RecordType != nil {
			qrRecordType = *o.RecordType
		}
		qRecordType := qrRecordType
		if qRecordType != "" {

			if err := r.SetQueryParam("recordType", qRecordType); err != nil {
				return err
			}
		}
	}

	// path param recording-status
	if err := r.SetPathParam("recording-status", swag.FormatBool(o.RecordingStatus)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
