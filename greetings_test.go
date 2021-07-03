package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "manu"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
