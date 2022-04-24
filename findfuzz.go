package findfuzz

import (
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "findfuzz is ..."

type Result struct {
	Events []Event
}

type Event struct {
	Name string
}

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "findfuzz",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	ResultType: reflect.TypeOf(new(Result)),
}

func run(pass *analysis.Pass) (interface{}, error) {
	result := &Result{}
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	if !strings.HasSuffix(pass.Pkg.Name(), "_test") {
		return result, nil
	}

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if strings.HasPrefix(n.Name.Name, "Fuzz") {
				if len(n.Type.Params.List) != 1 {
					return
				}

				for _, v := range n.Type.Params.List {
					typ := pass.TypesInfo.TypeOf(v.Type)
					typFuzz := analysisutil.TypeOf(pass, "testing", "*F")
					if types.Identical(typ, typFuzz) {
						pass.Reportf(n.Pos(), "Fuzz test here")
						event := Event{Name: n.Name.Name[4:]}

						result.Events = append(result.Events, event)
					}
				}
			}
		}
	})

	return result, nil
}
