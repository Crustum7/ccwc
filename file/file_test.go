package file

import (
	"os"
	"testing"
)

func TestFileInfo(t *testing.T) {
	name := "../test.txt"
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}
	defer file.Close()

	info, err := NewFileInfo(file)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}

	bytes := 342190
	if info.bytes != bytes {
		t.Fatalf(`Expected Bytes() to return %v bytes but returned %v bytes`, bytes, info.bytes)
	}

	lines := 7145
	if info.lines != lines {
		t.Fatalf(`Expected FileParseInfo() to return %v lines but returned %v lines`, lines, info.lines)
	}

	words := 58164
	if info.words != words {
		t.Fatalf(`Expected FileParseInfo() to return %v words but returned %v words`, words, info.words)
	}

	chars := 339292
	if info.chars != chars {
		t.Fatalf(`Expected FileParseInfo() to return %v chars but returned %v chars`, chars, info.chars)
	}
}

func TestFileInfoStdin(t *testing.T) {
	name := "../test.txt"
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}
	defer file.Close()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = file

	info, err := NewFileInfo(file)
	if err != nil {
		t.Fatalf(`Incorrect test: file %v not found`, name)
	}

	bytes := 342190
	if info.bytes != bytes {
		t.Fatalf(`Expected Bytes() to return %v bytes but returned %v bytes`, bytes, info.bytes)
	}

	lines := 7145
	if info.lines != lines {
		t.Fatalf(`Expected FileParseInfo() to return %v lines but returned %v lines`, lines, info.lines)
	}

	words := 58164
	if info.words != words {
		t.Fatalf(`Expected FileParseInfo() to return %v words but returned %v words`, words, info.words)
	}

	chars := 339292
	if info.chars != chars {
		t.Fatalf(`Expected FileParseInfo() to return %v chars but returned %v chars`, chars, info.chars)
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
