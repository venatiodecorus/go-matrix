// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

type SignedOneTimeKey struct {
	Key string `json:"key"`
	Signatures map[string]map[string]string `json:"signatures"`
}

// ClaimKeysOKBody claim keys o k body
// swagger:model claimKeysOKBody
type ClaimKeysOKBody struct {

	// If any remote homeservers could not be reached, they are
	// recorded here. The names of the properties are the names of
	// the unreachable servers.
	//
	// If the homeserver could be reached, but the user or device
	// was unknown, no failure is recorded. Instead, the corresponding
	// user or device is missing from the ``one_time_keys`` result.
	Failures map[string]interface{} `json:"failures,omitempty"`

	// One-time keys for the queried devices. A map from user ID, to a
	// map from devices to a map from ``<algorithm>:<key_id>`` to the key object.
	OneTimeKeys map[string]map[string]map[string]SignedOneTimeKey `json:"one_time_keys,omitempty"`
}

// Validate validates this claim keys o k body
func (m *ClaimKeysOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClaimKeysOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimKeysOKBody) UnmarshalBinary(b []byte) error {
	var res ClaimKeysOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
