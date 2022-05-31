package utils

import (
	"fmt"
	"testing"
)

func TestGetLocalIp(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "getip"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := GetLocalIp()
			fmt.Println(ip)
		})
	}
}
