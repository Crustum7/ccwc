package file

import (
	"os"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestOpenFile(t *testing.T) {
	name := "../test.txt"
	file, err := OpenFile(name)
	if err != nil {
		t.Fatalf(`OpenFile("%v") returned error %v`, name, err)
	}
	defer file.Close()
}

func TestBytes(t *testing.T) {
	name := "../test.txt"
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}
	bytes, err := Bytes(file)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}
	target := int64(342190)
	if bytes != target {
		t.Fatalf(`Expected Bytes() to return %v bytes but returned %v bytes`, target, bytes)
	}
}
