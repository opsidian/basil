package parser_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/opsidian/ocl/parser"
	"github.com/opsidian/ocl/test"
	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/combinator"
	"github.com/opsidian/parsley/parsley"
	"github.com/opsidian/parsley/text/terminal"
)

var _ = Describe("And", func() {

	q := combinator.Choice(
		terminal.Bool("true", "false"),
		terminal.Integer(),
		terminal.Word("nil", nil),
		test.EvalErrorParser(),
	).ReturnError("was expecting value")

	p := parser.And(q)

	DescribeTable("it evaluates the input correctly",
		func(input string, expected interface{}) {
			test.ExpectParserToEvaluate(p)(input, expected)
		},
		test.TableEntry("1", int64(1)),
		test.TableEntry("nil", nil),
		test.TableEntry("false", false),
		test.TableEntry("true", true),
		test.TableEntry("false && true", false),
		test.TableEntry("true && true", true),
		test.TableEntry("true && true && false", false),
	)

	DescribeTable("it returns a parse error",
		func(input string, expectedErr error) {
			test.ExpectParserToHaveParseError(p)(input, expectedErr)
		},
		test.TableEntry("true &&", errors.New("was expecting value at testfile:1:8")),
	)

	DescribeTable("it returns an eval error",
		func(input string, expectedErr error) {
			test.ExpectParserToHaveEvalError(p)(input, expectedErr)
		},
		test.TableEntry("1 && 2", errors.New("unsupported && operation on int64 at testfile:1:1")),
		test.TableEntry("nil && true", errors.New("unsupported && operation on <nil> at testfile:1:1")),
		test.TableEntry("true && 1", errors.New("unsupported && operation on int64 at testfile:1:6")),
		test.TableEntry("ERR", errors.New("ERR at testfile:1:1")),
		test.TableEntry("true && ERR", errors.New("ERR at testfile:1:9")),
	)

	Context("When there is only one node", func() {
		It("should return the node", func() {
			expectedNode := ast.NewTerminalNode("INT", int64(1), parsley.Pos(1), parsley.Pos(2))
			test.ExpectParserToReturn(p, "1", expectedNode)
		})
	})

})