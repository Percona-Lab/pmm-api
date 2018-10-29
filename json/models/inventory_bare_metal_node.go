// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryBareMetalNode BareMetalNode represents bare metal node configuration.
// swagger:model inventoryBareMetalNode
type InventoryBareMetalNode struct {

	// Hostname. Is not unique.
	Hostname string `json:"hostname,omitempty"`

	// Unique node identifier.
	ID int64 `json:"id,omitempty"`

	// Unique node name.
	Name string `json:"name,omitempty"`
}

// Validate validates this inventory bare metal node
func (m *InventoryBareMetalNode) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryBareMetalNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryBareMetalNode) UnmarshalBinary(b []byte) error {
	var res InventoryBareMetalNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
