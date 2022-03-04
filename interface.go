package clientnative

import (
	"context"

	"github.com/haproxytech/client-native/v3/configuration"
	"github.com/haproxytech/client-native/v3/options"
	"github.com/haproxytech/client-native/v3/runtime"
	"github.com/haproxytech/client-native/v3/spoe"
	"github.com/haproxytech/client-native/v3/storage"
)

type HAProxyClient interface {
	Configuration() (configuration.Configuration, error)
	Runtime() (runtime.Runtime, error)
	ReplaceConfiguration(configurationClient configuration.Configuration)
	ReplaceRuntime(runtimeClient runtime.Runtime)
	MapStorage() (storage.Storage, error)
	SSLCertStorage() (storage.Storage, error)
	GeneralStorage() (storage.Storage, error)
	Spoe() (spoe.Spoe, error)
}

func New(ctx context.Context, opt ...options.Option) (HAProxyClient, error) {
	o := options.Options{}
	var err error

	for _, option := range opt {
		err = option.Set(&o)
		if err != nil {
			return nil, err
		}
	}

	c := &haProxyClient{
		configuration:  o.Configuration,
		runtime:        o.Runtime,
		mapStorage:     o.MapStorage,
		sslCertStorage: o.SSLCertStorage,
		generalStorage: o.GeneralStorage,
		spoe:           o.Spoe,
	}

	return c, nil
}