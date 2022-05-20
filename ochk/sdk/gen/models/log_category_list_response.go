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

// LogCategoryListResponse LogCategoryListResponse
//
// swagger:model LogCategoryListResponse
type LogCategoryListResponse struct {

	// log category collection
	LogCategoryCollection []*LogCategory `json:"logCategoryCollection"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this log category list response
func (m *LogCategoryListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogCategoryCollection(formats); err != nil {
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

func (m *LogCategoryListResponse) validateLogCategoryCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.LogCategoryCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.LogCategoryCollection); i++ {
		if swag.IsZero(m.LogCategoryCollection[i]) { // not required
			continue
		}

		if m.LogCategoryCollection[i] != nil {
			if err := m.LogCategoryCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logCategoryCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogCategoryListResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogCategoryListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogCategoryListResponse) UnmarshalBinary(b []byte) error {
	var res LogCategoryListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}