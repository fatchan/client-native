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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TCPRequestRule TCP Request Rule
//
// HAProxy TCP Request Rule configuration (corresponds to tcp-request)
// Example: {"cond":"if","cond_test":"{ src 192.168.0.0/16 }","index":0,"type":"connection"}
//
// swagger:model tcp_request_rule
type TCPRequestRule struct {
	// action
	// Enum: [accept attach-srv capture do-resolve expect-netscaler-cip expect-proxy lua reject sc-add-gpc sc-inc-gpc sc-inc-gpc0 sc-inc-gpc1 sc-set-gpt sc-set-gpt0 send-spoe-group set-bandwidth-limit set-bc-mark set-bc-tos set-dst-port set-dst set-fc-mark set-fc-tos set-log-level set-mark set-nice set-priority-class set-priority-offset set-src set-src-port set-tos set-var set-var-fmt silent-drop switch-mode track-sc0 track-sc1 track-sc2 track-sc unset-var use-service]
	// +kubebuilder:validation:Enum=accept;attach-srv;capture;do-resolve;expect-netscaler-cip;expect-proxy;lua;reject;sc-add-gpc;sc-inc-gpc;sc-inc-gpc0;sc-inc-gpc1;sc-set-gpt;sc-set-gpt0;send-spoe-group;set-bandwidth-limit;set-bc-mark;set-bc-tos;set-dst-port;set-dst;set-fc-mark;set-fc-tos;set-log-level;set-mark;set-nice;set-priority-class;set-priority-offset;set-src;set-src-port;set-tos;set-var;set-var-fmt;silent-drop;switch-mode;track-sc0;track-sc1;track-sc2;track-sc;unset-var;use-service;
	Action string `json:"action,omitempty"`

	// bandwidth limit limit
	BandwidthLimitLimit string `json:"bandwidth_limit_limit,omitempty"`

	// bandwidth limit name
	BandwidthLimitName string `json:"bandwidth_limit_name,omitempty"`

	// bandwidth limit period
	BandwidthLimitPeriod string `json:"bandwidth_limit_period,omitempty"`

	// capture len
	CaptureLen int64 `json:"capture_len,omitempty"`

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

	// expr
	Expr string `json:"expr,omitempty"`

	// gpt value
	GptValue string `json:"gpt_value,omitempty"`

	// index
	// Required: true
	// +kubebuilder:validation:Optional
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

	// resolve protocol
	// Enum: [ipv4 ipv6]
	// +kubebuilder:validation:Enum=ipv4;ipv6;
	ResolveProtocol string `json:"resolve_protocol,omitempty"`

	// resolve resolvers
	ResolveResolvers string `json:"resolve_resolvers,omitempty"`

	// resolve var
	ResolveVar string `json:"resolve_var,omitempty"`

	// sc idx
	ScIdx string `json:"sc_idx,omitempty"`

	// sc inc id
	ScIncID string `json:"sc_inc_id,omitempty"`

	// sc int
	ScInt *int64 `json:"sc_int,omitempty"`

	// server name
	ServerName string `json:"server_name,omitempty"`

	// service name
	ServiceName string `json:"service_name,omitempty"`

	// spoe engine name
	SpoeEngineName string `json:"spoe_engine_name,omitempty"`

	// spoe group name
	SpoeGroupName string `json:"spoe_group_name,omitempty"`

	// switch mode proto
	SwitchModeProto string `json:"switch_mode_proto,omitempty"`

	// timeout
	Timeout *int64 `json:"timeout,omitempty"`

	// tos value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	// +kubebuilder:validation:Pattern=`^(0x[0-9A-Fa-f]+|[0-9]+)$`
	TosValue string `json:"tos_value,omitempty"`

	// track key
	TrackKey string `json:"track_key,omitempty"`

	// track stick counter
	TrackStickCounter *int64 `json:"track_stick_counter,omitempty"`

	// track table
	TrackTable string `json:"track_table,omitempty"`

	// type
	// Required: true
	// Enum: [connection content inspect-delay session]
	// +kubebuilder:validation:Enum=connection;content;inspect-delay;session;
	Type string `json:"type"`

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
}

// Validate validates this tcp request rule
func (m *TCPRequestRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaptureSample(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCond(formats); err != nil {
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

	if err := m.validateMarkValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNiceValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolveProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTosValue(formats); err != nil {
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

var tcpRequestRuleTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["accept","attach-srv","capture","do-resolve","expect-netscaler-cip","expect-proxy","lua","reject","sc-add-gpc","sc-inc-gpc","sc-inc-gpc0","sc-inc-gpc1","sc-set-gpt","sc-set-gpt0","send-spoe-group","set-bandwidth-limit","set-bc-mark","set-bc-tos","set-dst-port","set-dst","set-fc-mark","set-fc-tos","set-log-level","set-mark","set-nice","set-priority-class","set-priority-offset","set-src","set-src-port","set-tos","set-var","set-var-fmt","silent-drop","switch-mode","track-sc0","track-sc1","track-sc2","track-sc","unset-var","use-service"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpRequestRuleTypeActionPropEnum = append(tcpRequestRuleTypeActionPropEnum, v)
	}
}

const (

	// TCPRequestRuleActionAccept captures enum value "accept"
	TCPRequestRuleActionAccept string = "accept"

	// TCPRequestRuleActionAttachDashSrv captures enum value "attach-srv"
	TCPRequestRuleActionAttachDashSrv string = "attach-srv"

	// TCPRequestRuleActionCapture captures enum value "capture"
	TCPRequestRuleActionCapture string = "capture"

	// TCPRequestRuleActionDoDashResolve captures enum value "do-resolve"
	TCPRequestRuleActionDoDashResolve string = "do-resolve"

	// TCPRequestRuleActionExpectDashNetscalerDashCip captures enum value "expect-netscaler-cip"
	TCPRequestRuleActionExpectDashNetscalerDashCip string = "expect-netscaler-cip"

	// TCPRequestRuleActionExpectDashProxy captures enum value "expect-proxy"
	TCPRequestRuleActionExpectDashProxy string = "expect-proxy"

	// TCPRequestRuleActionLua captures enum value "lua"
	TCPRequestRuleActionLua string = "lua"

	// TCPRequestRuleActionReject captures enum value "reject"
	TCPRequestRuleActionReject string = "reject"

	// TCPRequestRuleActionScDashAddDashGpc captures enum value "sc-add-gpc"
	TCPRequestRuleActionScDashAddDashGpc string = "sc-add-gpc"

	// TCPRequestRuleActionScDashIncDashGpc captures enum value "sc-inc-gpc"
	TCPRequestRuleActionScDashIncDashGpc string = "sc-inc-gpc"

	// TCPRequestRuleActionScDashIncDashGpc0 captures enum value "sc-inc-gpc0"
	TCPRequestRuleActionScDashIncDashGpc0 string = "sc-inc-gpc0"

	// TCPRequestRuleActionScDashIncDashGpc1 captures enum value "sc-inc-gpc1"
	TCPRequestRuleActionScDashIncDashGpc1 string = "sc-inc-gpc1"

	// TCPRequestRuleActionScDashSetDashGpt captures enum value "sc-set-gpt"
	TCPRequestRuleActionScDashSetDashGpt string = "sc-set-gpt"

	// TCPRequestRuleActionScDashSetDashGpt0 captures enum value "sc-set-gpt0"
	TCPRequestRuleActionScDashSetDashGpt0 string = "sc-set-gpt0"

	// TCPRequestRuleActionSendDashSpoeDashGroup captures enum value "send-spoe-group"
	TCPRequestRuleActionSendDashSpoeDashGroup string = "send-spoe-group"

	// TCPRequestRuleActionSetDashBandwidthDashLimit captures enum value "set-bandwidth-limit"
	TCPRequestRuleActionSetDashBandwidthDashLimit string = "set-bandwidth-limit"

	// TCPRequestRuleActionSetDashBcDashMark captures enum value "set-bc-mark"
	TCPRequestRuleActionSetDashBcDashMark string = "set-bc-mark"

	// TCPRequestRuleActionSetDashBcDashTos captures enum value "set-bc-tos"
	TCPRequestRuleActionSetDashBcDashTos string = "set-bc-tos"

	// TCPRequestRuleActionSetDashDstDashPort captures enum value "set-dst-port"
	TCPRequestRuleActionSetDashDstDashPort string = "set-dst-port"

	// TCPRequestRuleActionSetDashDst captures enum value "set-dst"
	TCPRequestRuleActionSetDashDst string = "set-dst"

	// TCPRequestRuleActionSetDashFcDashMark captures enum value "set-fc-mark"
	TCPRequestRuleActionSetDashFcDashMark string = "set-fc-mark"

	// TCPRequestRuleActionSetDashFcDashTos captures enum value "set-fc-tos"
	TCPRequestRuleActionSetDashFcDashTos string = "set-fc-tos"

	// TCPRequestRuleActionSetDashLogDashLevel captures enum value "set-log-level"
	TCPRequestRuleActionSetDashLogDashLevel string = "set-log-level"

	// TCPRequestRuleActionSetDashMark captures enum value "set-mark"
	TCPRequestRuleActionSetDashMark string = "set-mark"

	// TCPRequestRuleActionSetDashNice captures enum value "set-nice"
	TCPRequestRuleActionSetDashNice string = "set-nice"

	// TCPRequestRuleActionSetDashPriorityDashClass captures enum value "set-priority-class"
	TCPRequestRuleActionSetDashPriorityDashClass string = "set-priority-class"

	// TCPRequestRuleActionSetDashPriorityDashOffset captures enum value "set-priority-offset"
	TCPRequestRuleActionSetDashPriorityDashOffset string = "set-priority-offset"

	// TCPRequestRuleActionSetDashSrc captures enum value "set-src"
	TCPRequestRuleActionSetDashSrc string = "set-src"

	// TCPRequestRuleActionSetDashSrcDashPort captures enum value "set-src-port"
	TCPRequestRuleActionSetDashSrcDashPort string = "set-src-port"

	// TCPRequestRuleActionSetDashTos captures enum value "set-tos"
	TCPRequestRuleActionSetDashTos string = "set-tos"

	// TCPRequestRuleActionSetDashVar captures enum value "set-var"
	TCPRequestRuleActionSetDashVar string = "set-var"

	// TCPRequestRuleActionSetDashVarDashFmt captures enum value "set-var-fmt"
	TCPRequestRuleActionSetDashVarDashFmt string = "set-var-fmt"

	// TCPRequestRuleActionSilentDashDrop captures enum value "silent-drop"
	TCPRequestRuleActionSilentDashDrop string = "silent-drop"

	// TCPRequestRuleActionSwitchDashMode captures enum value "switch-mode"
	TCPRequestRuleActionSwitchDashMode string = "switch-mode"

	// TCPRequestRuleActionTrackDashSc0 captures enum value "track-sc0"
	TCPRequestRuleActionTrackDashSc0 string = "track-sc0"

	// TCPRequestRuleActionTrackDashSc1 captures enum value "track-sc1"
	TCPRequestRuleActionTrackDashSc1 string = "track-sc1"

	// TCPRequestRuleActionTrackDashSc2 captures enum value "track-sc2"
	TCPRequestRuleActionTrackDashSc2 string = "track-sc2"

	// TCPRequestRuleActionTrackDashSc captures enum value "track-sc"
	TCPRequestRuleActionTrackDashSc string = "track-sc"

	// TCPRequestRuleActionUnsetDashVar captures enum value "unset-var"
	TCPRequestRuleActionUnsetDashVar string = "unset-var"

	// TCPRequestRuleActionUseDashService captures enum value "use-service"
	TCPRequestRuleActionUseDashService string = "use-service"
)

// prop value enum
func (m *TCPRequestRule) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tcpRequestRuleTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TCPRequestRule) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *TCPRequestRule) validateCaptureSample(formats strfmt.Registry) error {
	if swag.IsZero(m.CaptureSample) { // not required
		return nil
	}

	if err := validate.Pattern("capture_sample", "body", m.CaptureSample, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var tcpRequestRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpRequestRuleTypeCondPropEnum = append(tcpRequestRuleTypeCondPropEnum, v)
	}
}

const (

	// TCPRequestRuleCondIf captures enum value "if"
	TCPRequestRuleCondIf string = "if"

	// TCPRequestRuleCondUnless captures enum value "unless"
	TCPRequestRuleCondUnless string = "unless"
)

// prop value enum
func (m *TCPRequestRule) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tcpRequestRuleTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TCPRequestRule) validateCond(formats strfmt.Registry) error {
	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *TCPRequestRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var tcpRequestRuleTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug","silent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpRequestRuleTypeLogLevelPropEnum = append(tcpRequestRuleTypeLogLevelPropEnum, v)
	}
}

const (

	// TCPRequestRuleLogLevelEmerg captures enum value "emerg"
	TCPRequestRuleLogLevelEmerg string = "emerg"

	// TCPRequestRuleLogLevelAlert captures enum value "alert"
	TCPRequestRuleLogLevelAlert string = "alert"

	// TCPRequestRuleLogLevelCrit captures enum value "crit"
	TCPRequestRuleLogLevelCrit string = "crit"

	// TCPRequestRuleLogLevelErr captures enum value "err"
	TCPRequestRuleLogLevelErr string = "err"

	// TCPRequestRuleLogLevelWarning captures enum value "warning"
	TCPRequestRuleLogLevelWarning string = "warning"

	// TCPRequestRuleLogLevelNotice captures enum value "notice"
	TCPRequestRuleLogLevelNotice string = "notice"

	// TCPRequestRuleLogLevelInfo captures enum value "info"
	TCPRequestRuleLogLevelInfo string = "info"

	// TCPRequestRuleLogLevelDebug captures enum value "debug"
	TCPRequestRuleLogLevelDebug string = "debug"

	// TCPRequestRuleLogLevelSilent captures enum value "silent"
	TCPRequestRuleLogLevelSilent string = "silent"
)

// prop value enum
func (m *TCPRequestRule) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tcpRequestRuleTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TCPRequestRule) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogLevelEnum("log_level", "body", m.LogLevel); err != nil {
		return err
	}

	return nil
}

func (m *TCPRequestRule) validateLuaAction(formats strfmt.Registry) error {
	if swag.IsZero(m.LuaAction) { // not required
		return nil
	}

	if err := validate.Pattern("lua_action", "body", m.LuaAction, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *TCPRequestRule) validateMarkValue(formats strfmt.Registry) error {
	if swag.IsZero(m.MarkValue) { // not required
		return nil
	}

	if err := validate.Pattern("mark_value", "body", m.MarkValue, `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *TCPRequestRule) validateNiceValue(formats strfmt.Registry) error {
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

var tcpRequestRuleTypeResolveProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpRequestRuleTypeResolveProtocolPropEnum = append(tcpRequestRuleTypeResolveProtocolPropEnum, v)
	}
}

const (

	// TCPRequestRuleResolveProtocolIPV4 captures enum value "ipv4"
	TCPRequestRuleResolveProtocolIPV4 string = "ipv4"

	// TCPRequestRuleResolveProtocolIPV6 captures enum value "ipv6"
	TCPRequestRuleResolveProtocolIPV6 string = "ipv6"
)

// prop value enum
func (m *TCPRequestRule) validateResolveProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tcpRequestRuleTypeResolveProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TCPRequestRule) validateResolveProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolveProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateResolveProtocolEnum("resolve_protocol", "body", m.ResolveProtocol); err != nil {
		return err
	}

	return nil
}

func (m *TCPRequestRule) validateTosValue(formats strfmt.Registry) error {
	if swag.IsZero(m.TosValue) { // not required
		return nil
	}

	if err := validate.Pattern("tos_value", "body", m.TosValue, `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

var tcpRequestRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connection","content","inspect-delay","session"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpRequestRuleTypeTypePropEnum = append(tcpRequestRuleTypeTypePropEnum, v)
	}
}

const (

	// TCPRequestRuleTypeConnection captures enum value "connection"
	TCPRequestRuleTypeConnection string = "connection"

	// TCPRequestRuleTypeContent captures enum value "content"
	TCPRequestRuleTypeContent string = "content"

	// TCPRequestRuleTypeInspectDashDelay captures enum value "inspect-delay"
	TCPRequestRuleTypeInspectDashDelay string = "inspect-delay"

	// TCPRequestRuleTypeSession captures enum value "session"
	TCPRequestRuleTypeSession string = "session"
)

// prop value enum
func (m *TCPRequestRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tcpRequestRuleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TCPRequestRule) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *TCPRequestRule) validateVarName(formats strfmt.Registry) error {
	if swag.IsZero(m.VarName) { // not required
		return nil
	}

	if err := validate.Pattern("var_name", "body", m.VarName, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *TCPRequestRule) validateVarScope(formats strfmt.Registry) error {
	if swag.IsZero(m.VarScope) { // not required
		return nil
	}

	if err := validate.Pattern("var_scope", "body", m.VarScope, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tcp request rule based on context it is used
func (m *TCPRequestRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TCPRequestRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TCPRequestRule) UnmarshalBinary(b []byte) error {
	var res TCPRequestRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
