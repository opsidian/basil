package function_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	basilfunction "github.com/opsidian/basil/basil/function"
	"github.com/opsidian/basil/basil/variable"
	"github.com/opsidian/basil/function"
	"github.com/opsidian/basil/parser"
	"github.com/opsidian/basil/test"
	"github.com/opsidian/parsley/parsley"
)

var _ = Describe("HasPrefix", func() {

	registry := basilfunction.InterpreterRegistry{
		"test": function.HasPrefixInterpreter{},
	}

	DescribeTable("it evaluates the input correctly",
		func(input string, expected interface{}) {
			test.ExpectFunctionToEvaluate(parser.Expression(), registry)(input, expected)
		},
		test.TableEntry(`test("", "")`, true),
		test.TableEntry(`test("foo", "")`, true),
		test.TableEntry(`test("foo", "fo")`, true),
		test.TableEntry(`test("foo", "fox")`, false),
		test.TableEntry(`test("foo", "Fo")`, false),
	)

	DescribeTable("it will have a parse error",
		func(input string, expectedErr error) {
			test.ExpectFunctionToHaveParseError(parser.Expression(), registry)(input, expectedErr)
		},
		test.TableEntry(`test("foo")`, errors.New("test expects 2 arguments at testfile:1:1")),
		test.TableEntry(`test("a", "a", "a")`, errors.New("test expects 2 arguments at testfile:1:1")),
		test.TableEntry(`test(1, "a")`, errors.New("was expecting string at testfile:1:6")),
		test.TableEntry(`test("a", 1)`, errors.New("was expecting string at testfile:1:11")),
	)

	It("should return with boolean type", func() {
		test.ExpectFunctionNode(parser.Expression(), registry)(
			`test("", "")`,
			func(userCtx interface{}, node parsley.Node) {
				Expect(node.Type()).To(Equal(variable.TypeBool))
			},
		)
	})

})