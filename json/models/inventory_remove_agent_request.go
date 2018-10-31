// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryRemoveAgentRequest inventory remove agent request
// swagger:model inventoryRemoveAgentRequest
type InventoryRemoveAgentRequest struct {

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this inventory remove agent request
func (m *InventoryRemoveAgentRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryRemoveAgentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryRemoveAgentRequest) UnmarshalBinary(b []byte) error {
	var res InventoryRemoveAgentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
