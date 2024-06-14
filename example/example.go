package main

import (
	"github.com/golangci/plugin-module-register/register"
	example "github.com/ntrivix/goroutine-forbid-linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func init() {
	register.Plugin("example", example.NewPlugin)
}

func main() {
	singlechecker.Main(example.Analyzer)
}
