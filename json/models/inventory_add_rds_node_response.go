// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InventoryAddRDSNodeResponse inventory add RDS node response
// swagger:model inventoryAddRDSNodeResponse
type InventoryAddRDSNodeResponse struct {

	// rds
	RDS *InventoryRDSNode `json:"rds,omitempty"`
}

// Validate validates this inventory add RDS node response
func (m *InventoryAddRDSNodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRDS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryAddRDSNodeResponse) validateRDS(formats strfmt.Registry) error {

	if swag.IsZero(m.RDS) { // not required
		return nil
	}

	if m.RDS != nil {
		if err := m.RDS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryAddRDSNodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryAddRDSNodeResponse) UnmarshalBinary(b []byte) error {
	var res InventoryAddRDSNodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
