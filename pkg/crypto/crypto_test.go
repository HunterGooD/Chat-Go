package crypto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateKey(t *testing.T) {
	res := GenerateKey("key", 16)
	res1 := GenerateKey("key", 3)

	require.NotEmpty(t, res)
	require.Equal(t, res, []byte("keykeykeykeykeyk"))
	require.Equal(t, len(res), 16)

	require.NotEmpty(t, res1)
	require.Equal(t, res1, []byte("key"))
	require.Equal(t, len(res1), 3)
}
