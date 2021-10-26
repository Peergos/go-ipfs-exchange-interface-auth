// Package exchange defines the IPFS exchange interface
package exchange

import (
	"context"
	"io"

	auth "github.com/peergos/go-bitswap-auth/auth"
)

// Interface defines the functionality of the IPFS block exchange protocol.
type Interface interface { // type Exchanger interface
	Fetcher

	// TODO Should callers be concerned with whether the block was made
	// available on the network?
	HasBlock(auth.AuthBlock) error

	IsOnline() bool

	io.Closer
}

// Fetcher is an object that can be used to retrieve blocks
type Fetcher interface {
	// GetBlock returns the block associated with a given key.
	GetBlock(context.Context, auth.Want) (auth.AuthBlock, error)
	GetBlocks(context.Context, []auth.Want) (<-chan auth.AuthBlock, error)
}

// SessionExchange is an exchange.Interface which supports
// sessions.
type SessionExchange interface {
	Interface
	NewSession(context.Context) Fetcher
}
