/*
Copyright 2022 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"github.com/haproxytech/client-native/v6/configuration"
	"github.com/haproxytech/client-native/v6/runtime"
	"github.com/haproxytech/client-native/v6/spoe"
	"github.com/haproxytech/client-native/v6/storage"
)

type Options struct {
	Configuration  configuration.Configuration
	Runtime        runtime.Runtime
	MapStorage     storage.Storage
	SSLCertStorage storage.Storage
	GeneralStorage storage.Storage
	Spoe           spoe.Spoe
}

type Option interface {
	Set(p *Options) error
}
