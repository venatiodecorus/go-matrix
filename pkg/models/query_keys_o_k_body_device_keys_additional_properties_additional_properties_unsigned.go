// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// QueryKeysOKBodyDeviceKeysAdditionalPropertiesAdditionalPropertiesUnsigned UnsignedDeviceInfo
//
// Additional data added to the device key information
// by intermediate servers, and not covered by the
// signatures.
// swagger:model queryKeysOKBodyDeviceKeysAdditionalPropertiesAdditionalPropertiesUnsigned
type QueryKeysOKBodyDeviceKeysAdditionalPropertiesAdditionalPropertiesUnsigned struct {

	// The display name which the user set on the device.
	DeviceDisplayName string `json:"device_display_name,omitempty"`
}

// Validate validates this query keys o k body device keys additional properties additional properties unsigned
func (m *QueryKeysOKBodyDeviceKeysAdditionalPropertiesAdditionalPropertiesUnsigned) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryKeysOKBodyDeviceKeysAdditionalPropertiesAdditionalPropertiesUnsigned) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryKeysOKBodyDeviceKeysAdditionalPropertiesAdditionalPropertiesUnsigned) UnmarshalBinary(b []byte) error {
	var res QueryKeysOKBodyDeviceKeysAdditionalPropertiesAdditionalPropertiesUnsigned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
