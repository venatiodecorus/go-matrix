// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// QueryPublicRoomsParamsBodyFilter Filter
//
// Filter to apply to the results.
// swagger:model queryPublicRoomsParamsBodyFilter
type QueryPublicRoomsParamsBodyFilter struct {

	// A string to search for in the room metadata, e.g. name,
	// topic, canonical alias etc. (Optional).
	GenericSearchTerm string `json:"generic_search_term,omitempty"`
}

// Validate validates this query public rooms params body filter
func (m *QueryPublicRoomsParamsBodyFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryPublicRoomsParamsBodyFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryPublicRoomsParamsBodyFilter) UnmarshalBinary(b []byte) error {
	var res QueryPublicRoomsParamsBodyFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
