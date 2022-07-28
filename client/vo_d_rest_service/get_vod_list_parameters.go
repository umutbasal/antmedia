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
	"github.com/go-openapi/swag"
)

// NewGetVodListParams creates a new GetVodListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVodListParams() *GetVodListParams {
	return &GetVodListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVodListParamsWithTimeout creates a new GetVodListParams object
// with the ability to set a timeout on a request.
func NewGetVodListParamsWithTimeout(timeout time.Duration) *GetVodListParams {
	return &GetVodListParams{
		timeout: timeout,
	}
}

// NewGetVodListParamsWithContext creates a new GetVodListParams object
// with the ability to set a context for a request.
func NewGetVodListParamsWithContext(ctx context.Context) *GetVodListParams {
	return &GetVodListParams{
		Context: ctx,
	}
}

// NewGetVodListParamsWithHTTPClient creates a new GetVodListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVodListParamsWithHTTPClient(client *http.Client) *GetVodListParams {
	return &GetVodListParams{
		HTTPClient: client,
	}
}

/* GetVodListParams contains all the parameters to send to the API endpoint
   for the get vod list operation.

   Typically these are written to a http.Request.
*/
type GetVodListParams struct {

	/* Offset.

	   Offset of the list

	   Format: int32
	*/
	Offset int32

	/* OrderBy.

	   "asc" for Ascending, "desc" Descening order
	*/
	OrderBy *string

	/* Search.

	   Search string
	*/
	Search *string

	/* Size.

	   Number of items that will be fetched

	   Format: int32
	*/
	Size int32

	/* SortBy.

	   Field to sort. Possible values are "name", "date"
	*/
	SortBy *string

	/* StreamID.

	   Id of the stream to filter the results by stream id
	*/
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vod list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVodListParams) WithDefaults() *GetVodListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vod list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVodListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vod list params
func (o *GetVodListParams) WithTimeout(timeout time.Duration) *GetVodListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vod list params
func (o *GetVodListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vod list params
func (o *GetVodListParams) WithContext(ctx context.Context) *GetVodListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vod list params
func (o *GetVodListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vod list params
func (o *GetVodListParams) WithHTTPClient(client *http.Client) *GetVodListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vod list params
func (o *GetVodListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOffset adds the offset to the get vod list params
func (o *GetVodListParams) WithOffset(offset int32) *GetVodListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get vod list params
func (o *GetVodListParams) SetOffset(offset int32) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the get vod list params
func (o *GetVodListParams) WithOrderBy(orderBy *string) *GetVodListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get vod list params
func (o *GetVodListParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithSearch adds the search to the get vod list params
func (o *GetVodListParams) WithSearch(search *string) *GetVodListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get vod list params
func (o *GetVodListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSize adds the size to the get vod list params
func (o *GetVodListParams) WithSize(size int32) *GetVodListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get vod list params
func (o *GetVodListParams) SetSize(size int32) {
	o.Size = size
}

// WithSortBy adds the sortBy to the get vod list params
func (o *GetVodListParams) WithSortBy(sortBy *string) *GetVodListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get vod list params
func (o *GetVodListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithStreamID adds the streamID to the get vod list params
func (o *GetVodListParams) WithStreamID(streamID string) *GetVodListParams {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the get vod list params
func (o *GetVodListParams) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVodListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param offset
	if err := r.SetPathParam("offset", swag.FormatInt32(o.Offset)); err != nil {
		return err
	}

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	// path param size
	if err := r.SetPathParam("size", swag.FormatInt32(o.Size)); err != nil {
		return err
	}

	if o.SortBy != nil {

		// query param sort_by
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sort_by", qSortBy); err != nil {
				return err
			}
		}
	}

	// query param streamId
	qrStreamID := o.StreamID
	qStreamID := qrStreamID
	if qStreamID != "" {

		if err := r.SetQueryParam("streamId", qStreamID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}