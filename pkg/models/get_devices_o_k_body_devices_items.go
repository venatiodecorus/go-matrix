// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetDevicesOKBodyDevicesItems get devices o k body devices items
// swagger:model getDevicesOKBodyDevicesItems
type GetDevicesOKBodyDevicesItems struct {
	GetDevicesOKBodyDevicesItemsAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetDevicesOKBodyDevicesItems) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetDevicesOKBodyDevicesItemsAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetDevicesOKBodyDevicesItemsAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetDevicesOKBodyDevicesItems) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.GetDevicesOKBodyDevicesItemsAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get devices o k body devices items
func (m *GetDevicesOKBodyDevicesItems) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetDevicesOKBodyDevicesItemsAllOf0
	if err := m.GetDevicesOKBodyDevicesItemsAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetDevicesOKBodyDevicesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDevicesOKBodyDevicesItems) UnmarshalBinary(b []byte) error {
	var res GetDevicesOKBodyDevicesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
