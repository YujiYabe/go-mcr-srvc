package domain

import "testing"

func TestValidationWordRuleTypeFromBlacklistFlag(t *testing.T) {
	tests := []struct {
		name        string
		isBlacklist bool
		want        string
	}{
		{
			name:        "blacklist",
			isBlacklist: true,
			want:        ValidationWordRuleTypeBlacklist,
		},
		{
			name:        "whitelist",
			isBlacklist: false,
			want:        ValidationWordRuleTypeWhitelist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidationWordRuleTypeFromBlacklistFlag(tt.isBlacklist)
			if got != tt.want {
				t.Fatalf("expected %s, got %s", tt.want, got)
			}
		})
	}
}
