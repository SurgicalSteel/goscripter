package goscripter

import (
	"testing"
)

func TestIsOfKind(t *testing.T) {
	type teststruct struct {
		allkind []FileType
		kind    FileType
		expect  bool
	}
	testcases := []teststruct{
		{
			allkind: []FileType{FileType("csv"), JS},
			kind:    CSS,
			expect:  false,
		},
		{
			allkind: []FileType{JS, CSS},
			kind:    CSS,
			expect:  true,
		},
	}
	for _, v := range testcases {
		actual := IsOfKind(v.kind, v.allkind)
		if actual != v.expect {
			t.Errorf("Mismatched on testcase : %+v", v)
		}
	}
}
