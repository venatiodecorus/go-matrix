// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DefineFilterParamsBodyAllOf0Room RoomFilter
//
// Filters to be applied to room data.
// swagger:model defineFilterParamsBodyAllOf0Room
type DefineFilterParamsBodyAllOf0Room struct {

	// account data
	AccountData *DefineFilterParamsBodyAllOf0RoomAccountData `json:"account_data,omitempty"`

	// ephemeral
	Ephemeral *DefineFilterParamsBodyAllOf0RoomEphemeral `json:"ephemeral,omitempty"`

	// Include rooms that the user has left in the sync, default false
	IncludeLeave bool `json:"include_leave,omitempty"`

	// A list of room IDs to exclude. If this list is absent then no rooms are excluded. A matching room will be excluded even if it is listed in the ``'rooms'`` filter. This filter is applied before the filters in ``ephemeral``, ``state``, ``timeline`` or ``account_data``
	NotRooms []string `json:"not_rooms"`

	// A list of room IDs to include. If this list is absent then all rooms are included. This filter is applied before the filters in ``ephemeral``, ``state``, ``timeline`` or ``account_data``
	Rooms []string `json:"rooms"`

	// state
	State *DefineFilterParamsBodyAllOf0RoomState `json:"state,omitempty"`

	// timeline
	Timeline *DefineFilterParamsBodyAllOf0RoomTimeline `json:"timeline,omitempty"`
}

// Validate validates this define filter params body all of0 room
func (m *DefineFilterParamsBodyAllOf0Room) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEphemeral(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DefineFilterParamsBodyAllOf0Room) validateAccountData(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountData) { // not required
		return nil
	}

	if m.AccountData != nil {
		if err := m.AccountData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_data")
			}
			return err
		}
	}

	return nil
}

func (m *DefineFilterParamsBodyAllOf0Room) validateEphemeral(formats strfmt.Registry) error {

	if swag.IsZero(m.Ephemeral) { // not required
		return nil
	}

	if m.Ephemeral != nil {
		if err := m.Ephemeral.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeral")
			}
			return err
		}
	}

	return nil
}

func (m *DefineFilterParamsBodyAllOf0Room) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *DefineFilterParamsBodyAllOf0Room) validateTimeline(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeline) { // not required
		return nil
	}

	if m.Timeline != nil {
		if err := m.Timeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeline")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DefineFilterParamsBodyAllOf0Room) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefineFilterParamsBodyAllOf0Room) UnmarshalBinary(b []byte) error {
	var res DefineFilterParamsBodyAllOf0Room
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
