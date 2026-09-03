// Command gen writes the Basaltic Go SDK's service packages from the
// platform's OpenAPI specifications.
//
// The specifications live in their own repository and stay the single source
// of truth; nothing is vendored here. Point the generator at a local checkout:
//
//	make generate SPEC=../../openapi
//
// or run it directly — it is a module of its own, so it has to be run from
// its own directory:
//
//	cd internal/gen && go run . -spec ../../../openapi
//
// It rewrites one package per service in the SDK's root. The generated files
// are committed, because `go get` has to work without anyone running this.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// services are the packages the SDK ships.
//
// The platform serves more than these. Deliberately absent: database, registry,
// queue, notifications and email, which are not part of the release; and
// objects, which is S3-compatible and spoken to with an AWS SDK rather than
// this one.
//
// A service dropped from this list leaves its package behind — the generator
// rewrites packages, it does not remove them — so delete the directory in the
// same change, or `go get` keeps handing out a client for something unreleased.
var services = []string{
	"audit",
	"billing",
	"certificate",
	"compute",
	"dns",
	"iam",
	"kms",
	"loadbalancer",
	"network",
	"quota",
	"secrets",
	"storage",
	"telemetry",
}

func main() {
	specDir := flag.String("spec", "", "path to a checkout of the OpenAPI specification repository")
	outDir := flag.String("out", "../..", "SDK module root to write service packages into")
	only := flag.String("only", "", "comma-separated services to regenerate (default: all)")
	allowDirty := flag.Bool("allow-dirty", false, "generate even though the specification checkout has uncommitted changes")
	flag.Parse()

	if *specDir == "" {
		fmt.Fprintln(os.Stderr, "gen: -spec is required: point it at a checkout of the OpenAPI repository")
		flag.Usage()
		os.Exit(2)
	}
	if err := checkSpecClean(*specDir, *allowDirty); err != nil {
		fmt.Fprintln(os.Stderr, "gen:", err)
		os.Exit(1)
	}
	if err := run(*specDir, *outDir, *only); err != nil {
		fmt.Fprintln(os.Stderr, "gen:", err)
		os.Exit(1)
	}
}

// checkSpecClean refuses to generate from a specification checkout with
// uncommitted changes.
//
// The generated code is committed, so generating from a dirty tree bakes
// somebody's half-finished edit into a release. That is not hypothetical: this
// check exists because a run against a working tree mid-edit produced an SDK
// three operations short, and the only sign was a number in the summary line.
//
// A repository shared with other sessions is the normal case here, so the
// remedy is usually to generate from committed state rather than to wait:
//
//	git -C <spec> archive HEAD | tar -x -C /tmp/spec-head
//	go run . -spec /tmp/spec-head
func checkSpecClean(specDir string, allowDirty bool) error {
	if specDir == "" || allowDirty {
		return nil
	}
	cmd := exec.Command("git", "-C", specDir, "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		// Not a git checkout, or no git. Not a reason to refuse.
		return nil
	}
	dirty := strings.TrimSpace(string(out))
	if dirty == "" {
		return nil
	}
	n := len(strings.Split(dirty, "\n"))
	return fmt.Errorf(
		"the specification checkout at %s has %d uncommitted change(s):\n%s\n\n"+
			"Generated code is committed, so generating from a dirty tree would bake an\n"+
			"unfinished edit into the SDK. Generate from committed state instead:\n\n"+
			"    git -C %s archive HEAD | tar -x -C /tmp/spec-head\n"+
			"    go run . -spec /tmp/spec-head\n\n"+
			"or pass -allow-dirty if the changes are yours and you mean to include them",
		specDir, n, indentLines(dirty), specDir)
}

func indentLines(s string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = "    " + l
	}
	return strings.Join(lines, "\n")
}

func run(specDir, outDir, only string) error {
	specDir, err := filepath.Abs(specDir)
	if err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(specDir, "components", "schemas")); err != nil {
		return fmt.Errorf("%s does not look like the OpenAPI repository (no components/schemas): %w", specDir, err)
	}

	wanted := services
	if only != "" {
		wanted = nil
		for _, s := range strings.Split(only, ",") {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if !contains(services, s) {
				return fmt.Errorf("unknown service %q", s)
			}
			wanted = append(wanted, s)
		}
	}

	man := manifest{Module: "github.com/basaltic-sh/sdk-go", Version: sdkVersion()}
	totalOps, totalTypes := 0, 0
	for _, svc := range wanted {
		ops, types, ms, err := generate(specDir, outDir, svc)
		if err != nil {
			return fmt.Errorf("%s: %w", svc, err)
		}
		totalOps += ops
		totalTypes += types
		man.Services = append(man.Services, ms)
		fmt.Printf("%-14s %3d operations  %3d types\n", svc, ops, types)
	}
	// Regenerating a subset would otherwise drop the rest from the manifest.
	if only == "" {
		if err := writeManifest(outDir, man); err != nil {
			return fmt.Errorf("writing the manifest: %w", err)
		}
	}
	fmt.Printf("%-14s %3d operations  %3d types across %d services\n", "total", totalOps, totalTypes, len(wanted))
	return nil
}

func generate(specDir, outDir, svc string) (int, int, manifestService, error) {
	ld := newLoader(specDir)
	specFile := filepath.Join(specDir, svc, "openapi.yaml")
	spec, err := ld.load(specFile)
	if err != nil {
		return 0, 0, manifestService{}, err
	}

	b := newBuilder(ld, svc, specDir)
	if err := b.readServer(spec); err != nil {
		return 0, 0, manifestService{}, err
	}
	if err := b.buildOperations(spec, specFile); err != nil {
		return 0, 0, manifestService{}, err
	}
	info := obj(spec, "info")
	if err := emitService(filepath.Join(outDir, svc), b, info); err != nil {
		return 0, 0, manifestService{}, err
	}
	return len(b.ops), len(b.types), b.buildManifest(info), nil
}

// sdkVersion records which surface the manifest describes.
//
// Taken from the repository's most recent tag rather than a constant: a
// constant has to be edited in the same commit as the tag, and the one that
// used to be here was two releases stale before anyone noticed.
func sdkVersion() string {
	out, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		return "dev"
	}
	return strings.TrimPrefix(strings.TrimSpace(string(out)), "v")
}

// readServer takes the service's endpoint template from its specification, so
// the SDK does not carry a second list of which services are regional.
func (b *builder) readServer(spec map[string]any) error {
	servers := list(spec, "servers")
	if len(servers) == 0 {
		return fmt.Errorf("specification declares no servers")
	}
	sm, ok := mapOf(servers[0])
	if !ok {
		return fmt.Errorf("malformed servers entry")
	}
	url := str(sm, "url")
	if url == "" {
		return fmt.Errorf("first server has no url")
	}
	b.serverTemplate = url
	return nil
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
