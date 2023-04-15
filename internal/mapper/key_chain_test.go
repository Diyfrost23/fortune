package mapper

import (
	"testing"

	"github.com/shlima/fortune/internal/pkg/domain"
	"github.com/stretchr/testify/require"
)

func TestBlockchainLink(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := BlockchainLink("foo")
		require.Equal(t, "https://www.blockchain.com/explorer/addresses/btc/foo", got)
	})
}

func TestKeyChainHTML(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		chain := domain.KeyChain{
			Private:      "foo",
			Compressed:   "bar",
			Uncompressed: "baz",
		}

		got := KeyChainHTML(chain)
		require.Contains(t, got, chain.Private)
		require.Contains(t, got, chain.Compressed)
		require.Contains(t, got, chain.Uncompressed)
	})
}
