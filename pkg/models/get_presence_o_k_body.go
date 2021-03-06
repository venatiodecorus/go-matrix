// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetPresenceOKBody get presence o k body
// swagger:model getPresenceOKBody
type GetPresenceOKBody struct {

	// Whether the user is currently active
	CurrentlyActive bool `json:"currently_active,omitempty"`

	// The length of time in milliseconds since an action was performed
	// by this user.
	LastActiveAgo int64 `json:"last_active_ago,omitempty"`

	// This user's presence.
	// Required: true
	// Enum: [online offline unavailable]
	Presence *string `json:"presence"`

	// The state message for this user if one was set.
	StatusMsg string `json:"status_msg,omitempty"`
}

// Validate validates this get presence o k body
func (m *GetPresenceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePresence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getPresenceOKBodyTypePresencePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","offline","unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getPresenceOKBodyTypePresencePropEnum = append(getPresenceOKBodyTypePresencePropEnum, v)
	}
}

const (

	// GetPresenceOKBodyPresenceOnline captures enum value "online"
	GetPresenceOKBodyPresenceOnline string = "online"

	// GetPresenceOKBodyPresenceOffline captures enum value "offline"
	GetPresenceOKBodyPresenceOffline string = "offline"

	// GetPresenceOKBodyPresenceUnavailable captures enum value "unavailable"
	GetPresenceOKBodyPresenceUnavailable string = "unavailable"
)

// prop value enum
func (m *GetPresenceOKBody) validatePresenceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getPresenceOKBodyTypePresencePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetPresenceOKBody) validatePresence(formats strfmt.Registry) error {

	if err := validate.Required("presence", "body", m.Presence); err != nil {
		return err
	}

	// value enum
	if err := m.validatePresenceEnum("presence", "body", *m.Presence); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetPresenceOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPresenceOKBody) UnmarshalBinary(b []byte) error {
	var res GetPresenceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
