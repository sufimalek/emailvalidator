package emailvalidator

import (
	"emailvalidator/validators"
	"testing"
)

func TestValidatorChain(t *testing.T) {
	chain := NewValidatorChainBuilder().
		Add(&validators.SyntaxValidator{}).
		Add(&validators.RoleValidator{}).
		Add(&validators.MXValidator{}).
		Add(&validators.SMTPValidator{}).
		Add(validators.NewBanWordsUsernameValidator([]string{"spam", "junk"})).
		Add(validators.NewBlackListEmailsValidator([]string{"blacklisted@example.com"})).
		Add(&validators.DNSCheckValidator{}).
		Add(&validators.RFCValidator{}).
		Add(&validators.NoRFCWarningsValidator{}).
		// Add(&validators.GravatarValidator{}).
		Build()

	tests := []struct {
		email   string
		isValid bool
	}{
		{"admin@example.com", false},
		{"user@gmail.com", false},
		{"spamuser@example.com", false},
		{"blacklisted@example.com", false},
		{"invalid@invalid-domain.com", false},
		{"user@.com", false},
		{"@example.com", false},
		{"good@domain.com", true},
	}

	for _, test := range tests {
		if chain.Validate(test.email) != test.isValid {
			t.Errorf("expected %v for email %s, got %v", test.isValid, test.email, !test.isValid)
		}
	}
}
