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

// MailersSection Mailers Section
//
// A list of SMTP servers used by HAProxy to send emails.
//
// swagger:model mailers_section
type MailersSection struct {
	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_]+$
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9-_]+$`
	Name string `json:"name"`

	// timeout
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	Timeout *int64 `json:"timeout,omitempty"`
}

// Validate validates this mailers section
func (m *MailersSection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MailersSection) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9-_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *MailersSection) validateTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeout", "body", *m.Timeout, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mailers section based on context it is used
func (m *MailersSection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MailersSection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MailersSection) UnmarshalBinary(b []byte) error {
	var res MailersSection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
