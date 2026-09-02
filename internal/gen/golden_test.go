package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

// The golden test pins the generator's textual output — doc comments, field
// order, pointer-versus-value, everything — against a fixture specification
// shaped like the real ones.
//
// It is deliberately not a test of the generated SDK. The surface tests in the
// root package check that the emitted code behaves; this checks that
// refactoring the emitter does not change what it writes. The two catch
// different faults, and this one runs without the specification repository,
// which is private and so unavailable to public CI.
//
// Update with: go test ./internal/gen -update

var update = flag.Bool("update", false, "rewrite the golden files from current output")

func TestGeneratorOutputMatchesGolden(t *testing.T) {
	specDir, err := filepath.Abs("testdata/spec")
	if err != nil {
		t.Fatal(err)
	}
	outDir := t.TempDir()

	ops, types, err := generate(specDir, outDir, "widget")
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if ops != 6 {
		t.Errorf("generated %d operations from the fixture, want 6", ops)
	}
	if types == 0 {
		t.Error("generated no types")
	}

	for _, name := range []string{"client_gen.go", "types_gen.go", "api_gen.go"} {
		got, err := os.ReadFile(filepath.Join(outDir, "widget", name))
		if err != nil {
			t.Fatalf("reading generated %s: %v", name, err)
		}
		goldenPath := filepath.Join("testdata", "golden", name)

		if *update {
			if err := os.MkdirAll(filepath.Dir(goldenPath), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(goldenPath, got, 0o644); err != nil {
				t.Fatal(err)
			}
			t.Logf("updated %s", goldenPath)
			continue
		}

		want, err := os.ReadFile(goldenPath)
		if err != nil {
			t.Fatalf("reading golden %s (run: go test ./internal/gen -update): %v", name, err)
		}
		if string(got) != string(want) {
			t.Errorf("%s differs from its golden file.\n"+
				"If the change is intended, re-run with -update and review the diff.\n"+
				"%s", name, firstDifference(string(want), string(got)))
		}
	}
}

// firstDifference reports the first differing line, which is far easier to
// read than a whole-file dump of two generated sources.
func firstDifference(want, got string) string {
	wantLines, gotLines := splitLines(want), splitLines(got)
	for i := 0; i < len(wantLines) || i < len(gotLines); i++ {
		w, g := "", ""
		if i < len(wantLines) {
			w = wantLines[i]
		}
		if i < len(gotLines) {
			g = gotLines[i]
		}
		if w != g {
			return "first difference at line " + itoa(i+1) + ":\n  want: " + w + "\n  got:  " + g
		}
	}
	return "(files differ only in trailing content)"
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
