package main

import (
	"github.com/golangci/plugin-module-register/register"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var Analyzer = &analysis.Analyzer{
	Name: "todo",
	Doc:  "finds todos without author",
	Run:  run,
}

func init() {
	register.Plugin("example", NewPlugin)
}

type MySettings struct {
	One   string    `json:"one"`
	Two   []Element `json:"two"`
	Three Element   `json:"three"`
}

type Element struct {
	Name string `json:"name"`
}

type PluginExample struct {
	settings MySettings
}

func New(settings any) ([]*analysis.Analyzer, error) {
	plugin, err := NewPlugin(settings)
	if err != nil {
		return nil, err
	}

	return plugin.BuildAnalyzers()
}

func NewPlugin(settings any) (register.LinterPlugin, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	s, err := register.DecodeSettings[MySettings](settings)
	if err != nil {
		return nil, err
	}

	return &PluginExample{settings: s}, nil
}

func (f *PluginExample) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "todo",
			Doc:  "finds todos without author",
			Run:  run,
		},
	}, nil
}

func (f *PluginExample) GetLoadMode() string {
	return register.LoadModeSyntax
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if goStmt, ok := n.(*ast.GoStmt); ok {
				pass.Reportf(
					goStmt.Pos(),
					`"go" keyword usage is restricted. Use future/async wrapper instead.`,
				)
			}
			return true
		})
	}
	return nil, nil
}

func main() {
	singlechecker.Main(Analyzer)
}
