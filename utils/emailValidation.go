package utils

import (
	"github.com/wbrijesh/utils/helpers"
	"github.com/wbrijesh/utils/types"
)

func ValidateEmail(email string, rules types.EmailValidationRules) (validation bool, message string) {
	var trustedProviders []string = ["as", "as"]

	if rules.ValidEmailFormatCheck {
		if !helpers.ContainsAny(email, "@") || !helpers.ContainsAny(email, ".") {
			return false, "Invalid email format"
		}
		if helpers.IndexOf(email, '@') > helpers.IndexOf(email, '.') {
			return false, "@ should come before ."
		}
		if helpers.IndexOf(email, '@') == 0 {
			return false, "@ should not be the first character"
		}
		if helpers.IndexOf(email, '.') == len(email)-1 {
			return false, ". should not be the last character"
		}
	}
	return true, "No checks applied"
}
