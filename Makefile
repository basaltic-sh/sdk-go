# The OpenAPI specifications are the single source of truth and live in their
# own repository. Nothing is vendored here; point SPEC at a local checkout.
SPEC ?= ../openapi

.PHONY: all
all: test vet

.PHONY: generate
generate:
	@test -d "$(SPEC)/components/schemas" || { \
		echo "SPEC=$(SPEC) is not an OpenAPI checkout. Try: make generate SPEC=/path/to/openapi"; \
		exit 1; \
	}
	cd internal/gen && go run . -spec "$(abspath $(SPEC))" -out ../..
	gofmt -l . | grep -v '^internal/gen/' && exit 1 || true

.PHONY: test
test:
	go test ./...
	cd internal/gen && go test ./...

.PHONY: vet
vet:
	go vet ./...
	cd internal/gen && go vet ./...

# Fails when the committed output no longer matches what the specifications
# produce. Run it after changing either the generator or the specs.
.PHONY: check-generated
check-generated: generate
	@git diff --exit-code -- . ':!internal/gen' || { \
		echo; \
		echo "Generated code is out of date. Commit the regenerated files."; \
		exit 1; \
	}

.PHONY: doc
doc:
	go doc -all . | head -100
