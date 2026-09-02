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
	"path/filepath"
	"strings"
)

// services are the packages the SDK ships.
//
// The platform serves more than these. Deliberately absent: registry, queue,
// notifications and email, which are not part of the release; and objects,
// which is S3-compatible and spoken to with an AWS SDK rather than this one.
var services = []string{
	"audit",
	"billing",
	"certificate",
	"compute",
	"database",
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
	flag.Parse()

	if *specDir == "" {
		fmt.Fprintln(os.Stderr, "gen: -spec is required: point it at a checkout of the OpenAPI repository")
		flag.Usage()
		os.Exit(2)
	}
	if err := run(*specDir, *outDir, *only); err != nil {
		fmt.Fprintln(os.Stderr, "gen:", err)
		os.Exit(1)
	}
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

	totalOps, totalTypes := 0, 0
	for _, svc := range wanted {
		ops, types, err := generate(specDir, outDir, svc)
		if err != nil {
			return fmt.Errorf("%s: %w", svc, err)
		}
		totalOps += ops
		totalTypes += types
		fmt.Printf("%-14s %3d operations  %3d types\n", svc, ops, types)
	}
	fmt.Printf("%-14s %3d operations  %3d types across %d services\n", "total", totalOps, totalTypes, len(wanted))
	return nil
}

func generate(specDir, outDir, svc string) (int, int, error) {
	ld := newLoader(specDir)
	specFile := filepath.Join(specDir, svc, "openapi.yaml")
	spec, err := ld.load(specFile)
	if err != nil {
		return 0, 0, err
	}

	b := newBuilder(ld, svc, specDir)
	if err := b.readServer(spec); err != nil {
		return 0, 0, err
	}
	if err := b.buildOperations(spec, specFile); err != nil {
		return 0, 0, err
	}
	if err := emitService(filepath.Join(outDir, svc), b, obj(spec, "info")); err != nil {
		return 0, 0, err
	}
	return len(b.ops), len(b.types), nil
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
