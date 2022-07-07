// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Bind Bind
//
// HAProxy frontend bind configuration
//
// swagger:model bind
type Bind struct {
	BindParams

	// address
	// Pattern: ^[^\s]+$
	Address string `json:"address,omitempty"`

	// port
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`

	// port range end
	// Maximum: 65535
	// Minimum: 1
	PortRangeEnd *int64 `json:"port-range-end,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Bind) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BindParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BindParams = aO0

	// now for regular properties
	var propsBind struct {
		Address string `json:"address,omitempty"`

		Port *int64 `json:"port,omitempty"`

		PortRangeEnd *int64 `json:"port-range-end,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsBind); err != nil {
		return err
	}
	m.Address = propsBind.Address

	m.Port = propsBind.Port

	m.PortRangeEnd = propsBind.PortRangeEnd

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Bind) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BindParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsBind struct {
		Address string `json:"address,omitempty"`

		Port *int64 `json:"port,omitempty"`

		PortRangeEnd *int64 `json:"port-range-end,omitempty"`
	}
	propsBind.Address = m.Address

	propsBind.Port = m.Port

	propsBind.PortRangeEnd = m.PortRangeEnd

	jsonDataPropsBind, errBind := swag.WriteJSON(propsBind)
	if errBind != nil {
		return nil, errBind
	}
	_parts = append(_parts, jsonDataPropsBind)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this bind
func (m *Bind) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BindParams
	if err := m.BindParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortRangeEnd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Bind) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.Pattern("address", "body", m.Address, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", *m.Port, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", *m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validatePortRangeEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.PortRangeEnd) { // not required
		return nil
	}

	if err := validate.MinimumInt("port-range-end", "body", *m.PortRangeEnd, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port-range-end", "body", *m.PortRangeEnd, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bind based on the context it is used
func (m *Bind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BindParams
	if err := m.BindParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Bind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bind) UnmarshalBinary(b []byte) error {
	var res Bind
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
