package validators

import "testing"

func TestBlackListEmailsValidator(t *testing.T) {
	blacklistedEmails := []string{"blacklisted@example.com"}
	validator := NewBlackListEmailsValidator(blacklistedEmails)

	tests := []struct {
		email   string
		isValid bool
	}{
		{"blacklisted@example.com", false},
		{"user@example.com", true},
	}

	for _, test := range tests {
		if validator.Validate(test.email) != test.isValid {
			t.Errorf("expected %v for email %s, got %v", test.isValid, test.email, !test.isValid)
		}
	}
}
