// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetFilterOKBodyAllOf0RoomAccountData The per user account data to include for rooms.
// swagger:model getFilterOKBodyAllOf0RoomAccountData
type GetFilterOKBodyAllOf0RoomAccountData struct {
	GetFilterOKBodyAllOf0RoomAccountDataAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetFilterOKBodyAllOf0RoomAccountData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetFilterOKBodyAllOf0RoomAccountDataAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetFilterOKBodyAllOf0RoomAccountDataAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetFilterOKBodyAllOf0RoomAccountData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.GetFilterOKBodyAllOf0RoomAccountDataAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get filter o k body all of0 room account data
func (m *GetFilterOKBodyAllOf0RoomAccountData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetFilterOKBodyAllOf0RoomAccountDataAllOf0
	if err := m.GetFilterOKBodyAllOf0RoomAccountDataAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetFilterOKBodyAllOf0RoomAccountData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFilterOKBodyAllOf0RoomAccountData) UnmarshalBinary(b []byte) error {
	var res GetFilterOKBodyAllOf0RoomAccountData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
