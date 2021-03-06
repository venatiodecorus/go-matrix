// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems sync o k body rooms join additional properties ephemeral all of0 events items
// swagger:model syncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems
type SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems struct {
	SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItemsAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItemsAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItemsAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItemsAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sync o k body rooms join additional properties ephemeral all of0 events items
func (m *SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItemsAllOf0
	if err := m.SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItemsAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems) UnmarshalBinary(b []byte) error {
	var res SyncOKBodyRoomsJoinAdditionalPropertiesEphemeralAllOf0EventsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
