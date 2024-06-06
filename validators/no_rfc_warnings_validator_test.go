package validators

import "testing"

func TestNoRFCWarningsValidator(t *testing.T) {
	validator := NoRFCWarningsValidator{}

	tests := []struct {
		email   string
		isValid bool
	}{
		{"user@example.com", true},
		{"user@.com", false},
		{"@example.com", false},
	}

	for _, test := range tests {
		if validator.Validate(test.email) != test.isValid {
			t.Errorf("expected %v for email %s, got %v", test.isValid, test.email, !test.isValid)
		}
	}
}
