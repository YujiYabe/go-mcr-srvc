package domain

const (
	ValidationWordRuleTypeWhitelist     = "whitelist"
	ValidationWordRuleTypeBlacklist     = "blacklist"
	ValidationWordRuleMatchTypeContains = "contains"
)

func ValidationWordRuleTypeFromBlacklistFlag(
	isBlacklist bool,
) (
	value string,
) {
	if isBlacklist {
		return ValidationWordRuleTypeBlacklist
	}

	return ValidationWordRuleTypeWhitelist
}
