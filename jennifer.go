// Package jennifer is a code generator for Go
package jennifer

//go:generate go install github.com/mcrawfo2/jennifer/genjen
//go:generate genjen
//go:generate go install github.com/mcrawfo2/jennifer/gennames
//go:generate gennames -output "jen/hints.go" -package "jen" -name "standardLibraryHints" -standard -novendor -path "./..."
//go:generate go install github.com/mcrawfo2/rebecca/cmd/becca
//go:generate becca -package=github.com/mcrawfo2/jennifer/jen
