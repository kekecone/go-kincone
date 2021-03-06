// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FlagModel 休日等の種別
//
// swagger:model flagModel
type FlagModel struct {

	// 休日等の種別名
	Name string `json:"name,omitempty"`
}

// Validate validates this flag model
func (m *FlagModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this flag model based on context it is used
func (m *FlagModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlagModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlagModel) UnmarshalBinary(b []byte) error {
	var res FlagModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
