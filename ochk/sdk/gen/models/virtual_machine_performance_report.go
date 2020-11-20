// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualMachinePerformanceReport VirtualMachinePerformanceReport
//
// swagger:model VirtualMachinePerformanceReport
type VirtualMachinePerformanceReport struct {

	// availability percentage
	AvailabilityPercentage float32 `json:"availabilityPercentage,omitempty"`

	// avg Cpu utilization
	AvgCPUUtilization float32 `json:"avgCpuUtilization,omitempty"`

	// avg Ram utilization
	AvgRAMUtilization float32 `json:"avgRamUtilization,omitempty"`

	// avg storage utilization
	AvgStorageUtilization float32 `json:"avgStorageUtilization,omitempty"`

	// max Cpu utilization
	MaxCPUUtilization float32 `json:"maxCpuUtilization,omitempty"`

	// max Ram utilization
	MaxRAMUtilization float32 `json:"maxRamUtilization,omitempty"`

	// max storage utilization
	MaxStorageUtilization float32 `json:"maxStorageUtilization,omitempty"`

	// storage total used
	StorageTotalUsed float32 `json:"storageTotalUsed,omitempty"`

	// virtual machine guest
	VirtualMachineGuest string `json:"virtualMachineGuest,omitempty"`

	// virtual machine Id
	VirtualMachineID string `json:"virtualMachineId,omitempty"`

	// virtual machine name
	VirtualMachineName string `json:"virtualMachineName,omitempty"`
}

// Validate validates this virtual machine performance report
func (m *VirtualMachinePerformanceReport) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachinePerformanceReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachinePerformanceReport) UnmarshalBinary(b []byte) error {
	var res VirtualMachinePerformanceReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
