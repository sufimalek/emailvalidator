package validators

import "testing"

func TestGravatarValidator(t *testing.T) {
	validator := GravatarValidator{}

	tests := []struct {
		email   string
		isValid bool
	}{
		{"existingemail@domain.com", true},     // Replace with an email known to have a Gravatar
		{"nonexistingemail@domain.com", false}, // Replace with an email known to not have a Gravatar
	}

	for _, test := range tests {
		if validator.Validate(test.email) != test.isValid {
			t.Errorf("expected %v for email %s, got %v", test.isValid, test.email, !test.isValid)
		}
	}
}
