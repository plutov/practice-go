package functionfrequency

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
)

// Call is a count of a function call
type Call struct {
	Name  string
	Count int
}

// CallList is a list of call counts, sortable on count
type CallList []Call

func (l CallList) Len() int           { return len(l) }
func (l CallList) Less(i, j int) bool { return l[i].Count < l[j].Count }
func (l CallList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

// FunctionFrequency returns the top 3 most mentioned functions in the code sample
func FunctionFrequency(gocode []byte) []string {

	// Create the AST by parsing src.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", gocode, 0)
	if err != nil {
		return nil
	}

	// Print the AST for debug
	//	ast.Print(fset, f)

	// Get a frequency list of all function names cited
	freqencies := make(map[string]int)
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		// Find all function calls
		case *ast.CallExpr:
			switch tt := x.Fun.(type) {
			// Handle plain old function calls
			case *ast.Ident:
				name := tt.Name
				freqencies[name] = freqencies[name] + 1
			// Handle package calls
			case *ast.SelectorExpr:
				switch xx := tt.X.(type) {
				case *ast.Ident:
					name := xx.Name + "." + tt.Sel.Name
					freqencies[name] = freqencies[name] + 1
				}
			}
		}
		return true
	})

	// Store frequencies and sort
	var Calls CallList
	for k, v := range freqencies {
		Calls = append(Calls, Call{Name: k, Count: v})
	}
	sort.Sort(sort.Reverse(Calls))

	// Get a list of the top 3 Calls as strings
	var names []string
	for i, c := range Calls {
		names = append(names, c.Name)
		if i > 1 {
			break
		}
	}

	return names
}
