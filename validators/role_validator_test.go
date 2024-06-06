package validators

import "testing"

func TestRoleValidator(t *testing.T) {
	validator := RoleValidator{}

	tests := []struct {
		email   string
		isValid bool
	}{
		{"admin@example.com", false},
		{"support@example.com", false},
		{"info@example.com", false},
		{"user@example.com", true},
	}

	for _, test := range tests {
		if validator.Validate(test.email) != test.isValid {
			t.Errorf("expected %v for email %s, got %v", test.isValid, test.email, !test.isValid)
		}
	}
}
