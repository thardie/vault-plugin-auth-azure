package plugin

import (
	"context"
	"net/http"
	"sync"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

const (
	authorizationBaseURI = "https://login.windows.net"
	issuerBaseURI        = "https://sts.windows.net"
)

func Factory(ctx context.Context, c *logical.BackendConfig) (logical.Backend, error) {
	b := Backend(c)
	if err := b.Setup(ctx, c); err != nil {
		return nil, err
	}
	return b, nil
}

type azureAuthBackend struct {
	*framework.Backend

	l sync.RWMutex

	provider   provider
	httpClient *http.Client
}

func Backend(c *logical.BackendConfig) *azureAuthBackend {
	b := new(azureAuthBackend)

	b.Backend = &framework.Backend{
		AuthRenew:   b.pathLoginRenew,
		BackendType: logical.TypeCredential,
		Invalidate:  b.invalidate,
		Help:        backendHelp,
		PathsSpecial: &logical.Paths{
			Unauthenticated: []string{
				"login",
			},
			SealWrapStorage: []string{
				"config",
			},
		},
		Paths: framework.PathAppend(
			[]*framework.Path{
				pathLogin(b),
				pathConfig(b),
			},
			pathsRole(b),
		),
	}
	b.httpClient = cleanhttp.DefaultClient()

	return b
}

func (b *azureAuthBackend) invalidate(ctx context.Context, key string) {
	switch key {
	case "config":
		b.reset()
	}
}

func (b *azureAuthBackend) getProvider(config *azureConfig) (provider, error) {
	b.l.RLock()
	unlockFunc := b.l.RUnlock
	defer func() { unlockFunc() }()

	if b.provider != nil {
		return b.provider, nil
	}

	// Upgrade lock
	b.l.RUnlock()
	b.l.Lock()
	unlockFunc = b.l.Unlock

	if b.provider != nil {
		return b.provider, nil
	}

	provider, err := NewAzureProvider(config)
	if err != nil {
		return nil, err
	}

	b.provider = provider
	return b.provider, nil
}

func (b *azureAuthBackend) reset() {
	b.l.Lock()
	defer b.l.Unlock()

	b.provider = nil
}

const backendHelp = `
The Azure backend plugin allows authentication for Azure .
`
