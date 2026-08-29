package domain

const (
	ValidationWordRuleTypeWhitelist = "whitelist"
	ValidationWordRuleTypeBlacklist = "blacklist"
)

func ValidationWordRuleTypeFromBlacklistFlag(isBlacklist bool) string {
	if isBlacklist {
		return ValidationWordRuleTypeBlacklist
	}

	return ValidationWordRuleTypeWhitelist
}
