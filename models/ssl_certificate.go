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

// SslCertificate SSL File
//
// A file containing one or more SSL/TLS certificates and keys
//
// swagger:model ssl_certificate
type SslCertificate struct {

	// algorithm
	Algorithm string `json:"algorithm,omitempty"`

	// authority key id
	AuthorityKeyID string `json:"authority_key_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// domains
	// Read Only: true
	Domains string `json:"domains,omitempty"`

	// file
	File string `json:"file,omitempty"`

	// ip addresses
	// Read Only: true
	IPAddresses string `json:"ip_addresses,omitempty"`

	// issuers
	// Read Only: true
	Issuers string `json:"issuers,omitempty"`

	// not after
	// Read Only: true
	// Format: date-time
	// +kubebuilder:validation:Format=date-time
	NotAfter *strfmt.DateTime `json:"not_after,omitempty" gorm:"type:timestamp with time zone"`

	// not before
	// Read Only: true
	// Format: date-time
	// +kubebuilder:validation:Format=date-time
	NotBefore *strfmt.DateTime `json:"not_before,omitempty" gorm:"type:timestamp with time zone"`

	// serial
	Serial string `json:"serial,omitempty"`

	// sha1 finger print
	Sha1FingerPrint string `json:"sha1_finger_print,omitempty"`

	// sha256 finger print
	Sha256FingerPrint string `json:"sha256_finger_print,omitempty"`

	// File size in bytes.
	// Read Only: true
	Size *int64 `json:"size,omitempty"`

	// storage name
	StorageName string `json:"storage_name,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// subject alternative names
	SubjectAlternativeNames string `json:"subject_alternative_names,omitempty"`

	// subject key id
	SubjectKeyID string `json:"subject_key_id,omitempty"`
}

// Validate validates this ssl certificate
func (m *SslCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotBefore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SslCertificate) validateNotAfter(formats strfmt.Registry) error {
	if swag.IsZero(m.NotAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("not_after", "body", "date-time", m.NotAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SslCertificate) validateNotBefore(formats strfmt.Registry) error {
	if swag.IsZero(m.NotBefore) { // not required
		return nil
	}

	if err := validate.FormatOf("not_before", "body", "date-time", m.NotBefore.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ssl certificate based on the context it is used
func (m *SslCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssuers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotAfter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotBefore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SslCertificate) contextValidateDomains(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "domains", "body", string(m.Domains)); err != nil {
		return err
	}

	return nil
}

func (m *SslCertificate) contextValidateIPAddresses(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ip_addresses", "body", string(m.IPAddresses)); err != nil {
		return err
	}

	return nil
}

func (m *SslCertificate) contextValidateIssuers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "issuers", "body", string(m.Issuers)); err != nil {
		return err
	}

	return nil
}

func (m *SslCertificate) contextValidateNotAfter(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "not_after", "body", m.NotAfter); err != nil {
		return err
	}

	return nil
}

func (m *SslCertificate) contextValidateNotBefore(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "not_before", "body", m.NotBefore); err != nil {
		return err
	}

	return nil
}

func (m *SslCertificate) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SslCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SslCertificate) UnmarshalBinary(b []byte) error {
	var res SslCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
