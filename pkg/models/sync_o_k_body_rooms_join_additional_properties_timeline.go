// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SyncOKBodyRoomsJoinAdditionalPropertiesTimeline Timeline
//
// The timeline of messages and state changes in the
// room.
// swagger:model syncOKBodyRoomsJoinAdditionalPropertiesTimeline
type SyncOKBodyRoomsJoinAdditionalPropertiesTimeline struct {
	SyncOKBodyRoomsJoinAdditionalPropertiesTimelineAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SyncOKBodyRoomsJoinAdditionalPropertiesTimeline) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SyncOKBodyRoomsJoinAdditionalPropertiesTimelineAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SyncOKBodyRoomsJoinAdditionalPropertiesTimelineAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SyncOKBodyRoomsJoinAdditionalPropertiesTimeline) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SyncOKBodyRoomsJoinAdditionalPropertiesTimelineAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sync o k body rooms join additional properties timeline
func (m *SyncOKBodyRoomsJoinAdditionalPropertiesTimeline) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SyncOKBodyRoomsJoinAdditionalPropertiesTimelineAllOf0
	if err := m.SyncOKBodyRoomsJoinAdditionalPropertiesTimelineAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SyncOKBodyRoomsJoinAdditionalPropertiesTimeline) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncOKBodyRoomsJoinAdditionalPropertiesTimeline) UnmarshalBinary(b []byte) error {
	var res SyncOKBodyRoomsJoinAdditionalPropertiesTimeline
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
