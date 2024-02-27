package utils

import (
	"github.com/wbrijesh/utils/helpers"
	"github.com/wbrijesh/utils/types"
)

func ValidatePassword(password string, rules types.PasswordRules) (validation bool, message string) {
	MinLengthCheckResult := false
	MaxLengthCheckResult := false
	HasUpperCaseCheckResult := false
	HasLowerCaseCheckResult := false
	HasNumberCheckResult := false
	HasSpecialCharCheckResult := false

	if rules.MinLengthCheck {
		if len(password) < rules.MinLength {
			MinLengthCheckResult = false
		}
	} else {
		MinLengthCheckResult = true
	}

	if rules.MaxLengthCheck {
		if len(password) > rules.MaxLength {
			MaxLengthCheckResult = false
		}
	} else {
		MaxLengthCheckResult = true
	}

	if rules.HasUpperCaseCheck {
		upperCaseChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		if !helpers.ContainsAny(password, upperCaseChars) {
			HasUpperCaseCheckResult = false
		}
		HasUpperCaseCheckResult = true
	} else {
		HasUpperCaseCheckResult = true
	}

	if rules.HasLowerCaseCheck {
		lowerCaseChars := "abcdefghijklmnopqrstuvwxyz"
		if !helpers.ContainsAny(password, lowerCaseChars) {
			HasLowerCaseCheckResult = false
		}
		HasLowerCaseCheckResult = true
	} else {
		HasLowerCaseCheckResult = true
	}

	if rules.HasNumberCheck {
		numbers := "0123456789"
		if !helpers.ContainsAny(password, numbers) {
			HasNumberCheckResult = false
		}
		HasNumberCheckResult = true
	} else {
		HasNumberCheckResult = true
	}

	if rules.HasSpecialCharCheck {
		specialChars := "!@#$%^&*()_+{}|:<>?~"
		if !helpers.ContainsAny(password, specialChars) {
			HasSpecialCharCheckResult = false
		}
		HasSpecialCharCheckResult = true
	} else {
		HasSpecialCharCheckResult = true
	}

	if MinLengthCheckResult && MaxLengthCheckResult && HasUpperCaseCheckResult && HasLowerCaseCheckResult && HasNumberCheckResult && HasSpecialCharCheckResult {
		return true, "password matches all rules"
	} else {
		missingRules := ""
		if !MinLengthCheckResult {
			missingRules += "MinLengthCheck, "
		} else if !MaxLengthCheckResult {
			missingRules += "MaxLengthCheck, "
		} else if !HasUpperCaseCheckResult {
			missingRules += "HasUpperCaseCheck, "
		} else if !HasLowerCaseCheckResult {
			missingRules += "HasLowerCaseCheck, "
		} else if !HasNumberCheckResult {
			missingRules += "HasNumberCheck, "
		} else if !HasSpecialCharCheckResult {
			missingRules += "HasSpecialCharCheck"
		} else {
			missingRules += ""
		}
		return false, "password does not match following rules: "
	}
}
