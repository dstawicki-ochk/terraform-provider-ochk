// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VcsVirtualMachineInstance VcsVirtualMachineInstance
//
// swagger:model VcsVirtualMachineInstance
type VcsVirtualMachineInstance struct {

	// additional virtual disk device collection
	AdditionalVirtualDiskDeviceCollection []*VirtualDiskDevice `json:"additionalVirtualDiskDeviceCollection"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate,omitempty"`

	// deployment instance
	DeploymentInstance *DeploymentInstance `json:"deploymentInstance,omitempty"`

	// initial password
	InitialPassword string `json:"initialPassword,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate *strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// os virtual disk device
	OsVirtualDiskDevice *VirtualDiskDevice `json:"osVirtualDiskDevice,omitempty"`

	// power state
	// Enum: [poweredOff poweredOn suspended]
	PowerState string `json:"powerState,omitempty"`

	// resource profile
	// Enum: [CUSTOM SIZE_L SIZE_M SIZE_S SIZE_XL SIZE_XS]
	ResourceProfile string `json:"resourceProfile,omitempty"`

	// storage policy
	// Enum: [ENTERPRISE STANDARD UNKNOWN]
	StoragePolicy string `json:"storagePolicy,omitempty"`

	// subtenant ref Id
	SubtenantRefID string `json:"subtenantRefId,omitempty"`

	// virtual machine Id
	VirtualMachineID string `json:"virtualMachineId,omitempty"`

	// virtual machine name
	VirtualMachineName string `json:"virtualMachineName,omitempty"`

	// virtual network devices
	VirtualNetworkDevices []*VirtualNetworkDevice `json:"virtualNetworkDevices"`
}

// Validate validates this vcs virtual machine instance
func (m *VcsVirtualMachineInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalVirtualDiskDeviceCollection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsVirtualDiskDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsVirtualMachineInstance) validateAdditionalVirtualDiskDeviceCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalVirtualDiskDeviceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalVirtualDiskDeviceCollection); i++ {
		if swag.IsZero(m.AdditionalVirtualDiskDeviceCollection[i]) { // not required
			continue
		}

		if m.AdditionalVirtualDiskDeviceCollection[i] != nil {
			if err := m.AdditionalVirtualDiskDeviceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalVirtualDiskDeviceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateDeploymentInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.DeploymentInstance) { // not required
		return nil
	}

	if m.DeploymentInstance != nil {
		if err := m.DeploymentInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateModificationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateOsVirtualDiskDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.OsVirtualDiskDevice) { // not required
		return nil
	}

	if m.OsVirtualDiskDevice != nil {
		if err := m.OsVirtualDiskDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osVirtualDiskDevice")
			}
			return err
		}
	}

	return nil
}

var vcsVirtualMachineInstanceTypePowerStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["poweredOff","poweredOn","suspended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vcsVirtualMachineInstanceTypePowerStatePropEnum = append(vcsVirtualMachineInstanceTypePowerStatePropEnum, v)
	}
}

const (

	// VcsVirtualMachineInstancePowerStatePoweredOff captures enum value "poweredOff"
	VcsVirtualMachineInstancePowerStatePoweredOff string = "poweredOff"

	// VcsVirtualMachineInstancePowerStatePoweredOn captures enum value "poweredOn"
	VcsVirtualMachineInstancePowerStatePoweredOn string = "poweredOn"

	// VcsVirtualMachineInstancePowerStateSuspended captures enum value "suspended"
	VcsVirtualMachineInstancePowerStateSuspended string = "suspended"
)

// prop value enum
func (m *VcsVirtualMachineInstance) validatePowerStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vcsVirtualMachineInstanceTypePowerStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VcsVirtualMachineInstance) validatePowerState(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerState) { // not required
		return nil
	}

	// value enum
	if err := m.validatePowerStateEnum("powerState", "body", m.PowerState); err != nil {
		return err
	}

	return nil
}

var vcsVirtualMachineInstanceTypeResourceProfilePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","SIZE_L","SIZE_M","SIZE_S","SIZE_XL","SIZE_XS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vcsVirtualMachineInstanceTypeResourceProfilePropEnum = append(vcsVirtualMachineInstanceTypeResourceProfilePropEnum, v)
	}
}

const (

	// VcsVirtualMachineInstanceResourceProfileCUSTOM captures enum value "CUSTOM"
	VcsVirtualMachineInstanceResourceProfileCUSTOM string = "CUSTOM"

	// VcsVirtualMachineInstanceResourceProfileSIZEL captures enum value "SIZE_L"
	VcsVirtualMachineInstanceResourceProfileSIZEL string = "SIZE_L"

	// VcsVirtualMachineInstanceResourceProfileSIZEM captures enum value "SIZE_M"
	VcsVirtualMachineInstanceResourceProfileSIZEM string = "SIZE_M"

	// VcsVirtualMachineInstanceResourceProfileSIZES captures enum value "SIZE_S"
	VcsVirtualMachineInstanceResourceProfileSIZES string = "SIZE_S"

	// VcsVirtualMachineInstanceResourceProfileSIZEXL captures enum value "SIZE_XL"
	VcsVirtualMachineInstanceResourceProfileSIZEXL string = "SIZE_XL"

	// VcsVirtualMachineInstanceResourceProfileSIZEXS captures enum value "SIZE_XS"
	VcsVirtualMachineInstanceResourceProfileSIZEXS string = "SIZE_XS"
)

// prop value enum
func (m *VcsVirtualMachineInstance) validateResourceProfileEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vcsVirtualMachineInstanceTypeResourceProfilePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VcsVirtualMachineInstance) validateResourceProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceProfile) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceProfileEnum("resourceProfile", "body", m.ResourceProfile); err != nil {
		return err
	}

	return nil
}

var vcsVirtualMachineInstanceTypeStoragePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENTERPRISE","STANDARD","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vcsVirtualMachineInstanceTypeStoragePolicyPropEnum = append(vcsVirtualMachineInstanceTypeStoragePolicyPropEnum, v)
	}
}

const (

	// VcsVirtualMachineInstanceStoragePolicyENTERPRISE captures enum value "ENTERPRISE"
	VcsVirtualMachineInstanceStoragePolicyENTERPRISE string = "ENTERPRISE"

	// VcsVirtualMachineInstanceStoragePolicySTANDARD captures enum value "STANDARD"
	VcsVirtualMachineInstanceStoragePolicySTANDARD string = "STANDARD"

	// VcsVirtualMachineInstanceStoragePolicyUNKNOWN captures enum value "UNKNOWN"
	VcsVirtualMachineInstanceStoragePolicyUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *VcsVirtualMachineInstance) validateStoragePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vcsVirtualMachineInstanceTypeStoragePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VcsVirtualMachineInstance) validateStoragePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.StoragePolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateStoragePolicyEnum("storagePolicy", "body", m.StoragePolicy); err != nil {
		return err
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateVirtualNetworkDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualNetworkDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualNetworkDevices); i++ {
		if swag.IsZero(m.VirtualNetworkDevices[i]) { // not required
			continue
		}

		if m.VirtualNetworkDevices[i] != nil {
			if err := m.VirtualNetworkDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualNetworkDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcsVirtualMachineInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsVirtualMachineInstance) UnmarshalBinary(b []byte) error {
	var res VcsVirtualMachineInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}