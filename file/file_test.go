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

func TestLines(t *testing.T) {
	name := "../test.txt"
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}
	info, err := FileParseInfo(file)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}
	target := 7145
	if info.lines != target {
		t.Fatalf(`Expected FileParseInfo() to return %v lines but returned %v lines`, target, info.lines)
	}
}

func TestWordsInLine(t *testing.T) {
	line := "a b	c d e, f "
	words := WordsInLine(line)
	target := 6
	if words != target {
		t.Fatalf(`Expected WordsInLine() to return %d words but returned %d words`, target, words)
	}
}

func TestWords(t *testing.T) {
	name := "../test.txt"
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}
	info, err := FileParseInfo(file)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}
	target := 58164
	if info.words != target {
		t.Fatalf(`Expected FileParseInfo() to return %v words but returned %v words`, target, info.words)
	}
}
