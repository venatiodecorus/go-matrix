// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetOneEventOKBody get one event o k body
// swagger:model getOneEventOKBody
type GetOneEventOKBody struct {
	GetOneEventOKBodyAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetOneEventOKBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetOneEventOKBodyAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetOneEventOKBodyAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetOneEventOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.GetOneEventOKBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get one event o k body
func (m *GetOneEventOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetOneEventOKBodyAllOf0
	if err := m.GetOneEventOKBodyAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetOneEventOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOneEventOKBody) UnmarshalBinary(b []byte) error {
	var res GetOneEventOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
