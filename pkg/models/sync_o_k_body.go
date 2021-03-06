// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SyncOKBody sync o k body
// swagger:model syncOKBody
type SyncOKBody struct {

	// account data
	AccountData *SyncOKBodyAccountData `json:"account_data,omitempty"`

	// DeviceLists
	//
	// Information on end-to-end device updates, as specified in
	// |device_lists_sync|_.
	DeviceLists interface{} `json:"device_lists,omitempty"`

	// One-time keys count
	//
	// Information on end-to-end encryption keys, as specified
	// in |device_lists_sync|_.
	DeviceOneTimeKeysCount map[string]int64 `json:"device_one_time_keys_count,omitempty"`

	// The batch token to supply in the ``since`` param of the next
	// ``/sync`` request.
	// Required: true
	NextBatch *string `json:"next_batch"`

	// presence
	Presence *SyncOKBodyPresence `json:"presence,omitempty"`

	// rooms
	Rooms *SyncOKBodyRooms `json:"rooms,omitempty"`

	// ToDevice
	//
	// Information on the send-to-device messages for the client
	// device, as defined in |send_to_device_sync|_.
	ToDevice interface{} `json:"to_device,omitempty"`
}

// Validate validates this sync o k body
func (m *SyncOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextBatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRooms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncOKBody) validateAccountData(formats strfmt.Registry) error {

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

func (m *SyncOKBody) validateNextBatch(formats strfmt.Registry) error {

	if err := validate.Required("next_batch", "body", m.NextBatch); err != nil {
		return err
	}

	return nil
}

func (m *SyncOKBody) validatePresence(formats strfmt.Registry) error {

	if swag.IsZero(m.Presence) { // not required
		return nil
	}

	if m.Presence != nil {
		if err := m.Presence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("presence")
			}
			return err
		}
	}

	return nil
}

func (m *SyncOKBody) validateRooms(formats strfmt.Registry) error {

	if swag.IsZero(m.Rooms) { // not required
		return nil
	}

	if m.Rooms != nil {
		if err := m.Rooms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rooms")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncOKBody) UnmarshalBinary(b []byte) error {
	var res SyncOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
