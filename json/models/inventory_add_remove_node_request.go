// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryAddRemoveNodeRequest inventory add remove node request
// swagger:model inventoryAddRemoveNodeRequest
type InventoryAddRemoveNodeRequest struct {

	// Unique node name.
	Name string `json:"name,omitempty"`
}

// Validate validates this inventory add remove node request
func (m *InventoryAddRemoveNodeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryAddRemoveNodeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryAddRemoveNodeRequest) UnmarshalBinary(b []byte) error {
	var res InventoryAddRemoveNodeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}