// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetEventContextOKBodyEvent Details of the requested event.
// swagger:model getEventContextOKBodyEvent
type GetEventContextOKBodyEvent struct {
	GetEventContextOKBodyEventAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetEventContextOKBodyEvent) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetEventContextOKBodyEventAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetEventContextOKBodyEventAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetEventContextOKBodyEvent) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.GetEventContextOKBodyEventAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get event context o k body event
func (m *GetEventContextOKBodyEvent) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetEventContextOKBodyEventAllOf0
	if err := m.GetEventContextOKBodyEventAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetEventContextOKBodyEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEventContextOKBodyEvent) UnmarshalBinary(b []byte) error {
	var res GetEventContextOKBodyEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
