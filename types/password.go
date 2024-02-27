package types

type PasswordRules struct {
	MinLengthCheck bool
	MinLength int
	MaxLengthCheck bool
	MaxLength int
	HasUpperCaseCheck bool
	HasLowerCaseCheck bool
	HasNumberCheck bool
	HasSpecialCharCheck bool
}
