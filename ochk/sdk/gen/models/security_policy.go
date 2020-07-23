// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityPolicy security policy
//
// swagger:model SecurityPolicy
type SecurityPolicy struct {

	// category
	Category *Category `json:"category,omitempty"`

	// connectivity strategy
	ConnectivityStrategy *ConnectivityStrategy `json:"connectivityStrategy,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// resource type
	ResourceType *ResourceType `json:"resourceType,omitempty"`

	// security policy Id
	SecurityPolicyID string `json:"securityPolicyId,omitempty"`

	// stateful
	Stateful bool `json:"stateful,omitempty"`

	// tags
	Tags []*TagInstance `json:"tags"`

	// tcp strict
	TCPStrict bool `json:"tcpStrict,omitempty"`
}

// Validate validates this security policy
func (m *SecurityPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectivityStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityPolicy) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityPolicy) validateConnectivityStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectivityStrategy) { // not required
		return nil
	}

	if m.ConnectivityStrategy != nil {
		if err := m.ConnectivityStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivityStrategy")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityPolicy) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SecurityPolicy) validateModificationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SecurityPolicy) validateResourceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	if m.ResourceType != nil {
		if err := m.ResourceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceType")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityPolicy) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityPolicy) UnmarshalBinary(b []byte) error {
	var res SecurityPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
