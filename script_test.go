package goscripter

import (
	"github.com/stretchr/testify/assert"
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

func TestInitialize(t *testing.T) {
	// func Initialize(path string, kinds []FileType) (OurScript, error)
	path := "files/scripts"
	kinds := []FileType{CSS, JS, JSON}
	expected := OurScript{
		RawSlice: []ScriptFile{
			{
				Name:     "stylereference",
				FileName: "stylereference.css",
				Kind:     CSS,
				Body: `#container{background: #AAAAAA;}
`,
			},
			{
				Name:     "scriptreference",
				FileName: "scriptreference.js",
				Kind:     JS,
				Body: `alert("Awkarin is back!");
`,
			},
			{
				Name:     "jsonreference",
				FileName: "jsonreference.json",
				Kind:     JSON,
				Body: `[{"name":"James Bond","age":44,"sex":"male"},{"name":"Awkarin","age":20,"sex":"female"}]
`,
			},
		},
		Map: map[FileType]map[string]ScriptFile{
			CSS: {
				"stylereference": ScriptFile{
					Name:     "stylereference",
					FileName: "stylereference.css",
					Kind:     CSS,
					Body: `#container{background: #AAAAAA;}
`,
				},
			},
			JS: {
				"scriptreference": ScriptFile{
					Name:     "scriptreference",
					FileName: "scriptreference.js",
					Kind:     JS,
					Body: `alert("Awkarin is back!");
`,
				},
			},
			JSON: {
				"jsonreference": ScriptFile{
					Name:     "jsonreference",
					FileName: "jsonreference.json",
					Kind:     JSON,
					Body: `[{"name":"James Bond","age":44,"sex":"male"},{"name":"Awkarin","age":20,"sex":"female"}]
`,
				},
			},
		},
	}
	actual, err := Initialize(path, kinds)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFindScript(t *testing.T) {
	ourScript := OurScript{
		RawSlice: []ScriptFile{
			{
				Name:     "stylereference",
				FileName: "stylereference.css",
				Kind:     CSS,
				Body:     `#container{background: #AAAAAA;}`,
			},
			{
				Name:     "scriptreference",
				FileName: "scriptreference.js",
				Kind:     JS,
				Body:     `alert("Awkarin is back!");`,
			},
			{
				Name:     "jsonreference",
				FileName: "jsonreference.json",
				Kind:     JSON,
				Body:     `[{"name":"James Bond","age":44,"sex":"male"},{"name":"Awkarin","age":20,"sex":"female"}]`,
			},
		},
		Map: map[FileType]map[string]ScriptFile{
			CSS: {
				"stylereference": ScriptFile{
					Name:     "stylereference",
					FileName: "stylereference.css",
					Kind:     CSS,
					Body:     `#container{background: #AAAAAA;}`,
				},
			},
			JS: {
				"scriptreference": ScriptFile{
					Name:     "scriptreference",
					FileName: "scriptreference.js",
					Kind:     JS,
					Body:     `alert("Awkarin is back!");`,
				},
			},
			JSON: {
				"jsonreference": ScriptFile{
					Name:     "jsonreference",
					FileName: "jsonreference.json",
					Kind:     JSON,
					Body:     `[{"name":"James Bond","age":44,"sex":"male"},{"name":"Awkarin","age":20,"sex":"female"}]`,
				},
			},
		},
	}
	scriptItems := []ScriptItem{
		{
			Name: "scriptreference",
			Kind: JS,
		},
		{
			Name: "jsonreference",
			Kind: JSON,
		},
	}
	expected := map[FileType][]ScriptFile{
		JS: []ScriptFile{
			{
				Name:     "scriptreference",
				FileName: "scriptreference.js",
				Kind:     JS,
				Body:     `alert("Awkarin is back!");`,
			},
		},
		JSON: []ScriptFile{
			{
				Name:     "jsonreference",
				FileName: "jsonreference.json",
				Kind:     JSON,
				Body:     `[{"name":"James Bond","age":44,"sex":"male"},{"name":"Awkarin","age":20,"sex":"female"}]`,
			},
		},
	}

	actual := ourScript.FindScripts(scriptItems)
	assert.Equal(t, expected, actual)
}

func testFindBundledScripts(t *testing.T) {
	ourScript := OurScript{
		RawSlice: []ScriptFile{
			{
				Name:     "stylereference",
				FileName: "stylereference.css",
				Kind:     CSS,
				Body:     `#container{background: #AAAAAA;}`,
			},
			{
				Name:     "awkarinstylereference",
				FileName: "awkarinstylereference.css",
				Kind:     CSS,
				Body:     `#awkarincontainer{background: #000000;}`,
			},
			{
				Name:     "weirdstylereference",
				FileName: "weirdstylereference.css",
				Kind:     CSS,
				Body:     `#weirdcontainer{background: #FFFFFF;}`,
			},
			{
				Name:     "scriptreference",
				FileName: "scriptreference.js",
				Kind:     JS,
				Body:     "alert('Awkarin is back!');",
			},
			{
				Name:     "weirdscriptreference",
				FileName: "weirdscriptreference.js",
				Kind:     JS,
				Body:     "alert('Awkarin is coming....');",
			},
			{
				Name:     "jsonreference",
				FileName: "jsonreference.json",
				Kind:     JSON,
				Body:     `[{"name":"James Bond","age":44,"sex":"male"},{"name":"Awkarin","age":20,"sex":"female"}]`,
			},
		},
		Map: map[FileType]map[string]ScriptFile{
			CSS: {
				"stylereference": ScriptFile{
					Name:     "stylereference",
					FileName: "stylereference.css",
					Kind:     CSS,
					Body:     `#container{background: #AAAAAA;}`,
				},
				"awkarinstylereference": ScriptFile{
					Name:     "awkarinstylereference",
					FileName: "awkarinstylereference.css",
					Kind:     CSS,
					Body:     `#awkarincontainer{background: #000000;}`,
				},
				"weirdstylereference": ScriptFile{
					Name:     "weirdstylereference",
					FileName: "weirdstylereference.css",
					Kind:     CSS,
					Body:     `#weirdcontainer{background: #FFFFFF;}`,
				},
			},
			JS: {
				"scriptreference": ScriptFile{
					Name:     "scriptreference",
					FileName: "scriptreference.js",
					Kind:     JS,
					Body:     "alert('Awkarin is back!');",
				},
				"weirdscriptreference": ScriptFile{
					Name:     "weirdscriptreference",
					FileName: "weirdscriptreference.js",
					Kind:     JS,
					Body:     "alert('Awkarin is coming....');",
				},
			},
			JSON: {
				"jsonreference": ScriptFile{
					Name:     "jsonreference",
					FileName: "jsonreference.json",
					Kind:     JSON,
					Body:     `[{"name":"James Bond","age":44,"sex":"male"},{"name":"Awkarin","age":20,"sex":"female"}]`,
				},
			},
		},
	}
	bundleMap := make(map[FileType][]string)
	bundleMap[CSS] = []string{"stylereference", "awkarinstylereference"}
	bundleMap[JS] = []string{"weirdscriptreference"}
	expected := make(map[FileType]string)
	expected[CSS] = "<style>#container{background: #AAAAAA;}#awkarincontainer{background: #000000;}</style>"
	expected[JS] = "<script type='text/javascript'>alert('Awkarin is coming....');</script>"
	actual := ourScript.FindBundledScripts(bundleMap)
	assert.Equal(t, expected, actual)
}
