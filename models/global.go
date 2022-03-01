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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Global Global
//
// HAProxy global configuration
//
// swagger:model global
type Global struct {

	// CPU maps
	CPUMaps []*CPUMap `json:"cpu_maps"`

	// h1 case adjusts
	H1CaseAdjusts []*H1CaseAdjust `json:"h1_case_adjust"`

	// runtime a p is
	RuntimeAPIs []*RuntimeAPI `json:"runtime_apis"`

	// ca base
	CaBase string `json:"ca_base,omitempty"`

	// chroot
	// Pattern: ^[^\s]+$
	Chroot string `json:"chroot,omitempty"`

	// crt base
	CrtBase string `json:"crt_base,omitempty"`

	// daemon
	// Enum: [enabled disabled]
	Daemon string `json:"daemon,omitempty"`

	// external check
	ExternalCheck bool `json:"external_check,omitempty"`

	// group
	// Pattern: ^[^\s]+$
	Group string `json:"group,omitempty"`

	// h1 case adjust file
	H1CaseAdjustFile string `json:"h1_case_adjust_file,omitempty"`

	// hard stop after
	HardStopAfter *int64 `json:"hard_stop_after,omitempty"`

	// localpeer
	// Pattern: ^[^\s]+$
	Localpeer string `json:"localpeer,omitempty"`

	// log send hostname
	LogSendHostname *GlobalLogSendHostname `json:"log_send_hostname,omitempty"`

	// lua loads
	LuaLoads []*LuaLoad `json:"lua_loads"`

	// lua prepend path
	LuaPrependPath []*LuaPrependPath `json:"lua_prepend_path"`

	// master worker
	MasterWorker bool `json:"master-worker,omitempty"`

	// maxconn
	Maxconn int64 `json:"maxconn,omitempty"`

	// nbproc
	Nbproc int64 `json:"nbproc,omitempty"`

	// nbthread
	Nbthread int64 `json:"nbthread,omitempty"`

	// pidfile
	Pidfile string `json:"pidfile,omitempty"`

	// server state base
	// Pattern: ^[^\s]+$
	ServerStateBase string `json:"server_state_base,omitempty"`

	// server state file
	// Pattern: ^[^\s]+$
	ServerStateFile string `json:"server_state_file,omitempty"`

	// ssl default bind ciphers
	SslDefaultBindCiphers string `json:"ssl_default_bind_ciphers,omitempty"`

	// ssl default bind ciphersuites
	SslDefaultBindCiphersuites string `json:"ssl_default_bind_ciphersuites,omitempty"`

	// ssl default bind options
	SslDefaultBindOptions string `json:"ssl_default_bind_options,omitempty"`

	// ssl default server ciphers
	SslDefaultServerCiphers string `json:"ssl_default_server_ciphers,omitempty"`

	// ssl default server ciphersuites
	SslDefaultServerCiphersuites string `json:"ssl_default_server_ciphersuites,omitempty"`

	// ssl default server options
	SslDefaultServerOptions string `json:"ssl_default_server_options,omitempty"`

	// ssl mode async
	// Enum: [enabled disabled]
	SslModeAsync string `json:"ssl_mode_async,omitempty"`

	// stats timeout
	StatsTimeout *int64 `json:"stats_timeout,omitempty"`

	// tune options
	TuneOptions *GlobalTuneOptions `json:"tune_options,omitempty"`

	// tune ssl default dh param
	TuneSslDefaultDhParam int64 `json:"tune_ssl_default_dh_param,omitempty"`

	// user
	// Pattern: ^[^\s]+$
	User string `json:"user,omitempty"`
}

// Validate validates this global
func (m *Global) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUMaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateH1CaseAdjusts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimeAPIs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChroot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaemon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalpeer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogSendHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaLoads(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaPrependPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerStateBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerStateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslModeAsync(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTuneOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Global) validateCPUMaps(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUMaps) { // not required
		return nil
	}

	for i := 0; i < len(m.CPUMaps); i++ {
		if swag.IsZero(m.CPUMaps[i]) { // not required
			continue
		}

		if m.CPUMaps[i] != nil {
			if err := m.CPUMaps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpu_maps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateH1CaseAdjusts(formats strfmt.Registry) error {

	if swag.IsZero(m.H1CaseAdjusts) { // not required
		return nil
	}

	for i := 0; i < len(m.H1CaseAdjusts); i++ {
		if swag.IsZero(m.H1CaseAdjusts[i]) { // not required
			continue
		}

		if m.H1CaseAdjusts[i] != nil {
			if err := m.H1CaseAdjusts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("h1_case_adjust" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateRuntimeAPIs(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeAPIs) { // not required
		return nil
	}

	for i := 0; i < len(m.RuntimeAPIs); i++ {
		if swag.IsZero(m.RuntimeAPIs[i]) { // not required
			continue
		}

		if m.RuntimeAPIs[i] != nil {
			if err := m.RuntimeAPIs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runtime_apis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateChroot(formats strfmt.Registry) error {

	if swag.IsZero(m.Chroot) { // not required
		return nil
	}

	if err := validate.Pattern("chroot", "body", string(m.Chroot), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var globalTypeDaemonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTypeDaemonPropEnum = append(globalTypeDaemonPropEnum, v)
	}
}

const (

	// GlobalDaemonEnabled captures enum value "enabled"
	GlobalDaemonEnabled string = "enabled"

	// GlobalDaemonDisabled captures enum value "disabled"
	GlobalDaemonDisabled string = "disabled"
)

// prop value enum
func (m *Global) validateDaemonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalTypeDaemonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Global) validateDaemon(formats strfmt.Registry) error {

	if swag.IsZero(m.Daemon) { // not required
		return nil
	}

	// value enum
	if err := m.validateDaemonEnum("daemon", "body", m.Daemon); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if err := validate.Pattern("group", "body", string(m.Group), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateLocalpeer(formats strfmt.Registry) error {

	if swag.IsZero(m.Localpeer) { // not required
		return nil
	}

	if err := validate.Pattern("localpeer", "body", string(m.Localpeer), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateLogSendHostname(formats strfmt.Registry) error {

	if swag.IsZero(m.LogSendHostname) { // not required
		return nil
	}

	if m.LogSendHostname != nil {
		if err := m.LogSendHostname.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_send_hostname")
			}
			return err
		}
	}

	return nil
}

func (m *Global) validateLuaLoads(formats strfmt.Registry) error {

	if swag.IsZero(m.LuaLoads) { // not required
		return nil
	}

	for i := 0; i < len(m.LuaLoads); i++ {
		if swag.IsZero(m.LuaLoads[i]) { // not required
			continue
		}

		if m.LuaLoads[i] != nil {
			if err := m.LuaLoads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lua_loads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateLuaPrependPath(formats strfmt.Registry) error {

	if swag.IsZero(m.LuaPrependPath) { // not required
		return nil
	}

	for i := 0; i < len(m.LuaPrependPath); i++ {
		if swag.IsZero(m.LuaPrependPath[i]) { // not required
			continue
		}

		if m.LuaPrependPath[i] != nil {
			if err := m.LuaPrependPath[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lua_prepend_path" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateServerStateBase(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerStateBase) { // not required
		return nil
	}

	if err := validate.Pattern("server_state_base", "body", string(m.ServerStateBase), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateServerStateFile(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerStateFile) { // not required
		return nil
	}

	if err := validate.Pattern("server_state_file", "body", string(m.ServerStateFile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var globalTypeSslModeAsyncPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTypeSslModeAsyncPropEnum = append(globalTypeSslModeAsyncPropEnum, v)
	}
}

const (

	// GlobalSslModeAsyncEnabled captures enum value "enabled"
	GlobalSslModeAsyncEnabled string = "enabled"

	// GlobalSslModeAsyncDisabled captures enum value "disabled"
	GlobalSslModeAsyncDisabled string = "disabled"
)

// prop value enum
func (m *Global) validateSslModeAsyncEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalTypeSslModeAsyncPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Global) validateSslModeAsync(formats strfmt.Registry) error {

	if swag.IsZero(m.SslModeAsync) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslModeAsyncEnum("ssl_mode_async", "body", m.SslModeAsync); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateTuneOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.TuneOptions) { // not required
		return nil
	}

	if m.TuneOptions != nil {
		if err := m.TuneOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tune_options")
			}
			return err
		}
	}

	return nil
}

func (m *Global) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := validate.Pattern("user", "body", string(m.User), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Global) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Global) UnmarshalBinary(b []byte) error {
	var res Global
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CPUMap CPU map
//
// swagger:model CPUMap
type CPUMap struct {

	// cpu set
	// Required: true
	CPUSet *string `json:"cpu_set"`

	// process
	// Required: true
	Process *string `json:"process"`
}

// Validate validates this CPU map
func (m *CPUMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUMap) validateCPUSet(formats strfmt.Registry) error {

	if err := validate.Required("cpu_set", "body", m.CPUSet); err != nil {
		return err
	}

	return nil
}

func (m *CPUMap) validateProcess(formats strfmt.Registry) error {

	if err := validate.Required("process", "body", m.Process); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CPUMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUMap) UnmarshalBinary(b []byte) error {
	var res CPUMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// H1CaseAdjust h1 case adjust
//
// swagger:model H1CaseAdjust
type H1CaseAdjust struct {

	// from
	// Required: true
	From *string `json:"from"`

	// to
	// Required: true
	To *string `json:"to"`
}

// Validate validates this h1 case adjust
func (m *H1CaseAdjust) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *H1CaseAdjust) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *H1CaseAdjust) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *H1CaseAdjust) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *H1CaseAdjust) UnmarshalBinary(b []byte) error {
	var res H1CaseAdjust
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GlobalLogSendHostname global log send hostname
//
// swagger:model GlobalLogSendHostname
type GlobalLogSendHostname struct {

	// enabled
	// Required: true
	// Enum: [enabled disabled]
	Enabled *string `json:"enabled"`

	// param
	// Pattern: ^[^\s]+$
	Param string `json:"param,omitempty"`
}

// Validate validates this global log send hostname
func (m *GlobalLogSendHostname) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var globalLogSendHostnameTypeEnabledPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalLogSendHostnameTypeEnabledPropEnum = append(globalLogSendHostnameTypeEnabledPropEnum, v)
	}
}

const (

	// GlobalLogSendHostnameEnabledEnabled captures enum value "enabled"
	GlobalLogSendHostnameEnabledEnabled string = "enabled"

	// GlobalLogSendHostnameEnabledDisabled captures enum value "disabled"
	GlobalLogSendHostnameEnabledDisabled string = "disabled"
)

// prop value enum
func (m *GlobalLogSendHostname) validateEnabledEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalLogSendHostnameTypeEnabledPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GlobalLogSendHostname) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("log_send_hostname"+"."+"enabled", "body", m.Enabled); err != nil {
		return err
	}

	// value enum
	if err := m.validateEnabledEnum("log_send_hostname"+"."+"enabled", "body", *m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *GlobalLogSendHostname) validateParam(formats strfmt.Registry) error {

	if swag.IsZero(m.Param) { // not required
		return nil
	}

	if err := validate.Pattern("log_send_hostname"+"."+"param", "body", string(m.Param), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalLogSendHostname) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalLogSendHostname) UnmarshalBinary(b []byte) error {
	var res GlobalLogSendHostname
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LuaLoad lua load
//
// swagger:model LuaLoad
type LuaLoad struct {

	// file
	// Required: true
	// Pattern: ^[^\s]+$
	File *string `json:"file"`
}

// Validate validates this lua load
func (m *LuaLoad) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LuaLoad) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", m.File); err != nil {
		return err
	}

	if err := validate.Pattern("file", "body", string(*m.File), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LuaLoad) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LuaLoad) UnmarshalBinary(b []byte) error {
	var res LuaLoad
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LuaPrependPath lua prepend path
//
// swagger:model LuaPrependPath
type LuaPrependPath struct {

	// path
	// Required: true
	// Pattern: ^[^\s]+$
	Path *string `json:"path"`

	// type
	// Enum: [path cpath]
	Type string `json:"type,omitempty"`
}

// Validate validates this lua prepend path
func (m *LuaPrependPath) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LuaPrependPath) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	if err := validate.Pattern("path", "body", string(*m.Path), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var luaPrependPathTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["path","cpath"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		luaPrependPathTypeTypePropEnum = append(luaPrependPathTypeTypePropEnum, v)
	}
}

const (

	// LuaPrependPathTypePath captures enum value "path"
	LuaPrependPathTypePath string = "path"

	// LuaPrependPathTypeCpath captures enum value "cpath"
	LuaPrependPathTypeCpath string = "cpath"
)

// prop value enum
func (m *LuaPrependPath) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, luaPrependPathTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LuaPrependPath) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LuaPrependPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LuaPrependPath) UnmarshalBinary(b []byte) error {
	var res LuaPrependPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RuntimeAPI runtime API
//
// swagger:model RuntimeAPI
type RuntimeAPI struct {
	BindParams

	// address
	// Required: true
	// Pattern: ^[^\s]+$
	Address *string `json:"address"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RuntimeAPI) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BindParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BindParams = aO0

	// now for regular properties
	var propsRuntimeAPI struct {
		Address *string `json:"address"`
	}
	if err := swag.ReadJSON(raw, &propsRuntimeAPI); err != nil {
		return err
	}
	m.Address = propsRuntimeAPI.Address

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RuntimeAPI) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BindParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsRuntimeAPI struct {
		Address *string `json:"address"`
	}
	propsRuntimeAPI.Address = m.Address

	jsonDataPropsRuntimeAPI, errRuntimeAPI := swag.WriteJSON(propsRuntimeAPI)
	if errRuntimeAPI != nil {
		return nil, errRuntimeAPI
	}
	_parts = append(_parts, jsonDataPropsRuntimeAPI)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this runtime API
func (m *RuntimeAPI) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BindParams
	if err := m.BindParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuntimeAPI) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if err := validate.Pattern("address", "body", string(*m.Address), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeAPI) UnmarshalBinary(b []byte) error {
	var res RuntimeAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GlobalTuneOptions global tune options
//
// swagger:model GlobalTuneOptions
type GlobalTuneOptions struct {

	// buffers limit
	BuffersLimit *int64 `json:"buffers_limit,omitempty"`

	// buffers reserve
	// Minimum: 2
	BuffersReserve int64 `json:"buffers_reserve,omitempty"`

	// bufsize
	Bufsize int64 `json:"bufsize,omitempty"`

	// comp maxlevel
	CompMaxlevel int64 `json:"comp_maxlevel,omitempty"`

	// fail alloc
	FailAlloc bool `json:"fail_alloc,omitempty"`

	// h2 header table size
	// Maximum: 65535
	H2HeaderTableSize int64 `json:"h2_header_table_size,omitempty"`

	// h2 initial window size
	H2InitialWindowSize *int64 `json:"h2_initial_window_size,omitempty"`

	// h2 max concurrent streams
	H2MaxConcurrentStreams int64 `json:"h2_max_concurrent_streams,omitempty"`

	// h2 max frame size
	H2MaxFrameSize int64 `json:"h2_max_frame_size,omitempty"`

	// http cookielen
	HTTPCookielen int64 `json:"http_cookielen,omitempty"`

	// http logurilen
	HTTPLogurilen int64 `json:"http_logurilen,omitempty"`

	// http maxhdr
	// Maximum: 32767
	// Minimum: 1
	HTTPMaxhdr int64 `json:"http_maxhdr,omitempty"`

	// idle pool shared
	// Enum: [enabled disabled]
	IdlePoolShared string `json:"idle_pool_shared,omitempty"`

	// idletimer
	// Maximum: 65535
	// Minimum: 0
	Idletimer *int64 `json:"idletimer,omitempty"`

	// listener multi queue
	// Enum: [enabled disabled]
	ListenerMultiQueue string `json:"listener_multi_queue,omitempty"`

	// lua forced yield
	LuaForcedYield int64 `json:"lua_forced_yield,omitempty"`

	// lua maxmem
	LuaMaxmem bool `json:"lua_maxmem,omitempty"`

	// lua service timeout
	LuaServiceTimeout *int64 `json:"lua_service_timeout,omitempty"`

	// lua session timeout
	LuaSessionTimeout *int64 `json:"lua_session_timeout,omitempty"`

	// lua task timeout
	LuaTaskTimeout *int64 `json:"lua_task_timeout,omitempty"`

	// maxaccept
	Maxaccept int64 `json:"maxaccept,omitempty"`

	// maxpollevents
	Maxpollevents int64 `json:"maxpollevents,omitempty"`

	// maxrewrite
	Maxrewrite int64 `json:"maxrewrite,omitempty"`

	// pattern cache size
	PatternCacheSize *int64 `json:"pattern_cache_size,omitempty"`

	// pipesize
	Pipesize int64 `json:"pipesize,omitempty"`

	// pool high fd ratio
	PoolHighFdRatio int64 `json:"pool_high_fd_ratio,omitempty"`

	// pool low fd ratio
	PoolLowFdRatio int64 `json:"pool_low_fd_ratio,omitempty"`

	// rcvbuf client
	RcvbufClient *int64 `json:"rcvbuf_client,omitempty"`

	// rcvbuf server
	RcvbufServer *int64 `json:"rcvbuf_server,omitempty"`

	// recv enough
	RecvEnough int64 `json:"recv_enough,omitempty"`

	// runqueue depth
	RunqueueDepth int64 `json:"runqueue_depth,omitempty"`

	// sched low latency
	// Enum: [enabled disabled]
	SchedLowLatency string `json:"sched_low_latency,omitempty"`

	// sndbuf client
	SndbufClient *int64 `json:"sndbuf_client,omitempty"`

	// sndbuf server
	SndbufServer *int64 `json:"sndbuf_server,omitempty"`

	// ssl cachesize
	SslCachesize *int64 `json:"ssl_cachesize,omitempty"`

	// ssl capture buffer size
	SslCaptureBufferSize *int64 `json:"ssl_capture_buffer_size,omitempty"`

	// ssl ctx cache size
	SslCtxCacheSize int64 `json:"ssl_ctx_cache_size,omitempty"`

	// ssl default dh param
	SslDefaultDhParam int64 `json:"ssl_default_dh_param,omitempty"`

	// ssl force private cache
	SslForcePrivateCache bool `json:"ssl_force_private_cache,omitempty"`

	// ssl keylog
	// Enum: [enabled disabled]
	SslKeylog string `json:"ssl_keylog,omitempty"`

	// ssl lifetime
	SslLifetime *int64 `json:"ssl_lifetime,omitempty"`

	// ssl maxrecord
	SslMaxrecord *int64 `json:"ssl_maxrecord,omitempty"`

	// vars global max size
	VarsGlobalMaxSize *int64 `json:"vars_global_max_size,omitempty"`

	// vars proc max size
	VarsProcMaxSize *int64 `json:"vars_proc_max_size,omitempty"`

	// vars reqres max size
	VarsReqresMaxSize *int64 `json:"vars_reqres_max_size,omitempty"`

	// vars sess max size
	VarsSessMaxSize *int64 `json:"vars_sess_max_size,omitempty"`

	// vars txn max size
	VarsTxnMaxSize *int64 `json:"vars_txn_max_size,omitempty"`

	// zlib memlevel
	// Maximum: 9
	// Minimum: 1
	ZlibMemlevel int64 `json:"zlib_memlevel,omitempty"`

	// zlib windowsize
	// Maximum: 15
	// Minimum: 8
	ZlibWindowsize int64 `json:"zlib_windowsize,omitempty"`
}

// Validate validates this global tune options
func (m *GlobalTuneOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuffersReserve(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateH2HeaderTableSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPMaxhdr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdlePoolShared(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdletimer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListenerMultiQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedLowLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslKeylog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZlibMemlevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZlibWindowsize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalTuneOptions) validateBuffersReserve(formats strfmt.Registry) error {

	if swag.IsZero(m.BuffersReserve) { // not required
		return nil
	}

	if err := validate.MinimumInt("tune_options"+"."+"buffers_reserve", "body", int64(m.BuffersReserve), 2, false); err != nil {
		return err
	}

	return nil
}

func (m *GlobalTuneOptions) validateH2HeaderTableSize(formats strfmt.Registry) error {

	if swag.IsZero(m.H2HeaderTableSize) { // not required
		return nil
	}

	if err := validate.MaximumInt("tune_options"+"."+"h2_header_table_size", "body", int64(m.H2HeaderTableSize), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *GlobalTuneOptions) validateHTTPMaxhdr(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPMaxhdr) { // not required
		return nil
	}

	if err := validate.MinimumInt("tune_options"+"."+"http_maxhdr", "body", int64(m.HTTPMaxhdr), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("tune_options"+"."+"http_maxhdr", "body", int64(m.HTTPMaxhdr), 32767, false); err != nil {
		return err
	}

	return nil
}

var globalTuneOptionsTypeIdlePoolSharedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTuneOptionsTypeIdlePoolSharedPropEnum = append(globalTuneOptionsTypeIdlePoolSharedPropEnum, v)
	}
}

const (

	// GlobalTuneOptionsIdlePoolSharedEnabled captures enum value "enabled"
	GlobalTuneOptionsIdlePoolSharedEnabled string = "enabled"

	// GlobalTuneOptionsIdlePoolSharedDisabled captures enum value "disabled"
	GlobalTuneOptionsIdlePoolSharedDisabled string = "disabled"
)

// prop value enum
func (m *GlobalTuneOptions) validateIdlePoolSharedEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalTuneOptionsTypeIdlePoolSharedPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GlobalTuneOptions) validateIdlePoolShared(formats strfmt.Registry) error {

	if swag.IsZero(m.IdlePoolShared) { // not required
		return nil
	}

	// value enum
	if err := m.validateIdlePoolSharedEnum("tune_options"+"."+"idle_pool_shared", "body", m.IdlePoolShared); err != nil {
		return err
	}

	return nil
}

func (m *GlobalTuneOptions) validateIdletimer(formats strfmt.Registry) error {

	if swag.IsZero(m.Idletimer) { // not required
		return nil
	}

	if err := validate.MinimumInt("tune_options"+"."+"idletimer", "body", int64(*m.Idletimer), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("tune_options"+"."+"idletimer", "body", int64(*m.Idletimer), 65535, false); err != nil {
		return err
	}

	return nil
}

var globalTuneOptionsTypeListenerMultiQueuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTuneOptionsTypeListenerMultiQueuePropEnum = append(globalTuneOptionsTypeListenerMultiQueuePropEnum, v)
	}
}

const (

	// GlobalTuneOptionsListenerMultiQueueEnabled captures enum value "enabled"
	GlobalTuneOptionsListenerMultiQueueEnabled string = "enabled"

	// GlobalTuneOptionsListenerMultiQueueDisabled captures enum value "disabled"
	GlobalTuneOptionsListenerMultiQueueDisabled string = "disabled"
)

// prop value enum
func (m *GlobalTuneOptions) validateListenerMultiQueueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalTuneOptionsTypeListenerMultiQueuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GlobalTuneOptions) validateListenerMultiQueue(formats strfmt.Registry) error {

	if swag.IsZero(m.ListenerMultiQueue) { // not required
		return nil
	}

	// value enum
	if err := m.validateListenerMultiQueueEnum("tune_options"+"."+"listener_multi_queue", "body", m.ListenerMultiQueue); err != nil {
		return err
	}

	return nil
}

var globalTuneOptionsTypeSchedLowLatencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTuneOptionsTypeSchedLowLatencyPropEnum = append(globalTuneOptionsTypeSchedLowLatencyPropEnum, v)
	}
}

const (

	// GlobalTuneOptionsSchedLowLatencyEnabled captures enum value "enabled"
	GlobalTuneOptionsSchedLowLatencyEnabled string = "enabled"

	// GlobalTuneOptionsSchedLowLatencyDisabled captures enum value "disabled"
	GlobalTuneOptionsSchedLowLatencyDisabled string = "disabled"
)

// prop value enum
func (m *GlobalTuneOptions) validateSchedLowLatencyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalTuneOptionsTypeSchedLowLatencyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GlobalTuneOptions) validateSchedLowLatency(formats strfmt.Registry) error {

	if swag.IsZero(m.SchedLowLatency) { // not required
		return nil
	}

	// value enum
	if err := m.validateSchedLowLatencyEnum("tune_options"+"."+"sched_low_latency", "body", m.SchedLowLatency); err != nil {
		return err
	}

	return nil
}

var globalTuneOptionsTypeSslKeylogPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTuneOptionsTypeSslKeylogPropEnum = append(globalTuneOptionsTypeSslKeylogPropEnum, v)
	}
}

const (

	// GlobalTuneOptionsSslKeylogEnabled captures enum value "enabled"
	GlobalTuneOptionsSslKeylogEnabled string = "enabled"

	// GlobalTuneOptionsSslKeylogDisabled captures enum value "disabled"
	GlobalTuneOptionsSslKeylogDisabled string = "disabled"
)

// prop value enum
func (m *GlobalTuneOptions) validateSslKeylogEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalTuneOptionsTypeSslKeylogPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GlobalTuneOptions) validateSslKeylog(formats strfmt.Registry) error {

	if swag.IsZero(m.SslKeylog) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslKeylogEnum("tune_options"+"."+"ssl_keylog", "body", m.SslKeylog); err != nil {
		return err
	}

	return nil
}

func (m *GlobalTuneOptions) validateZlibMemlevel(formats strfmt.Registry) error {

	if swag.IsZero(m.ZlibMemlevel) { // not required
		return nil
	}

	if err := validate.MinimumInt("tune_options"+"."+"zlib_memlevel", "body", int64(m.ZlibMemlevel), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("tune_options"+"."+"zlib_memlevel", "body", int64(m.ZlibMemlevel), 9, false); err != nil {
		return err
	}

	return nil
}

func (m *GlobalTuneOptions) validateZlibWindowsize(formats strfmt.Registry) error {

	if swag.IsZero(m.ZlibWindowsize) { // not required
		return nil
	}

	if err := validate.MinimumInt("tune_options"+"."+"zlib_windowsize", "body", int64(m.ZlibWindowsize), 8, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("tune_options"+"."+"zlib_windowsize", "body", int64(m.ZlibWindowsize), 15, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalTuneOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalTuneOptions) UnmarshalBinary(b []byte) error {
	var res GlobalTuneOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
