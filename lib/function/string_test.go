package function_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	basilfunction "github.com/opsidian/basil/function"
	"github.com/opsidian/basil/lib/function"
	"github.com/opsidian/basil/parser"
	"github.com/opsidian/basil/test"
	"github.com/opsidian/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

var _ = Describe("String", func() {

	registry := basilfunction.Registry{
		"test": function.StringInterpreter{},
	}

	DescribeTable("it evaluates the input correctly",
		func(input string, expected interface{}) {
			test.ExpectFunctionToEvaluate(parser.Expression(), registry)(input, expected)
		},
		test.TableEntry(`test(1)`, "1"),
		test.TableEntry(`test(1.1)`, "1.1"),
		test.TableEntry(`test(false)`, "false"),
		test.TableEntry(`test(true)`, "true"),
		test.TableEntry(`test("foo")`, "foo"),
		test.TableEntry(`test(1m30s)`, "1m30s"),
	)

	DescribeTable("it will have a parse error",
		func(input string, expectedErr error) {
			test.ExpectFunctionToHaveParseError(parser.Expression(), registry)(input, expectedErr)
		},
		test.TableEntry(`test()`, errors.New("test expects 1 arguments at testfile:1:1")),
		test.TableEntry(`test(1, 2)`, errors.New("test expects 1 arguments at testfile:1:1")),
		test.TableEntry(`test([])`, errors.New("was expecting any basic type at testfile:1:6")),
	)

	It("should return with string type", func() {
		test.ExpectFunctionNode(parser.Expression(), registry)(
			"test(1)",
			func(userCtx interface{}, node parsley.Node) {
				Expect(node.Type()).To(Equal(variable.TypeString))
			},
		)
	})

})