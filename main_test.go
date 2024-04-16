package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/elliotthill/golox/interpreter"
)

type OutAssertTest struct {
    name        string
    syntax      string
    expectedOut string
    expectedErr string
}

var (
    //Output assertion
    assertTests []OutAssertTest = []OutAssertTest{
        {name: "Print()", syntax:"print 'hello';", expectedOut:"hello", expectedErr:""},
        {name: "Equality", syntax:"print 1==1;", expectedOut: "true", expectedErr: ""},
        {name: "Equality String", syntax:"print 'hello'=='hello';", expectedOut: "true", expectedErr: ""},
        {name: "Compare", syntax:"print 1<2;", expectedOut: "true", expectedErr: ""},
        {name: "OOO", syntax:"print 2*(1+1+(2*10));", expectedOut: "44", expectedErr: ""},
    }
)

func TestRun(t *testing.T) {

    var outBuf bytes.Buffer = bytes.Buffer{}
    var errBuf bytes.Buffer = bytes.Buffer{}

    interp := interpreter.NewInterpreter(&outBuf, &errBuf)

    for _, test := range assertTests {

        Run(test.syntax, interp, false);

        output := outBuf.String()
        output = StripAll(output);

        if (output != test.expectedOut) {
            t.Errorf("Got %s, expected %s", strconv.Quote(output), strconv.Quote(test.expectedOut))
        }
        outBuf.Reset()
        errBuf.Reset()
    }



}



func StripAll(str string) string {

    str = strings.ReplaceAll(str, " ", "")
    str = strings.ReplaceAll(str, "\n", "")
    return str
}
