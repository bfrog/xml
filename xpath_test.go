package golibxml

import "testing"

func TestXPathCompile(t *testing.T) {
	xpath := CompileXPath("*")
	if xpath == nil {
		t.Fail()
	}
}

func TestXPathCompileNeg(t *testing.T) {
	xpath := CompileXPath("")
	if xpath != nil {
		t.Fail()
	}
}
