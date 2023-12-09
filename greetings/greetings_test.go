package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	//name := "Clive"
	name := "Vic"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Vic")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Vic") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHellosNames calls greetings.Hellos with a slice of names,
// checking for a valid return value.
func TestHellosNames(t *testing.T) {
	//names := []string{"Clive", "Vic", "Bert"}
	names := []string{"Vic", "Bert"}
	want := regexp.MustCompile(`\b`+names[0]+`\b`)
	messages, err := Hellos(names)
	if !want.MatchString(messages[names[0]]) || err != nil {
		t.Fatalf(`Hellos(%q) = %q, %v, want match for %#q, nil`, names, messages, err, want)
	}
}

// TestHellosEmpty calls greetings.Hellos with an empty slice,
// checking for an error.
// func TestHellosEmpty(t *testing.T) {
// 	messages, err := Hellos(nil)
// 	if messages != nil || err == nil {
// 		t.Fatalf(`Hellos(nil) = %q, %v, want nil, error`, messages, err)
// 	}
// }
