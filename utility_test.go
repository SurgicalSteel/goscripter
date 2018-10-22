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
		actual := BuildJavascriptBundle(vtc.body)
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
		actual := BuildCSSBundle(vtc.body)
		if actual != vtc.expect {
			t.Errorf("Mismatched on testcase : %+v", vtc)
		}
	}
}
func TestIsValidJSONString(t *testing.T) {
	type testValidJSONString struct {
		paramValue    string
		expectedValid bool
		actualValid   bool
	}
	testCases := []testValidJSONString{
		{paramValue: "I still love my ex-girlfriend", expectedValid: false},
		{paramValue: "{\"data\":\"bla bla bla\"}", expectedValid: true},
		{paramValue: "{\"data\":{\"productID\":666,\"productName\":\"Jenglot\",\"productPrice\":45000,\"productPriceCurrency\":\"USD\"}}", expectedValid: true},
		{paramValue: "{666}", expectedValid: false},
		{paramValue: "{data:666}", expectedValid: false},
		{paramValue: "{\"productIDs\":[123,234,345,456,567,678,789]}", expectedValid: true},
	}
	for _, test := range testCases {
		test.actualValid = ValidateJSON(test.paramValue)
		if test.actualValid != test.expectedValid {
			t.Error("Result JSON Validation mismatched. Expected : ", test.expectedValid, " But got ", test.actualValid)
		}
	}
}
