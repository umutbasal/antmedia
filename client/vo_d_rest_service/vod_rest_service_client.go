// Code generated by go-swagger;

package vo_d_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vo d rest service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vo d rest service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteVoD(params *DeleteVoDParams, opts ...ClientOption) (*DeleteVoDOK, error)

	DeleteVoDs(params *DeleteVoDsParams, opts ...ClientOption) (*DeleteVoDsOK, error)

	GetTotalVodNumber(params *GetTotalVodNumberParams, opts ...ClientOption) (*GetTotalVodNumberOK, error)

	GetVoD(params *GetVoDParams, opts ...ClientOption) (*GetVoDOK, error)

	GetVodList(params *GetVodListParams, opts ...ClientOption) (*GetVodListOK, error)

	ImportVoDsToStalker(params *ImportVoDsToStalkerParams, opts ...ClientOption) (*ImportVoDsToStalkerOK, error)

	SynchUserVodList(params *SynchUserVodListParams, opts ...ClientOption) (*SynchUserVodListOK, error)

	UploadVoDFile(params *UploadVoDFileParams, opts ...ClientOption) (*UploadVoDFileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteVoD deletes specific vo d file
*/
func (a *Client) DeleteVoD(params *DeleteVoDParams, opts ...ClientOption) (*DeleteVoDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVoDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVoD",
		Method:             "DELETE",
		PathPattern:        "/v2/vods/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVoDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVoDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVoD: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVoDs deletes bulk vo d files based on vod Id
*/
func (a *Client) DeleteVoDs(params *DeleteVoDsParams, opts ...ClientOption) (*DeleteVoDsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVoDsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVoDs",
		Method:             "DELETE",
		PathPattern:        "/v2/vods/bulk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVoDsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVoDsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVoDs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTotalVodNumber gets the partial number of vo ds depending on the searched items
*/
func (a *Client) GetTotalVodNumber(params *GetTotalVodNumberParams, opts ...ClientOption) (*GetTotalVodNumberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTotalVodNumberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTotalVodNumber",
		Method:             "GET",
		PathPattern:        "/v2/vods/count/{search}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTotalVodNumberReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTotalVodNumberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTotalVodNumber: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVoD vos d file from database
*/
func (a *Client) GetVoD(params *GetVoDParams, opts ...ClientOption) (*GetVoDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVoDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVoD",
		Method:             "GET",
		PathPattern:        "/v2/vods/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVoDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVoDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVoD: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVodList gets the vo d list from database it returns 50 items at max you can use offset value to get result page by page
*/
func (a *Client) GetVodList(params *GetVodListParams, opts ...ClientOption) (*GetVodListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVodListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVodList",
		Method:             "GET",
		PathPattern:        "/v2/vods/list/{offset}/{size}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVodListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVodListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVodList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportVoDsToStalker imports vo ds to stalker portal
*/
func (a *Client) ImportVoDsToStalker(params *ImportVoDsToStalkerParams, opts ...ClientOption) (*ImportVoDsToStalkerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportVoDsToStalkerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "importVoDsToStalker",
		Method:             "POST",
		PathPattern:        "/v2/vods/import-to-stalker",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ImportVoDsToStalkerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportVoDsToStalkerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importVoDsToStalker: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SynchUserVodList synchronizes vo d folder and add them to vo d database if any file exist and create symbolic links to that folder

  Notes here
*/
func (a *Client) SynchUserVodList(params *SynchUserVodListParams, opts ...ClientOption) (*SynchUserVodListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSynchUserVodListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "synchUserVodList",
		Method:             "POST",
		PathPattern:        "/v2/vods/synch-user-vod-list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SynchUserVodListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SynchUserVodListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for synchUserVodList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadVoDFile uploads external vo d file to ant media server
*/
func (a *Client) UploadVoDFile(params *UploadVoDFileParams, opts ...ClientOption) (*UploadVoDFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadVoDFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadVoDFile",
		Method:             "POST",
		PathPattern:        "/v2/vods/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadVoDFileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadVoDFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadVoDFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
