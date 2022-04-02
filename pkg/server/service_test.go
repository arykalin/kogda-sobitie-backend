package server

import (
	"testing"
)

func TestService_Start(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestService_Start"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewService(nil)
			if err != nil {
				t.Fatal(err)
			}
			err = s.Start()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
