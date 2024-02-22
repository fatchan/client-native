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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPResponseRule HTTP Response Rule
//
// HAProxy HTTP response rule configuration (corresponds to http-response directives)
// Example: {"cond":"unless","cond_test":"{ src 192.168.0.0/16 }","hdr_format":"%T","hdr_name":"X-Haproxy-Current-Date","index":0,"type":"add-header"}
//
// swagger:model http_response_rule
type HTTPResponseRule struct {

	// return headers
	ReturnHeaders []*ReturnHeader `json:"return_hdrs,omitempty"`

	// acl file
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	ACLFile string `json:"acl_file,omitempty"`

	// acl keyfmt
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	ACLKeyfmt string `json:"acl_keyfmt,omitempty"`

	// bandwidth limit limit
	BandwidthLimitLimit string `json:"bandwidth_limit_limit,omitempty"`

	// bandwidth limit name
	BandwidthLimitName string `json:"bandwidth_limit_name,omitempty"`

	// bandwidth limit period
	BandwidthLimitPeriod string `json:"bandwidth_limit_period,omitempty"`

	// cache name
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	CacheName string `json:"cache_name,omitempty"`

	// capture id
	CaptureID *int64 `json:"capture_id,omitempty"`

	// capture sample
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	CaptureSample string `json:"capture_sample,omitempty"`

	// cond
	// Enum: [if unless]
	// +kubebuilder:validation:Enum=if;unless;
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// deny status
	// Maximum: 599
	// Minimum: 200
	// +kubebuilder:validation:Maximum=599
	// +kubebuilder:validation:Minimum=200
	DenyStatus *int64 `json:"deny_status,omitempty"`

	// expr
	Expr string `json:"expr,omitempty"`

	// hdr format
	HdrFormat string `json:"hdr_format,omitempty"`

	// hdr match
	HdrMatch string `json:"hdr_match,omitempty"`

	// hdr method
	HdrMethod string `json:"hdr_method,omitempty"`

	// hdr name
	HdrName string `json:"hdr_name,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// log level
	// Enum: [emerg alert crit err warning notice info debug silent]
	// +kubebuilder:validation:Enum=emerg;alert;crit;err;warning;notice;info;debug;silent;
	LogLevel string `json:"log_level,omitempty"`

	// lua action
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	LuaAction string `json:"lua_action,omitempty"`

	// lua params
	LuaParams string `json:"lua_params,omitempty"`

	// map file
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	MapFile string `json:"map_file,omitempty"`

	// map keyfmt
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	MapKeyfmt string `json:"map_keyfmt,omitempty"`

	// map valuefmt
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	MapValuefmt string `json:"map_valuefmt,omitempty"`

	// mark value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	// +kubebuilder:validation:Pattern=`^(0x[0-9A-Fa-f]+|[0-9]+)$`
	MarkValue string `json:"mark_value,omitempty"`

	// nice value
	// Maximum: 1024
	// Minimum: -1024
	// +kubebuilder:validation:Maximum=1024
	// +kubebuilder:validation:Minimum=-1024
	NiceValue int64 `json:"nice_value,omitempty"`

	// redir code
	// Enum: [301 302 303 307 308]
	// +kubebuilder:validation:Enum=301;302;303;307;308;
	RedirCode *int64 `json:"redir_code,omitempty"`

	// redir option
	RedirOption string `json:"redir_option,omitempty"`

	// redir type
	// Enum: [location prefix scheme]
	// +kubebuilder:validation:Enum=location;prefix;scheme;
	RedirType string `json:"redir_type,omitempty"`

	// redir value
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	RedirValue string `json:"redir_value,omitempty"`

	// return content
	ReturnContent string `json:"return_content,omitempty"`

	// return content format
	// Enum: [default-errorfiles errorfile errorfiles file lf-file string lf-string]
	// +kubebuilder:validation:Enum=default-errorfiles;errorfile;errorfiles;file;lf-file;string;lf-string;
	ReturnContentFormat string `json:"return_content_format,omitempty"`

	// return content type
	ReturnContentType *string `json:"return_content_type,omitempty"`

	// return status code
	// Maximum: 599
	// Minimum: 200
	// +kubebuilder:validation:Maximum=599
	// +kubebuilder:validation:Minimum=200
	ReturnStatusCode *int64 `json:"return_status_code,omitempty"`

	// sc expr
	ScExpr string `json:"sc_expr,omitempty"`

	// sc id
	ScID int64 `json:"sc_id,omitempty"`

	// sc idx
	ScIdx int64 `json:"sc_idx,omitempty"`

	// sc int
	ScInt *int64 `json:"sc_int,omitempty"`

	// spoe engine
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	SpoeEngine string `json:"spoe_engine,omitempty"`

	// spoe group
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	SpoeGroup string `json:"spoe_group,omitempty"`

	// status
	// Maximum: 999
	// Minimum: 100
	// +kubebuilder:validation:Maximum=999
	// +kubebuilder:validation:Minimum=100
	Status int64 `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// strict mode
	// Enum: [on off]
	// +kubebuilder:validation:Enum=on;off;
	StrictMode string `json:"strict_mode,omitempty"`

	// timeout
	Timeout string `json:"timeout,omitempty"`

	// timeout type
	// Enum: [server tunnel client]
	// +kubebuilder:validation:Enum=server;tunnel;client;
	TimeoutType string `json:"timeout_type,omitempty"`

	// tos value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	// +kubebuilder:validation:Pattern=`^(0x[0-9A-Fa-f]+|[0-9]+)$`
	TosValue string `json:"tos_value,omitempty"`

	// track sc0 key
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TrackSc0Key string `json:"track-sc0-key,omitempty"`

	// track sc0 table
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TrackSc0Table string `json:"track-sc0-table,omitempty"`

	// track sc1 key
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TrackSc1Key string `json:"track-sc1-key,omitempty"`

	// track sc1 table
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TrackSc1Table string `json:"track-sc1-table,omitempty"`

	// track sc2 key
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TrackSc2Key string `json:"track-sc2-key,omitempty"`

	// track sc2 table
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TrackSc2Table string `json:"track-sc2-table,omitempty"`

	// track sc key
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TrackScKey string `json:"track_sc_key,omitempty"`

	// track sc stick counter
	TrackScStickCounter *int64 `json:"track_sc_stick_counter,omitempty"`

	// track sc table
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TrackScTable string `json:"track_sc_table,omitempty"`

	// type
	// Required: true
	// Enum: [add-acl add-header allow cache-store capture del-acl del-header del-map deny lua redirect replace-header replace-value return sc-add-gpc sc-inc-gpc sc-inc-gpc0 sc-inc-gpc1 sc-set-gpt0 send-spoe-group set-fc-mark set-fc-tos set-header set-log-level set-map set-mark set-nice set-status set-timeout set-tos set-var set-var-fmt silent-drop strict-mode track-sc0 track-sc1 track-sc2 track-sc unset-var wait-for-body set-bandwidth-limit]
	// +kubebuilder:validation:Enum=add-acl;add-header;allow;cache-store;capture;del-acl;del-header;del-map;deny;lua;redirect;replace-header;replace-value;return;sc-add-gpc;sc-inc-gpc;sc-inc-gpc0;sc-inc-gpc1;sc-set-gpt0;send-spoe-group;set-fc-mark;set-fc-tos;set-header;set-log-level;set-map;set-mark;set-nice;set-status;set-timeout;set-tos;set-var;set-var-fmt;silent-drop;strict-mode;track-sc0;track-sc1;track-sc2;track-sc;unset-var;wait-for-body;set-bandwidth-limit;
	Type string `json:"type"`

	// var expr
	VarExpr string `json:"var_expr,omitempty"`

	// var format
	VarFormat string `json:"var_format,omitempty"`

	// var name
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	VarName string `json:"var_name,omitempty"`

	// var scope
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	VarScope string `json:"var_scope,omitempty"`

	// wait at least
	WaitAtLeast *int64 `json:"wait_at_least,omitempty"`

	// wait time
	WaitTime *int64 `json:"wait_time,omitempty"`
}

// Validate validates this http response rule
func (m *HTTPResponseRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReturnHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateACLFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateACLKeyfmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCacheName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaptureSample(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDenyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapKeyfmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapValuefmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNiceValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnContentFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrictMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeoutType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTosValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc0Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc0Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc1Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc1Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc2Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc2Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackScKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackScTable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPResponseRule) validateReturnHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnHeaders) { // not required
		return nil
	}

	for i := 0; i < len(m.ReturnHeaders); i++ {
		if swag.IsZero(m.ReturnHeaders[i]) { // not required
			continue
		}

		if m.ReturnHeaders[i] != nil {
			if err := m.ReturnHeaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("return_hdrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("return_hdrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HTTPResponseRule) validateACLFile(formats strfmt.Registry) error {
	if swag.IsZero(m.ACLFile) { // not required
		return nil
	}

	if err := validate.Pattern("acl_file", "body", m.ACLFile, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateACLKeyfmt(formats strfmt.Registry) error {
	if swag.IsZero(m.ACLKeyfmt) { // not required
		return nil
	}

	if err := validate.Pattern("acl_keyfmt", "body", m.ACLKeyfmt, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateCacheName(formats strfmt.Registry) error {
	if swag.IsZero(m.CacheName) { // not required
		return nil
	}

	if err := validate.Pattern("cache_name", "body", m.CacheName, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateCaptureSample(formats strfmt.Registry) error {
	if swag.IsZero(m.CaptureSample) { // not required
		return nil
	}

	if err := validate.Pattern("capture_sample", "body", m.CaptureSample, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeCondPropEnum = append(httpResponseRuleTypeCondPropEnum, v)
	}
}

const (

	// HTTPResponseRuleCondIf captures enum value "if"
	HTTPResponseRuleCondIf string = "if"

	// HTTPResponseRuleCondUnless captures enum value "unless"
	HTTPResponseRuleCondUnless string = "unless"
)

// prop value enum
func (m *HTTPResponseRule) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateCond(formats strfmt.Registry) error {
	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateDenyStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.DenyStatus) { // not required
		return nil
	}

	if err := validate.MinimumInt("deny_status", "body", *m.DenyStatus, 200, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("deny_status", "body", *m.DenyStatus, 599, false); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug","silent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeLogLevelPropEnum = append(httpResponseRuleTypeLogLevelPropEnum, v)
	}
}

const (

	// HTTPResponseRuleLogLevelEmerg captures enum value "emerg"
	HTTPResponseRuleLogLevelEmerg string = "emerg"

	// HTTPResponseRuleLogLevelAlert captures enum value "alert"
	HTTPResponseRuleLogLevelAlert string = "alert"

	// HTTPResponseRuleLogLevelCrit captures enum value "crit"
	HTTPResponseRuleLogLevelCrit string = "crit"

	// HTTPResponseRuleLogLevelErr captures enum value "err"
	HTTPResponseRuleLogLevelErr string = "err"

	// HTTPResponseRuleLogLevelWarning captures enum value "warning"
	HTTPResponseRuleLogLevelWarning string = "warning"

	// HTTPResponseRuleLogLevelNotice captures enum value "notice"
	HTTPResponseRuleLogLevelNotice string = "notice"

	// HTTPResponseRuleLogLevelInfo captures enum value "info"
	HTTPResponseRuleLogLevelInfo string = "info"

	// HTTPResponseRuleLogLevelDebug captures enum value "debug"
	HTTPResponseRuleLogLevelDebug string = "debug"

	// HTTPResponseRuleLogLevelSilent captures enum value "silent"
	HTTPResponseRuleLogLevelSilent string = "silent"
)

// prop value enum
func (m *HTTPResponseRule) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogLevelEnum("log_level", "body", m.LogLevel); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateLuaAction(formats strfmt.Registry) error {
	if swag.IsZero(m.LuaAction) { // not required
		return nil
	}

	if err := validate.Pattern("lua_action", "body", m.LuaAction, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapFile(formats strfmt.Registry) error {
	if swag.IsZero(m.MapFile) { // not required
		return nil
	}

	if err := validate.Pattern("map_file", "body", m.MapFile, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapKeyfmt(formats strfmt.Registry) error {
	if swag.IsZero(m.MapKeyfmt) { // not required
		return nil
	}

	if err := validate.Pattern("map_keyfmt", "body", m.MapKeyfmt, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapValuefmt(formats strfmt.Registry) error {
	if swag.IsZero(m.MapValuefmt) { // not required
		return nil
	}

	if err := validate.Pattern("map_valuefmt", "body", m.MapValuefmt, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMarkValue(formats strfmt.Registry) error {
	if swag.IsZero(m.MarkValue) { // not required
		return nil
	}

	if err := validate.Pattern("mark_value", "body", m.MarkValue, `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateNiceValue(formats strfmt.Registry) error {
	if swag.IsZero(m.NiceValue) { // not required
		return nil
	}

	if err := validate.MinimumInt("nice_value", "body", m.NiceValue, -1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nice_value", "body", m.NiceValue, 1024, false); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeRedirCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[301,302,303,307,308]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeRedirCodePropEnum = append(httpResponseRuleTypeRedirCodePropEnum, v)
	}
}

// prop value enum
func (m *HTTPResponseRule) validateRedirCodeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeRedirCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateRedirCode(formats strfmt.Registry) error {
	if swag.IsZero(m.RedirCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirCodeEnum("redir_code", "body", *m.RedirCode); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeRedirTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["location","prefix","scheme"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeRedirTypePropEnum = append(httpResponseRuleTypeRedirTypePropEnum, v)
	}
}

const (

	// HTTPResponseRuleRedirTypeLocation captures enum value "location"
	HTTPResponseRuleRedirTypeLocation string = "location"

	// HTTPResponseRuleRedirTypePrefix captures enum value "prefix"
	HTTPResponseRuleRedirTypePrefix string = "prefix"

	// HTTPResponseRuleRedirTypeScheme captures enum value "scheme"
	HTTPResponseRuleRedirTypeScheme string = "scheme"
)

// prop value enum
func (m *HTTPResponseRule) validateRedirTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeRedirTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateRedirType(formats strfmt.Registry) error {
	if swag.IsZero(m.RedirType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirTypeEnum("redir_type", "body", m.RedirType); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateRedirValue(formats strfmt.Registry) error {
	if swag.IsZero(m.RedirValue) { // not required
		return nil
	}

	if err := validate.Pattern("redir_value", "body", m.RedirValue, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeReturnContentFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default-errorfiles","errorfile","errorfiles","file","lf-file","string","lf-string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeReturnContentFormatPropEnum = append(httpResponseRuleTypeReturnContentFormatPropEnum, v)
	}
}

const (

	// HTTPResponseRuleReturnContentFormatDefaultDashErrorfiles captures enum value "default-errorfiles"
	HTTPResponseRuleReturnContentFormatDefaultDashErrorfiles string = "default-errorfiles"

	// HTTPResponseRuleReturnContentFormatErrorfile captures enum value "errorfile"
	HTTPResponseRuleReturnContentFormatErrorfile string = "errorfile"

	// HTTPResponseRuleReturnContentFormatErrorfiles captures enum value "errorfiles"
	HTTPResponseRuleReturnContentFormatErrorfiles string = "errorfiles"

	// HTTPResponseRuleReturnContentFormatFile captures enum value "file"
	HTTPResponseRuleReturnContentFormatFile string = "file"

	// HTTPResponseRuleReturnContentFormatLfDashFile captures enum value "lf-file"
	HTTPResponseRuleReturnContentFormatLfDashFile string = "lf-file"

	// HTTPResponseRuleReturnContentFormatString captures enum value "string"
	HTTPResponseRuleReturnContentFormatString string = "string"

	// HTTPResponseRuleReturnContentFormatLfDashString captures enum value "lf-string"
	HTTPResponseRuleReturnContentFormatLfDashString string = "lf-string"
)

// prop value enum
func (m *HTTPResponseRule) validateReturnContentFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeReturnContentFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateReturnContentFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnContentFormat) { // not required
		return nil
	}

	// value enum
	if err := m.validateReturnContentFormatEnum("return_content_format", "body", m.ReturnContentFormat); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateReturnStatusCode(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnStatusCode) { // not required
		return nil
	}

	if err := validate.MinimumInt("return_status_code", "body", *m.ReturnStatusCode, 200, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("return_status_code", "body", *m.ReturnStatusCode, 599, false); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateSpoeEngine(formats strfmt.Registry) error {
	if swag.IsZero(m.SpoeEngine) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_engine", "body", m.SpoeEngine, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateSpoeGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.SpoeGroup) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_group", "body", m.SpoeGroup, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.MinimumInt("status", "body", m.Status, 100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("status", "body", m.Status, 999, false); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeStrictModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeStrictModePropEnum = append(httpResponseRuleTypeStrictModePropEnum, v)
	}
}

const (

	// HTTPResponseRuleStrictModeOn captures enum value "on"
	HTTPResponseRuleStrictModeOn string = "on"

	// HTTPResponseRuleStrictModeOff captures enum value "off"
	HTTPResponseRuleStrictModeOff string = "off"
)

// prop value enum
func (m *HTTPResponseRule) validateStrictModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeStrictModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateStrictMode(formats strfmt.Registry) error {
	if swag.IsZero(m.StrictMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStrictModeEnum("strict_mode", "body", m.StrictMode); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeTimeoutTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["server","tunnel","client"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeTimeoutTypePropEnum = append(httpResponseRuleTypeTimeoutTypePropEnum, v)
	}
}

const (

	// HTTPResponseRuleTimeoutTypeServer captures enum value "server"
	HTTPResponseRuleTimeoutTypeServer string = "server"

	// HTTPResponseRuleTimeoutTypeTunnel captures enum value "tunnel"
	HTTPResponseRuleTimeoutTypeTunnel string = "tunnel"

	// HTTPResponseRuleTimeoutTypeClient captures enum value "client"
	HTTPResponseRuleTimeoutTypeClient string = "client"
)

// prop value enum
func (m *HTTPResponseRule) validateTimeoutTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeTimeoutTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateTimeoutType(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeoutType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeoutTypeEnum("timeout_type", "body", m.TimeoutType); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTosValue(formats strfmt.Registry) error {
	if swag.IsZero(m.TosValue) { // not required
		return nil
	}

	if err := validate.Pattern("tos_value", "body", m.TosValue, `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc0Key(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackSc0Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc0-key", "body", m.TrackSc0Key, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc0Table(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackSc0Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc0-table", "body", m.TrackSc0Table, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc1Key(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackSc1Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc1-key", "body", m.TrackSc1Key, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc1Table(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackSc1Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc1-table", "body", m.TrackSc1Table, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc2Key(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackSc2Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc2-key", "body", m.TrackSc2Key, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc2Table(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackSc2Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc2-table", "body", m.TrackSc2Table, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackScKey(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackScKey) { // not required
		return nil
	}

	if err := validate.Pattern("track_sc_key", "body", m.TrackScKey, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackScTable(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackScTable) { // not required
		return nil
	}

	if err := validate.Pattern("track_sc_table", "body", m.TrackScTable, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add-acl","add-header","allow","cache-store","capture","del-acl","del-header","del-map","deny","lua","redirect","replace-header","replace-value","return","sc-add-gpc","sc-inc-gpc","sc-inc-gpc0","sc-inc-gpc1","sc-set-gpt0","send-spoe-group","set-fc-mark","set-fc-tos","set-header","set-log-level","set-map","set-mark","set-nice","set-status","set-timeout","set-tos","set-var","set-var-fmt","silent-drop","strict-mode","track-sc0","track-sc1","track-sc2","track-sc","unset-var","wait-for-body","set-bandwidth-limit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeTypePropEnum = append(httpResponseRuleTypeTypePropEnum, v)
	}
}

const (

	// HTTPResponseRuleTypeAddDashACL captures enum value "add-acl"
	HTTPResponseRuleTypeAddDashACL string = "add-acl"

	// HTTPResponseRuleTypeAddDashHeader captures enum value "add-header"
	HTTPResponseRuleTypeAddDashHeader string = "add-header"

	// HTTPResponseRuleTypeAllow captures enum value "allow"
	HTTPResponseRuleTypeAllow string = "allow"

	// HTTPResponseRuleTypeCacheDashStore captures enum value "cache-store"
	HTTPResponseRuleTypeCacheDashStore string = "cache-store"

	// HTTPResponseRuleTypeCapture captures enum value "capture"
	HTTPResponseRuleTypeCapture string = "capture"

	// HTTPResponseRuleTypeDelDashACL captures enum value "del-acl"
	HTTPResponseRuleTypeDelDashACL string = "del-acl"

	// HTTPResponseRuleTypeDelDashHeader captures enum value "del-header"
	HTTPResponseRuleTypeDelDashHeader string = "del-header"

	// HTTPResponseRuleTypeDelDashMap captures enum value "del-map"
	HTTPResponseRuleTypeDelDashMap string = "del-map"

	// HTTPResponseRuleTypeDeny captures enum value "deny"
	HTTPResponseRuleTypeDeny string = "deny"

	// HTTPResponseRuleTypeLua captures enum value "lua"
	HTTPResponseRuleTypeLua string = "lua"

	// HTTPResponseRuleTypeRedirect captures enum value "redirect"
	HTTPResponseRuleTypeRedirect string = "redirect"

	// HTTPResponseRuleTypeReplaceDashHeader captures enum value "replace-header"
	HTTPResponseRuleTypeReplaceDashHeader string = "replace-header"

	// HTTPResponseRuleTypeReplaceDashValue captures enum value "replace-value"
	HTTPResponseRuleTypeReplaceDashValue string = "replace-value"

	// HTTPResponseRuleTypeReturn captures enum value "return"
	HTTPResponseRuleTypeReturn string = "return"

	// HTTPResponseRuleTypeScDashAddDashGpc captures enum value "sc-add-gpc"
	HTTPResponseRuleTypeScDashAddDashGpc string = "sc-add-gpc"

	// HTTPResponseRuleTypeScDashIncDashGpc captures enum value "sc-inc-gpc"
	HTTPResponseRuleTypeScDashIncDashGpc string = "sc-inc-gpc"

	// HTTPResponseRuleTypeScDashIncDashGpc0 captures enum value "sc-inc-gpc0"
	HTTPResponseRuleTypeScDashIncDashGpc0 string = "sc-inc-gpc0"

	// HTTPResponseRuleTypeScDashIncDashGpc1 captures enum value "sc-inc-gpc1"
	HTTPResponseRuleTypeScDashIncDashGpc1 string = "sc-inc-gpc1"

	// HTTPResponseRuleTypeScDashSetDashGpt0 captures enum value "sc-set-gpt0"
	HTTPResponseRuleTypeScDashSetDashGpt0 string = "sc-set-gpt0"

	// HTTPResponseRuleTypeSendDashSpoeDashGroup captures enum value "send-spoe-group"
	HTTPResponseRuleTypeSendDashSpoeDashGroup string = "send-spoe-group"

	// HTTPResponseRuleTypeSetDashFcDashMark captures enum value "set-fc-mark"
	HTTPResponseRuleTypeSetDashFcDashMark string = "set-fc-mark"

	// HTTPResponseRuleTypeSetDashFcDashTos captures enum value "set-fc-tos"
	HTTPResponseRuleTypeSetDashFcDashTos string = "set-fc-tos"

	// HTTPResponseRuleTypeSetDashHeader captures enum value "set-header"
	HTTPResponseRuleTypeSetDashHeader string = "set-header"

	// HTTPResponseRuleTypeSetDashLogDashLevel captures enum value "set-log-level"
	HTTPResponseRuleTypeSetDashLogDashLevel string = "set-log-level"

	// HTTPResponseRuleTypeSetDashMap captures enum value "set-map"
	HTTPResponseRuleTypeSetDashMap string = "set-map"

	// HTTPResponseRuleTypeSetDashMark captures enum value "set-mark"
	HTTPResponseRuleTypeSetDashMark string = "set-mark"

	// HTTPResponseRuleTypeSetDashNice captures enum value "set-nice"
	HTTPResponseRuleTypeSetDashNice string = "set-nice"

	// HTTPResponseRuleTypeSetDashStatus captures enum value "set-status"
	HTTPResponseRuleTypeSetDashStatus string = "set-status"

	// HTTPResponseRuleTypeSetDashTimeout captures enum value "set-timeout"
	HTTPResponseRuleTypeSetDashTimeout string = "set-timeout"

	// HTTPResponseRuleTypeSetDashTos captures enum value "set-tos"
	HTTPResponseRuleTypeSetDashTos string = "set-tos"

	// HTTPResponseRuleTypeSetDashVar captures enum value "set-var"
	HTTPResponseRuleTypeSetDashVar string = "set-var"

	// HTTPResponseRuleTypeSetDashVarDashFmt captures enum value "set-var-fmt"
	HTTPResponseRuleTypeSetDashVarDashFmt string = "set-var-fmt"

	// HTTPResponseRuleTypeSilentDashDrop captures enum value "silent-drop"
	HTTPResponseRuleTypeSilentDashDrop string = "silent-drop"

	// HTTPResponseRuleTypeStrictDashMode captures enum value "strict-mode"
	HTTPResponseRuleTypeStrictDashMode string = "strict-mode"

	// HTTPResponseRuleTypeTrackDashSc0 captures enum value "track-sc0"
	HTTPResponseRuleTypeTrackDashSc0 string = "track-sc0"

	// HTTPResponseRuleTypeTrackDashSc1 captures enum value "track-sc1"
	HTTPResponseRuleTypeTrackDashSc1 string = "track-sc1"

	// HTTPResponseRuleTypeTrackDashSc2 captures enum value "track-sc2"
	HTTPResponseRuleTypeTrackDashSc2 string = "track-sc2"

	// HTTPResponseRuleTypeTrackDashSc captures enum value "track-sc"
	HTTPResponseRuleTypeTrackDashSc string = "track-sc"

	// HTTPResponseRuleTypeUnsetDashVar captures enum value "unset-var"
	HTTPResponseRuleTypeUnsetDashVar string = "unset-var"

	// HTTPResponseRuleTypeWaitDashForDashBody captures enum value "wait-for-body"
	HTTPResponseRuleTypeWaitDashForDashBody string = "wait-for-body"

	// HTTPResponseRuleTypeSetDashBandwidthDashLimit captures enum value "set-bandwidth-limit"
	HTTPResponseRuleTypeSetDashBandwidthDashLimit string = "set-bandwidth-limit"
)

// prop value enum
func (m *HTTPResponseRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateVarName(formats strfmt.Registry) error {
	if swag.IsZero(m.VarName) { // not required
		return nil
	}

	if err := validate.Pattern("var_name", "body", m.VarName, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateVarScope(formats strfmt.Registry) error {
	if swag.IsZero(m.VarScope) { // not required
		return nil
	}

	if err := validate.Pattern("var_scope", "body", m.VarScope, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this http response rule based on the context it is used
func (m *HTTPResponseRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReturnHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPResponseRule) contextValidateReturnHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReturnHeaders); i++ {

		if m.ReturnHeaders[i] != nil {
			if err := m.ReturnHeaders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("return_hdrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("return_hdrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPResponseRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPResponseRule) UnmarshalBinary(b []byte) error {
	var res HTTPResponseRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
