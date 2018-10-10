package goscripter

import (
	"testing"
)

func TestBuildJavascript(t *testing.T) {
	type testbuild struct {
		body   string
		expect string
	}
	testcases := []testbuild{
		{
			body:   "console.log('Awkarin is back!');",
			expect: "<script type='text/javascript'>console.log('Awkarin is back!');</script>",
		},
	}
	for _, vtc := range testcases {
		actual := BuildJavascript(vtc.body)
		if actual != vtc.expect {
			t.Errorf("Mismatched on testcase : %+v", vtc)
		}
	}
}
func TestBuildCSS(t *testing.T) {
	type testbuild struct {
		body   string
		expect string
	}
	testcases := []testbuild{
		{
			body:   "#awkarin{background:#AAAAAA;}",
			expect: "<style>#awkarin{background:#AAAAAA;}</style>",
		},
	}
	for _, vtc := range testcases {
		actual := BuildCSS(vtc.body)
		if actual != vtc.expect {
			t.Errorf("Mismatched on testcase : %+v", vtc)
		}
	}
}
