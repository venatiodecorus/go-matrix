// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchParamsBodySearchCategoriesRoomEvents Room Events Criteria
//
// Mapping of category name to search criteria.
// swagger:model searchParamsBodySearchCategoriesRoomEvents
type SearchParamsBodySearchCategoriesRoomEvents struct {

	// event context
	EventContext *SearchParamsBodySearchCategoriesRoomEventsEventContext `json:"event_context,omitempty"`

	// filter
	Filter *SearchParamsBodySearchCategoriesRoomEventsFilter `json:"filter,omitempty"`

	// groupings
	Groupings *SearchParamsBodySearchCategoriesRoomEventsGroupings `json:"groupings,omitempty"`

	// Include current state
	//
	// Requests the server return the current state for
	// each room returned.
	IncludeState bool `json:"include_state,omitempty"`

	// The keys to search. Defaults to all.
	Keys []string `json:"keys"`

	// Ordering
	//
	// The order in which to search for results.
	// By default, this is ``"rank"``.
	// Enum: [recent rank]
	OrderBy string `json:"order_by,omitempty"`

	// The string to search events for
	// Required: true
	SearchTerm *string `json:"search_term"`
}

// Validate validates this search params body search categories room events
func (m *SearchParamsBodySearchCategoriesRoomEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchTerm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchParamsBodySearchCategoriesRoomEvents) validateEventContext(formats strfmt.Registry) error {

	if swag.IsZero(m.EventContext) { // not required
		return nil
	}

	if m.EventContext != nil {
		if err := m.EventContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_context")
			}
			return err
		}
	}

	return nil
}

func (m *SearchParamsBodySearchCategoriesRoomEvents) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchParamsBodySearchCategoriesRoomEvents) validateGroupings(formats strfmt.Registry) error {

	if swag.IsZero(m.Groupings) { // not required
		return nil
	}

	if m.Groupings != nil {
		if err := m.Groupings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupings")
			}
			return err
		}
	}

	return nil
}

var searchParamsBodySearchCategoriesRoomEventsKeysItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["content.body","content.name","content.topic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchParamsBodySearchCategoriesRoomEventsKeysItemsEnum = append(searchParamsBodySearchCategoriesRoomEventsKeysItemsEnum, v)
	}
}

func (m *SearchParamsBodySearchCategoriesRoomEvents) validateKeysItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, searchParamsBodySearchCategoriesRoomEventsKeysItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *SearchParamsBodySearchCategoriesRoomEvents) validateKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.Keys) { // not required
		return nil
	}

	for i := 0; i < len(m.Keys); i++ {

		// value enum
		if err := m.validateKeysItemsEnum("keys"+"."+strconv.Itoa(i), "body", m.Keys[i]); err != nil {
			return err
		}

	}

	return nil
}

var searchParamsBodySearchCategoriesRoomEventsTypeOrderByPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["recent","rank"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchParamsBodySearchCategoriesRoomEventsTypeOrderByPropEnum = append(searchParamsBodySearchCategoriesRoomEventsTypeOrderByPropEnum, v)
	}
}

const (

	// SearchParamsBodySearchCategoriesRoomEventsOrderByRecent captures enum value "recent"
	SearchParamsBodySearchCategoriesRoomEventsOrderByRecent string = "recent"

	// SearchParamsBodySearchCategoriesRoomEventsOrderByRank captures enum value "rank"
	SearchParamsBodySearchCategoriesRoomEventsOrderByRank string = "rank"
)

// prop value enum
func (m *SearchParamsBodySearchCategoriesRoomEvents) validateOrderByEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, searchParamsBodySearchCategoriesRoomEventsTypeOrderByPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SearchParamsBodySearchCategoriesRoomEvents) validateOrderBy(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrderByEnum("order_by", "body", m.OrderBy); err != nil {
		return err
	}

	return nil
}

func (m *SearchParamsBodySearchCategoriesRoomEvents) validateSearchTerm(formats strfmt.Registry) error {

	if err := validate.Required("search_term", "body", m.SearchTerm); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchParamsBodySearchCategoriesRoomEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchParamsBodySearchCategoriesRoomEvents) UnmarshalBinary(b []byte) error {
	var res SearchParamsBodySearchCategoriesRoomEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
