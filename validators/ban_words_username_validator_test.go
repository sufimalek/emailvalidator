package validators

import "testing"

func TestBanWordsUsernameValidator(t *testing.T) {
	bannedWords := []string{"spam", "junk"}
	validator := NewBanWordsUsernameValidator(bannedWords)

	tests := []struct {
		email   string
		isValid bool
	}{
		{"spamuser@example.com", false},
		{"junkuser@example.com", false},
		{"validuser@example.com", true},
	}

	for _, test := range tests {
		if validator.Validate(test.email) != test.isValid {
			t.Errorf("expected %v for email %s, got %v", test.isValid, test.email, !test.isValid)
		}
	}
}
