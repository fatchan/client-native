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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NativeStatsCollection Stats collection
//
// Stats from one runtime API
//
// swagger:model native_stats_collection
type NativeStatsCollection struct {

	// error
	Error string `json:"error,omitempty"`

	// runtime API
	RuntimeAPI string `json:"runtimeAPI,omitempty"`

	// stats
	Stats []*NativeStat `json:"stats"`
}

// Validate validates this native stats collection
func (m *NativeStatsCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NativeStatsCollection) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	for i := 0; i < len(m.Stats); i++ {
		if swag.IsZero(m.Stats[i]) { // not required
			continue
		}

		if m.Stats[i] != nil {
			if err := m.Stats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this native stats collection based on the context it is used
func (m *NativeStatsCollection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NativeStatsCollection) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stats); i++ {

		if m.Stats[i] != nil {
			if err := m.Stats[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NativeStatsCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NativeStatsCollection) UnmarshalBinary(b []byte) error {
	var res NativeStatsCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
