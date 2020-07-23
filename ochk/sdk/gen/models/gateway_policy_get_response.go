// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayPolicyGetResponse gateway policy get response
//
// swagger:model GatewayPolicyGetResponse
type GatewayPolicyGetResponse struct {

	// gateway policy
	GatewayPolicy *GatewayPolicy `json:"gatewayPolicy,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this gateway policy get response
func (m *GatewayPolicyGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGatewayPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayPolicyGetResponse) validateGatewayPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayPolicy) { // not required
		return nil
	}

	if m.GatewayPolicy != nil {
		if err := m.GatewayPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gatewayPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayPolicyGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayPolicyGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayPolicyGetResponse) UnmarshalBinary(b []byte) error {
	var res GatewayPolicyGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
