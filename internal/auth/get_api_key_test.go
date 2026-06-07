package auth

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey fa234bbaff494"},
	}

	key, err := GetAPIKey(headers)
	require.NoError(t, err)
	assert.Equal(t, key, "fa234bbaff494")
}

func TestGetAPIKey_Fail(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"Bearer fa234bbaff494"},
	}

	_, err := GetAPIKey(headers)
	require.Error(t, err)
}
