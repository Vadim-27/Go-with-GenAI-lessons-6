// Homework — Task 2: extend emailCases and/or phoneCases below to at
// least 8 cases each (your mentor may ask for just one of the two
// functions), then implement validate.go until every subtest passes.
//
// Like todo_test.go, this file uses t.Run per case and t.Errorf (not
// t.Fatalf) for the actual assertions, so one wrong case never hides
// the others. Run `go test -v ./validate/...` and read every FAIL line.
package validate

import (
	"strings"
	"testing"
)

const minCases = 8

// emailCases is the table of test cases for ValidateEmail.
//
// TODO: add at least 5 more cases here — for example: whitespace inside
// the address, a missing domain, a trailing dot, consecutive dots, a
// very long local part, or a unicode character.
var emailCases = []struct {
	name  string
	input string
	want  bool
}{
	{"valid simple", "student@softserve.academy", true},
	{"missing at sign", "student-softserve.academy", false},
	{"empty string", "", false},
	{"valid with subdomain", "a.b@mail.example.com", true},
	{"valid unicode local part", "студент@example.com", true},
	{"space inside", "stud ent@softserve.academy", false},
	{"tab inside", "stud\tent@softserve.academy", false},
	{"two at signs", "a@b@softserve.academy", false},
	{"missing local part", "@softserve.academy", false},
	{"missing domain", "student@", false},
	{"trailing dot in domain", "student@softserve.", false},
	{"exactly 40 chars", strings.Repeat("a", 28) + "@example.com", true},
	{"41 chars is too long", strings.Repeat("a", 29) + "@example.com", false},
}

func TestValidateEmail(t *testing.T) {
	if len(emailCases) < minCases {
		t.Fatalf(
			"emailCases has %d case(s), need at least %d — add more edge cases to validate_test.go before this test can pass",
			len(emailCases), minCases,
		)
	}

	for _, tc := range emailCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateEmail(tc.input)
			if got != tc.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

// phoneCases is the table of test cases for ValidatePhone.
// This is only required if your mentor asked you to validate phone
// numbers instead of (or in addition to) email addresses.
//
// TODO: add at least 5 more cases here — for example: missing digits,
// letters mixed in, an unexpected country code format, or extra
// separators like spaces, dots or parentheses.
var phoneCases = []struct {
	name  string
	input string
	want  bool
}{
	{"valid with plus", "+380501234567", true},
	{"contains letters", "050-abc-4567", false},
	{"empty string", "", false},
	{"valid local with dashes", "050-123-4567", true},
	{"valid with parentheses", "(050)1234567", true},
	{"too short", "+38050", false},
	{"too long (16 digits)", "+3805012345678901", false},
	{"space inside", "050 123 4567", false},
	{"plus in the middle", "050+1234567", false},
	{"uppercase letters", "050-ABC-4567", false},
	{"only separators", "----------", false},
	{"arabic-indic digits", "٠٥٠١٢٣٤٥٦٧", false},
}

func TestValidatePhone(t *testing.T) {
	if len(phoneCases) < minCases {
		t.Fatalf(
			"phoneCases has %d case(s), need at least %d — add more edge cases to validate_test.go before this test can pass",
			len(phoneCases), minCases,
		)
	}

	for _, tc := range phoneCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ValidatePhone(tc.input)
			if got != tc.want {
				t.Errorf("ValidatePhone(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
