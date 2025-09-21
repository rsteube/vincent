package main

import "github.com/rsteube/vincent/cmd/vincent/cmd"

var version = "develop"

//go:generate go run ../../generate/gen.go
func main() {
	cmd.Execute(version)
}
