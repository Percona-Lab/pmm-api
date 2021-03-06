// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InventoryAddNodeResponse inventory add node response
// swagger:model inventoryAddNodeResponse
type InventoryAddNodeResponse struct {

	// bare metal
	BareMetal *InventoryBareMetalNode `json:"bare_metal,omitempty"`

	// container
	Container *InventoryContainerNode `json:"container,omitempty"`

	// virtual machine
	VirtualMachine *InventoryVirtualMachineNode `json:"virtual_machine,omitempty"`
}

// Validate validates this inventory add node response
func (m *InventoryAddNodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBareMetal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryAddNodeResponse) validateBareMetal(formats strfmt.Registry) error {

	if swag.IsZero(m.BareMetal) { // not required
		return nil
	}

	if m.BareMetal != nil {
		if err := m.BareMetal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bare_metal")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryAddNodeResponse) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(m.Container) { // not required
		return nil
	}

	if m.Container != nil {
		if err := m.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryAddNodeResponse) validateVirtualMachine(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMachine) { // not required
		return nil
	}

	if m.VirtualMachine != nil {
		if err := m.VirtualMachine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_machine")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryAddNodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryAddNodeResponse) UnmarshalBinary(b []byte) error {
	var res InventoryAddNodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
