// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new snapshots API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for snapshots API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SnapshotGetUsingGET(params *SnapshotGetUsingGETParams) (*SnapshotGetUsingGETOK, error)

	SnapshotListUsingGET(params *SnapshotListUsingGETParams) (*SnapshotListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SnapshotGetUsingGET gets

  Get vSphere vCenter virtual machine snapshot
*/
func (a *Client) SnapshotGetUsingGET(params *SnapshotGetUsingGETParams) (*SnapshotGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotGetUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "snapshotGetUsingGET",
		Method:             "GET",
		PathPattern:        "/vcs/snapshots/{snapshotId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SnapshotGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for snapshotGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SnapshotListUsingGET lists

  List vSphere vCenter virtual machines snapshots
*/
func (a *Client) SnapshotListUsingGET(params *SnapshotListUsingGETParams) (*SnapshotListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotListUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "snapshotListUsingGET",
		Method:             "GET",
		PathPattern:        "/vcs/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotListUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SnapshotListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for snapshotListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}