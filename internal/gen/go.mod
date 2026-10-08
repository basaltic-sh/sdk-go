// The generator is a module of its own so the SDK it emits stays free of
// dependencies. Consumers of github.com/basaltic-sh/sdk-go never see the YAML
// parser this needs — a nested module is excluded from the parent's graph.
module github.com/basaltic-sh/sdk-go/internal/gen

go 1.23.0

require gopkg.in/yaml.v3 v3.0.1
