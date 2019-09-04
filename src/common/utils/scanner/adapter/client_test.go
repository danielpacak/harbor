package adapter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetMetadata(t *testing.T) {
	client := NewClient("http://localhost:8080/api/v1")
	metadata, err := client.GetMetadata()
	require.NoError(t, err)
	t.Logf("Metadata: %v", metadata)
}
