package parser

import (
	"testing"

	"github.com/ncw/gpython/ast"
)

// FIXME test pos is correct

func TestGrammar(t *testing.T) {
	for _, test := range []struct {
		in   string
		mode string
		out  string
	}{
		// START TESTS
		{"", "exec", "Module(body=[])"},
		{"pass", "exec", "Module(body=[Pass()])"},
		{"()", "eval", "Expression(body=Tuple(elts=[], ctx=Load()))"},
		{"()", "exec", "Module(body=[Expr(value=Tuple(elts=[], ctx=Load()))])"},
		{"[ ]", "exec", "Module(body=[Expr(value=List(elts=[], ctx=Load()))])"},
		// END TESTS
	} {
		Ast, err := ParseString(test.in, test.mode)
		if err != nil {
			t.Errorf("Parse(%q) returned error: %v", test.in, err)
		} else {
			out := ast.Dump(Ast)
			if out != test.out {
				t.Errorf("Parse(%q) expecting %q actual %q", test.in, test.out, out)
			}
		}
	}
}
