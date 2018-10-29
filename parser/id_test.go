package parser_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/opsidian/ocl/parser"
	"github.com/opsidian/ocl/test"
)

var _ = Describe("ID", func() {

	p := parser.ID()

	DescribeTable("it evaluates the input correctly",
		func(input string, expected interface{}) {
			test.ExpectParserToEvaluate(p)(input, expected)
		},
		test.TableEntry(`a`, "a"),
		test.TableEntry(`a_b`, "a_b"),
		test.TableEntry(`abcdefghijklmnopqrstuvwxyz_0123456789`, "abcdefghijklmnopqrstuvwxyz_0123456789"),
	)

	DescribeTable("it returns a parse error",
		func(input string, expectedErr error) {
			test.ExpectParserToHaveParseError(p)(input, expectedErr)
		},
		test.TableEntry(`testkeyword`, errors.New("testkeyword is a reserved keyword at testfile:1:1")),
		test.TableEntry(`a__b`, errors.New("was expecting the end of input at testfile:1:2")),
		test.TableEntry(`_b`, errors.New("was expecting identifier at testfile:1:1")),
		test.TableEntry(`0ab`, errors.New("was expecting identifier at testfile:1:1")),
	)

})