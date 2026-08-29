package domain

const (
	ValidationWordRuleTypeWhitelist     = "whitelist"
	ValidationWordRuleTypeBlacklist     = "blacklist"
	ValidationWordRuleMatchTypeContains = "contains"
)

func ValidationWordRuleTypeFromBlacklistFlag(isBlacklist bool) string {
	if isBlacklist {
		return ValidationWordRuleTypeBlacklist
	}

	return ValidationWordRuleTypeWhitelist
}
