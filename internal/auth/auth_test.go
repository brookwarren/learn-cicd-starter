package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name    string
		input   http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "Valid API Key",
			input:   http.Header{"Authorization": []string{"ApiKey 1234567890"}},
			wantKey: "1234567890",
			wantErr: false,
		},
		{
			name:    "Missing Header",
			input:   http.Header{},
			wantErr: true,
		},
		{
			name:    "Invalid Format",
			input:   http.Header{"Authorization": []string{"Basic abcdefgh"}},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if gotKey != tc.wantKey {
				t.Errorf("GetAPIKey() = %v, want %v", gotKey, tc.wantKey)
			}
		})
	}
}
