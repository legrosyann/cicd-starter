package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"simple":                  {input: map[string][]string{"Authorization": {"ApiKey key"}}, want: "key", err: nil},
		"no header Authorization": {input: map[string][]string{"Auth": {"ApiKey key"}}, want: "", err: errors.New("no authorization header included")},
		"malformed no apiKey":     {input: map[string][]string{"Authorization": {"ApiKey"}}, want: "", err: errors.New("malformed authorization header")},
		"malformed":               {input: map[string][]string{"Authorization": {"Key key"}}, want: "", err: errors.New("malformed authorization header")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
