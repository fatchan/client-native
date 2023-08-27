// Code generated with struct_equal_generator; DO NOT EDIT.

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

import (
	"strconv"
)

// Equal checks if two structs of type Frontend are equal
//
// By default empty maps and slices are equal to nil:
//
//	var a, b Frontend
//	equal := a.Equal(b)
//
// For more advanced use case you can configure these options (default values are shown):
//
//	var a, b Frontend
//	equal := a.Equal(b,Options{
//		NilSameAsEmpty: true,
//	})
func (s Frontend) Equal(t Frontend, opts ...Options) bool {
	opt := getOptions(opts...)

	if !CheckSameNilAndLen(s.ErrorFiles, t.ErrorFiles, opt) {
		return false
	}
	for i := range s.ErrorFiles {
		if !s.ErrorFiles[i].Equal(*t.ErrorFiles[i], opt) {
			return false
		}
	}

	if !CheckSameNilAndLen(s.ErrorFilesFromHTTPErrors, t.ErrorFilesFromHTTPErrors, opt) {
		return false
	}
	for i := range s.ErrorFilesFromHTTPErrors {
		if !s.ErrorFilesFromHTTPErrors[i].Equal(*t.ErrorFilesFromHTTPErrors[i], opt) {
			return false
		}
	}

	if s.AcceptInvalidHTTPRequest != t.AcceptInvalidHTTPRequest {
		return false
	}

	if !equalPointers(s.Backlog, t.Backlog) {
		return false
	}

	if s.BindProcess != t.BindProcess {
		return false
	}

	if s.Clflog != t.Clflog {
		return false
	}

	if !equalPointers(s.ClientFinTimeout, t.ClientFinTimeout) {
		return false
	}

	if !equalPointers(s.ClientTimeout, t.ClientTimeout) {
		return false
	}

	if s.Clitcpka != t.Clitcpka {
		return false
	}

	if !equalPointers(s.ClitcpkaCnt, t.ClitcpkaCnt) {
		return false
	}

	if !equalPointers(s.ClitcpkaIdle, t.ClitcpkaIdle) {
		return false
	}

	if !equalPointers(s.ClitcpkaIntvl, t.ClitcpkaIntvl) {
		return false
	}

	if !s.Compression.Equal(*t.Compression, opt) {
		return false
	}

	if s.Contstats != t.Contstats {
		return false
	}

	if s.DefaultBackend != t.DefaultBackend {
		return false
	}

	if s.Description != t.Description {
		return false
	}

	if s.DisableH2Upgrade != t.DisableH2Upgrade {
		return false
	}

	if s.Disabled != t.Disabled {
		return false
	}

	if s.DontlogNormal != t.DontlogNormal {
		return false
	}

	if s.Dontlognull != t.Dontlognull {
		return false
	}

	if !s.EmailAlert.Equal(*t.EmailAlert, opt) {
		return false
	}

	if s.Enabled != t.Enabled {
		return false
	}

	if s.ErrorLogFormat != t.ErrorLogFormat {
		return false
	}

	if !s.Errorloc302.Equal(*t.Errorloc302, opt) {
		return false
	}

	if !s.Errorloc303.Equal(*t.Errorloc303, opt) {
		return false
	}

	if !s.Forwardfor.Equal(*t.Forwardfor, opt) {
		return false
	}

	if s.From != t.From {
		return false
	}

	if s.H1CaseAdjustBogusClient != t.H1CaseAdjustBogusClient {
		return false
	}

	if s.HTTPBufferRequest != t.HTTPBufferRequest {
		return false
	}

	if s.HTTPUseHtx != t.HTTPUseHtx {
		return false
	}

	if s.HTTPConnectionMode != t.HTTPConnectionMode {
		return false
	}

	if s.HTTPIgnoreProbes != t.HTTPIgnoreProbes {
		return false
	}

	if !equalPointers(s.HTTPKeepAliveTimeout, t.HTTPKeepAliveTimeout) {
		return false
	}

	if s.HTTPNoDelay != t.HTTPNoDelay {
		return false
	}

	if !equalPointers(s.HTTPRequestTimeout, t.HTTPRequestTimeout) {
		return false
	}

	if s.HTTPRestrictReqHdrNames != t.HTTPRestrictReqHdrNames {
		return false
	}

	if s.HTTPUseProxyHeader != t.HTTPUseProxyHeader {
		return false
	}

	if s.Httplog != t.Httplog {
		return false
	}

	if s.Httpslog != t.Httpslog {
		return false
	}

	if !equalPointers(s.ID, t.ID) {
		return false
	}

	if s.IdleCloseOnResponse != t.IdleCloseOnResponse {
		return false
	}

	if s.IndependentStreams != t.IndependentStreams {
		return false
	}

	if s.LogFormat != t.LogFormat {
		return false
	}

	if s.LogFormatSd != t.LogFormatSd {
		return false
	}

	if s.LogSeparateErrors != t.LogSeparateErrors {
		return false
	}

	if s.LogTag != t.LogTag {
		return false
	}

	if s.Logasap != t.Logasap {
		return false
	}

	if !equalPointers(s.Maxconn, t.Maxconn) {
		return false
	}

	if s.Mode != t.Mode {
		return false
	}

	if !s.MonitorFail.Equal(*t.MonitorFail, opt) {
		return false
	}

	if !s.MonitorURI.Equal(t.MonitorURI, opt) {
		return false
	}

	if s.Name != t.Name {
		return false
	}

	if s.Nolinger != t.Nolinger {
		return false
	}

	if !s.Originalto.Equal(*t.Originalto, opt) {
		return false
	}

	if s.SocketStats != t.SocketStats {
		return false
	}

	if s.SpliceAuto != t.SpliceAuto {
		return false
	}

	if s.SpliceRequest != t.SpliceRequest {
		return false
	}

	if s.SpliceResponse != t.SpliceResponse {
		return false
	}

	if !s.StatsOptions.Equal(*t.StatsOptions, opt) {
		return false
	}

	if !s.StickTable.Equal(*t.StickTable, opt) {
		return false
	}

	if !equalPointers(s.TarpitTimeout, t.TarpitTimeout) {
		return false
	}

	if s.TCPSmartAccept != t.TCPSmartAccept {
		return false
	}

	if s.Tcpka != t.Tcpka {
		return false
	}

	if s.Tcplog != t.Tcplog {
		return false
	}

	if s.UniqueIDFormat != t.UniqueIDFormat {
		return false
	}

	if s.UniqueIDHeader != t.UniqueIDHeader {
		return false
	}

	return true
}

// Diff checks if two structs of type Frontend are equal
//
// By default empty arrays, maps and slices are equal to nil:
//
//	var a, b Frontend
//	diff := a.Diff(b)
//
// For more advanced use case you can configure the options (default values are shown):
//
//	var a, b Frontend
//	equal := a.Diff(b,Options{
//		NilSameAsEmpty: true,
//	})
func (s Frontend) Diff(t Frontend, opts ...Options) map[string][]interface{} {
	opt := getOptions(opts...)

	diff := make(map[string][]interface{})
	if len(s.ErrorFiles) != len(t.ErrorFiles) {
		diff["ErrorFiles"] = []interface{}{s.ErrorFiles, t.ErrorFiles}
	} else {
		diff2 := make(map[string][]interface{})
		for i := range s.ErrorFiles {
			diffSub := s.ErrorFiles[i].Diff(*t.ErrorFiles[i], opt)
			if len(diffSub) > 0 {
				diff2[strconv.Itoa(i)] = []interface{}{diffSub}
			}
		}
		if len(diff2) > 0 {
			diff["ErrorFiles"] = []interface{}{diff2}
		}
	}

	if len(s.ErrorFilesFromHTTPErrors) != len(t.ErrorFilesFromHTTPErrors) {
		diff["ErrorFilesFromHTTPErrors"] = []interface{}{s.ErrorFilesFromHTTPErrors, t.ErrorFilesFromHTTPErrors}
	} else {
		diff2 := make(map[string][]interface{})
		for i := range s.ErrorFilesFromHTTPErrors {
			diffSub := s.ErrorFilesFromHTTPErrors[i].Diff(*t.ErrorFilesFromHTTPErrors[i], opt)
			if len(diffSub) > 0 {
				diff2[strconv.Itoa(i)] = []interface{}{diffSub}
			}
		}
		if len(diff2) > 0 {
			diff["ErrorFilesFromHTTPErrors"] = []interface{}{diff2}
		}
	}

	if s.AcceptInvalidHTTPRequest != t.AcceptInvalidHTTPRequest {
		diff["AcceptInvalidHTTPRequest"] = []interface{}{s.AcceptInvalidHTTPRequest, t.AcceptInvalidHTTPRequest}
	}

	if !equalPointers(s.Backlog, t.Backlog) {
		diff["Backlog"] = []interface{}{s.Backlog, t.Backlog}
	}

	if s.BindProcess != t.BindProcess {
		diff["BindProcess"] = []interface{}{s.BindProcess, t.BindProcess}
	}

	if s.Clflog != t.Clflog {
		diff["Clflog"] = []interface{}{s.Clflog, t.Clflog}
	}

	if !equalPointers(s.ClientFinTimeout, t.ClientFinTimeout) {
		diff["ClientFinTimeout"] = []interface{}{s.ClientFinTimeout, t.ClientFinTimeout}
	}

	if !equalPointers(s.ClientTimeout, t.ClientTimeout) {
		diff["ClientTimeout"] = []interface{}{s.ClientTimeout, t.ClientTimeout}
	}

	if s.Clitcpka != t.Clitcpka {
		diff["Clitcpka"] = []interface{}{s.Clitcpka, t.Clitcpka}
	}

	if !equalPointers(s.ClitcpkaCnt, t.ClitcpkaCnt) {
		diff["ClitcpkaCnt"] = []interface{}{s.ClitcpkaCnt, t.ClitcpkaCnt}
	}

	if !equalPointers(s.ClitcpkaIdle, t.ClitcpkaIdle) {
		diff["ClitcpkaIdle"] = []interface{}{s.ClitcpkaIdle, t.ClitcpkaIdle}
	}

	if !equalPointers(s.ClitcpkaIntvl, t.ClitcpkaIntvl) {
		diff["ClitcpkaIntvl"] = []interface{}{s.ClitcpkaIntvl, t.ClitcpkaIntvl}
	}

	if !s.Compression.Equal(*t.Compression, opt) {
		diff["Compression"] = []interface{}{s.Compression, t.Compression}
	}

	if s.Contstats != t.Contstats {
		diff["Contstats"] = []interface{}{s.Contstats, t.Contstats}
	}

	if s.DefaultBackend != t.DefaultBackend {
		diff["DefaultBackend"] = []interface{}{s.DefaultBackend, t.DefaultBackend}
	}

	if s.Description != t.Description {
		diff["Description"] = []interface{}{s.Description, t.Description}
	}

	if s.DisableH2Upgrade != t.DisableH2Upgrade {
		diff["DisableH2Upgrade"] = []interface{}{s.DisableH2Upgrade, t.DisableH2Upgrade}
	}

	if s.Disabled != t.Disabled {
		diff["Disabled"] = []interface{}{s.Disabled, t.Disabled}
	}

	if s.DontlogNormal != t.DontlogNormal {
		diff["DontlogNormal"] = []interface{}{s.DontlogNormal, t.DontlogNormal}
	}

	if s.Dontlognull != t.Dontlognull {
		diff["Dontlognull"] = []interface{}{s.Dontlognull, t.Dontlognull}
	}

	if !s.EmailAlert.Equal(*t.EmailAlert, opt) {
		diff["EmailAlert"] = []interface{}{s.EmailAlert, t.EmailAlert}
	}

	if s.Enabled != t.Enabled {
		diff["Enabled"] = []interface{}{s.Enabled, t.Enabled}
	}

	if s.ErrorLogFormat != t.ErrorLogFormat {
		diff["ErrorLogFormat"] = []interface{}{s.ErrorLogFormat, t.ErrorLogFormat}
	}

	if !s.Errorloc302.Equal(*t.Errorloc302, opt) {
		diff["Errorloc302"] = []interface{}{s.Errorloc302, t.Errorloc302}
	}

	if !s.Errorloc303.Equal(*t.Errorloc303, opt) {
		diff["Errorloc303"] = []interface{}{s.Errorloc303, t.Errorloc303}
	}

	if !s.Forwardfor.Equal(*t.Forwardfor, opt) {
		diff["Forwardfor"] = []interface{}{s.Forwardfor, t.Forwardfor}
	}

	if s.From != t.From {
		diff["From"] = []interface{}{s.From, t.From}
	}

	if s.H1CaseAdjustBogusClient != t.H1CaseAdjustBogusClient {
		diff["H1CaseAdjustBogusClient"] = []interface{}{s.H1CaseAdjustBogusClient, t.H1CaseAdjustBogusClient}
	}

	if s.HTTPBufferRequest != t.HTTPBufferRequest {
		diff["HTTPBufferRequest"] = []interface{}{s.HTTPBufferRequest, t.HTTPBufferRequest}
	}

	if s.HTTPUseHtx != t.HTTPUseHtx {
		diff["HTTPUseHtx"] = []interface{}{s.HTTPUseHtx, t.HTTPUseHtx}
	}

	if s.HTTPConnectionMode != t.HTTPConnectionMode {
		diff["HTTPConnectionMode"] = []interface{}{s.HTTPConnectionMode, t.HTTPConnectionMode}
	}

	if s.HTTPIgnoreProbes != t.HTTPIgnoreProbes {
		diff["HTTPIgnoreProbes"] = []interface{}{s.HTTPIgnoreProbes, t.HTTPIgnoreProbes}
	}

	if !equalPointers(s.HTTPKeepAliveTimeout, t.HTTPKeepAliveTimeout) {
		diff["HTTPKeepAliveTimeout"] = []interface{}{s.HTTPKeepAliveTimeout, t.HTTPKeepAliveTimeout}
	}

	if s.HTTPNoDelay != t.HTTPNoDelay {
		diff["HTTPNoDelay"] = []interface{}{s.HTTPNoDelay, t.HTTPNoDelay}
	}

	if !equalPointers(s.HTTPRequestTimeout, t.HTTPRequestTimeout) {
		diff["HTTPRequestTimeout"] = []interface{}{s.HTTPRequestTimeout, t.HTTPRequestTimeout}
	}

	if s.HTTPRestrictReqHdrNames != t.HTTPRestrictReqHdrNames {
		diff["HTTPRestrictReqHdrNames"] = []interface{}{s.HTTPRestrictReqHdrNames, t.HTTPRestrictReqHdrNames}
	}

	if s.HTTPUseProxyHeader != t.HTTPUseProxyHeader {
		diff["HTTPUseProxyHeader"] = []interface{}{s.HTTPUseProxyHeader, t.HTTPUseProxyHeader}
	}

	if s.Httplog != t.Httplog {
		diff["Httplog"] = []interface{}{s.Httplog, t.Httplog}
	}

	if s.Httpslog != t.Httpslog {
		diff["Httpslog"] = []interface{}{s.Httpslog, t.Httpslog}
	}

	if !equalPointers(s.ID, t.ID) {
		diff["ID"] = []interface{}{s.ID, t.ID}
	}

	if s.IdleCloseOnResponse != t.IdleCloseOnResponse {
		diff["IdleCloseOnResponse"] = []interface{}{s.IdleCloseOnResponse, t.IdleCloseOnResponse}
	}

	if s.IndependentStreams != t.IndependentStreams {
		diff["IndependentStreams"] = []interface{}{s.IndependentStreams, t.IndependentStreams}
	}

	if s.LogFormat != t.LogFormat {
		diff["LogFormat"] = []interface{}{s.LogFormat, t.LogFormat}
	}

	if s.LogFormatSd != t.LogFormatSd {
		diff["LogFormatSd"] = []interface{}{s.LogFormatSd, t.LogFormatSd}
	}

	if s.LogSeparateErrors != t.LogSeparateErrors {
		diff["LogSeparateErrors"] = []interface{}{s.LogSeparateErrors, t.LogSeparateErrors}
	}

	if s.LogTag != t.LogTag {
		diff["LogTag"] = []interface{}{s.LogTag, t.LogTag}
	}

	if s.Logasap != t.Logasap {
		diff["Logasap"] = []interface{}{s.Logasap, t.Logasap}
	}

	if !equalPointers(s.Maxconn, t.Maxconn) {
		diff["Maxconn"] = []interface{}{s.Maxconn, t.Maxconn}
	}

	if s.Mode != t.Mode {
		diff["Mode"] = []interface{}{s.Mode, t.Mode}
	}

	if !s.MonitorFail.Equal(*t.MonitorFail, opt) {
		diff["MonitorFail"] = []interface{}{s.MonitorFail, t.MonitorFail}
	}

	if !s.MonitorURI.Equal(t.MonitorURI, opt) {
		diff["MonitorURI"] = []interface{}{s.MonitorURI, t.MonitorURI}
	}

	if s.Name != t.Name {
		diff["Name"] = []interface{}{s.Name, t.Name}
	}

	if s.Nolinger != t.Nolinger {
		diff["Nolinger"] = []interface{}{s.Nolinger, t.Nolinger}
	}

	if !s.Originalto.Equal(*t.Originalto, opt) {
		diff["Originalto"] = []interface{}{s.Originalto, t.Originalto}
	}

	if s.SocketStats != t.SocketStats {
		diff["SocketStats"] = []interface{}{s.SocketStats, t.SocketStats}
	}

	if s.SpliceAuto != t.SpliceAuto {
		diff["SpliceAuto"] = []interface{}{s.SpliceAuto, t.SpliceAuto}
	}

	if s.SpliceRequest != t.SpliceRequest {
		diff["SpliceRequest"] = []interface{}{s.SpliceRequest, t.SpliceRequest}
	}

	if s.SpliceResponse != t.SpliceResponse {
		diff["SpliceResponse"] = []interface{}{s.SpliceResponse, t.SpliceResponse}
	}

	if !s.StatsOptions.Equal(*t.StatsOptions, opt) {
		diff["StatsOptions"] = []interface{}{s.StatsOptions, t.StatsOptions}
	}

	if !s.StickTable.Equal(*t.StickTable, opt) {
		diff["StickTable"] = []interface{}{s.StickTable, t.StickTable}
	}

	if !equalPointers(s.TarpitTimeout, t.TarpitTimeout) {
		diff["TarpitTimeout"] = []interface{}{s.TarpitTimeout, t.TarpitTimeout}
	}

	if s.TCPSmartAccept != t.TCPSmartAccept {
		diff["TCPSmartAccept"] = []interface{}{s.TCPSmartAccept, t.TCPSmartAccept}
	}

	if s.Tcpka != t.Tcpka {
		diff["Tcpka"] = []interface{}{s.Tcpka, t.Tcpka}
	}

	if s.Tcplog != t.Tcplog {
		diff["Tcplog"] = []interface{}{s.Tcplog, t.Tcplog}
	}

	if s.UniqueIDFormat != t.UniqueIDFormat {
		diff["UniqueIDFormat"] = []interface{}{s.UniqueIDFormat, t.UniqueIDFormat}
	}

	if s.UniqueIDHeader != t.UniqueIDHeader {
		diff["UniqueIDHeader"] = []interface{}{s.UniqueIDHeader, t.UniqueIDHeader}
	}

	return diff
}
