package utils

import "testing"

func TestEthAddressRegex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		addr string
		want bool
	}{
		{"valid", "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed", true},
		{"invalid", "invalid_eth_address", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEthAddress(tt.addr); got != tt.want {
				t.Errorf("ValidateEthAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
