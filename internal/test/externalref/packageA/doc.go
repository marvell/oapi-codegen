package packageA

//go:generate go run github.com/marvell/oapi-codegen/cmd/oapi-codegen -generate types,skip-prune --package=packageA -o externalref.gen.go --import-mapping=../packageB/spec.yaml:github.com/marvell/oapi-codegen/internal/test/externalref/packageB spec.yaml
