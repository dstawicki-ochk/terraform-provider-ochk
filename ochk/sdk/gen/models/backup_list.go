// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupList BackupList
//
// swagger:model BackupList
type BackupList struct {

	// backup list code
	BackupListCode string `json:"backupListCode,omitempty"`

	// backup list Id
	BackupListID string `json:"backupListId,omitempty"`

	// backup list name
	BackupListName string `json:"backupListName,omitempty"`
}

// Validate validates this backup list
func (m *BackupList) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupList) UnmarshalBinary(b []byte) error {
	var res BackupList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
