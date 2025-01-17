// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new networks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for networks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	VcsVirtualMachineGroupGetUsingGET(params *VcsVirtualMachineGroupGetUsingGETParams) (*VcsVirtualMachineGroupGetUsingGETOK, error)

	VcsVirtualMachineListUsingGET(params *VcsVirtualMachineListUsingGETParams) (*VcsVirtualMachineListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  VcsVirtualMachineGroupGetUsingGET gets

  Get vSphere vCenter networks
*/
func (a *Client) VcsVirtualMachineGroupGetUsingGET(params *VcsVirtualMachineGroupGetUsingGETParams) (*VcsVirtualMachineGroupGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVcsVirtualMachineGroupGetUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "vcsVirtualMachineGroupGetUsingGET",
		Method:             "GET",
		PathPattern:        "/vcs/networks/{networkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VcsVirtualMachineGroupGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VcsVirtualMachineGroupGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for vcsVirtualMachineGroupGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VcsVirtualMachineListUsingGET lists

  List vSphere vCenter networks
*/
func (a *Client) VcsVirtualMachineListUsingGET(params *VcsVirtualMachineListUsingGETParams) (*VcsVirtualMachineListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVcsVirtualMachineListUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "vcsVirtualMachineListUsingGET",
		Method:             "GET",
		PathPattern:        "/vcs/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VcsVirtualMachineListUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VcsVirtualMachineListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for vcsVirtualMachineListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
