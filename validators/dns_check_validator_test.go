package validators

import "testing"

func TestDNSCheckValidator(t *testing.T) {
	validator := DNSCheckValidator{}

	tests := []struct {
		email   string
		isValid bool
	}{
		{"user@gmail.com", true},
		{"user@invalid-domain.com", false},
	}

	for _, test := range tests {
		if validator.Validate(test.email) != test.isValid {
			t.Errorf("expected %v for email %s, got %v", test.isValid, test.email, !test.isValid)
		}
	}
}
