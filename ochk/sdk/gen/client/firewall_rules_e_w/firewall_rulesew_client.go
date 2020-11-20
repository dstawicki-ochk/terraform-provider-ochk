// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_e_w

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new firewall rules e w API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for firewall rules e w API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DfwRuleCreateUsingPUT(params *DfwRuleCreateUsingPUTParams) (*DfwRuleCreateUsingPUTOK, *DfwRuleCreateUsingPUTCreated, error)

	DfwRuleDeleteUsingDELETE(params *DfwRuleDeleteUsingDELETEParams) (*DfwRuleDeleteUsingDELETEOK, error)

	DfwRuleGetUsingGET(params *DfwRuleGetUsingGETParams) (*DfwRuleGetUsingGETOK, error)

	DfwRuleListUsingGET(params *DfwRuleListUsingGETParams) (*DfwRuleListUsingGETOK, error)

	DfwRuleUpdateUsingPUT(params *DfwRuleUpdateUsingPUTParams) (*DfwRuleUpdateUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DfwRuleCreateUsingPUT creates

  Create firewall rule (east-west) in NSX-T
*/
func (a *Client) DfwRuleCreateUsingPUT(params *DfwRuleCreateUsingPUTParams) (*DfwRuleCreateUsingPUTOK, *DfwRuleCreateUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDfwRuleCreateUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dfwRuleCreateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/network/firewall/security-policies/{securityPolicyId}/rules/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DfwRuleCreateUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DfwRuleCreateUsingPUTOK:
		return value, nil, nil
	case *DfwRuleCreateUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for firewall_rules_e_w: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DfwRuleDeleteUsingDELETE deletes

  Delete firewall rule (east-west) from NSX-T
*/
func (a *Client) DfwRuleDeleteUsingDELETE(params *DfwRuleDeleteUsingDELETEParams) (*DfwRuleDeleteUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDfwRuleDeleteUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dfwRuleDeleteUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/network/firewall/security-policies/{securityPolicyId}/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DfwRuleDeleteUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DfwRuleDeleteUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dfwRuleDeleteUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DfwRuleGetUsingGET gets

  Get firewall rule (east-west) from NSX-T
*/
func (a *Client) DfwRuleGetUsingGET(params *DfwRuleGetUsingGETParams) (*DfwRuleGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDfwRuleGetUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dfwRuleGetUsingGET",
		Method:             "GET",
		PathPattern:        "/network/firewall/security-policies/{securityPolicyId}/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DfwRuleGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DfwRuleGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dfwRuleGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DfwRuleListUsingGET lists

  List firewall rules (east-west) from NSX-T
*/
func (a *Client) DfwRuleListUsingGET(params *DfwRuleListUsingGETParams) (*DfwRuleListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDfwRuleListUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dfwRuleListUsingGET",
		Method:             "GET",
		PathPattern:        "/network/firewall/security-policies/{securityPolicyId}/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DfwRuleListUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DfwRuleListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dfwRuleListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DfwRuleUpdateUsingPUT updates

  Update firewall rule (east-west) in NSX-T
*/
func (a *Client) DfwRuleUpdateUsingPUT(params *DfwRuleUpdateUsingPUTParams) (*DfwRuleUpdateUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDfwRuleUpdateUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dfwRuleUpdateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/network/firewall/security-policies/{securityPolicyId}/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DfwRuleUpdateUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DfwRuleUpdateUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dfwRuleUpdateUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
