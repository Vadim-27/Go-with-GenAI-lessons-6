package validate

import (
	"strings"
	"testing"
)

// Тести згенеровано ШІ «наосліп» — лише за сигнатурою
// func ValidateEmail(s string) bool
// Промпт: «для функції ValidateEmail напиши прості табличні тести».
func TestValidateEmail_AIGenerated(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid simple", "user@example.com", true},
		{"valid with subdomain", "user@mail.example.com", true},
		{"valid with dot in local", "first.last@example.com", true},
		{"valid with plus tag", "user+tag@example.com", true},
		{"valid with digits", "user123@example123.com", true},
		{"valid with hyphen in domain", "user@my-domain.com", true},
		{"valid uppercase", "User@Example.COM", true},
		{"empty string", "", false},
		{"missing at", "userexample.com", false},
		{"multiple at", "user@@example.com", false},
		{"two at in different places", "a@b@example.com", false},
		{"missing local part", "@example.com", false},
		{"missing domain", "user@", false},
		{"missing tld", "user@example", false},
		{"space inside", "us er@example.com", false},
		{"leading space", " user@example.com", false},
		{"trailing space", "user@example.com ", false},
		{"tab character", "user\t@example.com", false},
		{"newline character", "user@example.com\n", false},
		{"consecutive dots in local", "us..er@example.com", false},
		{"consecutive dots in domain", "user@example..com", false},
		{"leading dot in local", ".user@example.com", false},
		{"trailing dot in local", "user.@example.com", false},
		{"leading dot in domain", "user@.example.com", false},
		{"trailing dot in domain", "user@example.com.", false},
		{"domain starts with hyphen", "user@-example.com", false},
		{"invalid character in local", "us(er@example.com", false},
		{"invalid character in domain", "user@exa_mple.com", false},
		{"only at sign", "@", false},
		// Розбіжність зі специфікацією проєкту: unicode у локальній частині дозволено.
		{"unicode local part", "юзер@example.com", true},
		{"long local part over 64", strings.Repeat("a", 65) + "@example.com", false},
		{"long address over 254", strings.Repeat("a", 64) + "@" + strings.Repeat("b", 250) + ".com", false},
		// Розбіжність зі специфікацією проєкту: загальний ліміт адреси — 40, а не RFC-шні 64/254.
		{"max valid local part 64", strings.Repeat("a", 64) + "@example.com", false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.input); got != tt.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
