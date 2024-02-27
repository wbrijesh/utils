package types

type EmailValidationRules struct {
	TrustedProviderCheck        bool
	TrustedDomainExtensionCheck bool
	ValidEmailFormatCheck       bool
	TrustedProviders            []string
	TrustedDomainExtensions     []string
}
