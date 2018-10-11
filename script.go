package goscripter

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ScriptFile defines a single script file
type ScriptFile struct {
	Name     string   // Name is a script name (can be used as an identifier)
	FileName string   // FileName is a full file name of the script
	Kind     FileType // Kind is the script' file type
	Body     string   // Body is the script' body (a.k.a the file content)
}

// ScriptItem defines informations needed to access the big map of ScriptData
type ScriptItem struct {
	Name string   // Name is a script name (can be used as an identifier)
	Kind FileType // Kind is the script' file type
}

// FileType defines types of file. This is used to define what kind of files will we use on our directory
type FileType string

type OurScript struct {
	RawSlice []ScriptFile
	Map      map[FileType]map[string]ScriptFile
}

type ScriptCollectionMap map[FileType]map[string]ScriptFile

var ScriptCollection OurScript

func Initialize(path string, kinds []FileType) (OurScript, error) {
	err := scanDir(path, kinds)
	if err != nil {
		return OurScript{}, err
	}
	buildMap(ScriptCollection.RawSlice)
	return ScriptCollection, nil
}

func scanDir(path string, kinds []FileType) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, vf := range files {
		if !vf.IsDir() {
			name, kind := getScriptNameAndKind(vf.Name())
			if name == "" && kind == "" {
				continue
			} else {
				scriptFile := ScriptFile{
					Name:     name,
					FileName: vf.Name(),
					Kind:     kind,
				}
				body, err := readScriptBody(fmt.Sprintf("%s/%s", path, vf.Name()))
				if err != nil {
					continue
				} else {
					scriptFile.Body = body
					if IsOfKind(kind, kinds) {
						ScriptCollection.RawSlice = append(ScriptCollection.RawSlice, scriptFile)
					}
				}
			}
		} else {
			deeperPath := fmt.Sprintf("%s/%s", path, vf.Name())
			err := scanDir(deeperPath, kinds)
			if err != nil {
				continue
			}
		}
	}
	return err
}

func getScriptNameAndKind(fileName string) (string, FileType) {
	// WARNING! AVOID USING SPLIT BY ".", AS FILE NAME MAY CONTAIN THAT CHARACTER
	length := len(fileName)
	kind, name := "", ""
	breaker := false
	iterator := length - 1
	if length == 0 {
		return "", FileType("")
	}
	for !breaker {
		if fileName[iterator:iterator+1] == "." {
			breaker = true
		} else {
			kind = fileName[iterator:iterator+1] + kind
		}
		if iterator == 0 {
			breaker = true
		}
		iterator--
	}
	if len(kind) == 0 {
		return "", FileType("")
	}
	name = strings.TrimSuffix(fileName, "."+kind)
	kind = strings.ToLower(kind)
	return name, FileType(kind)
}

func readScriptBody(path string) (string, error) {
	rawFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer rawFile.Close()
	rawBody, err := ioutil.ReadAll(rawFile)
	if err != nil {
		return "", err
	}

	return string(rawBody), nil
}

func IsOfKind(kind FileType, kinds []FileType) bool {
	for _, vk := range kinds {
		if strings.ToLower(string(kind)) == strings.ToLower(string(vk)) {
			return true
		}
	}
	return false
}

func buildMap(scriptFiles []ScriptFile) {
	resultMap := make(map[FileType]map[string]ScriptFile)
	for _, vsf := range scriptFiles {
		if _, okrm := resultMap[vsf.Kind]; !okrm {
			resultMap[vsf.Kind] = make(map[string]ScriptFile)
		}
		resultMap[vsf.Kind][vsf.Name] = vsf
	}
	ScriptCollection.Map = resultMap
}

func (o *OurScript) FindScripts(scriptItems []ScriptItem) map[FileType][]ScriptFile {
	result := make(map[FileType][]ScriptFile)
	for _, vsi := range scriptItems {
		if vsi.Kind == "" || vsi.Name == "" {
			continue
		} else {
			if sk, okk := o.Map[vsi.Kind]; okk {
				if sn, okn := sk[vsi.Name]; okn {
					result[vsi.Kind] = append(result[vsi.Kind], sn)
				}
			}
		}
	}
	return result
}

func (o *OurScript) FindBundledScripts(mappedScriptItem map[FileType][]string) map[FileType]string {
	result := make(map[FileType]string)
	cssScript, jsScript := "", ""
	for kmsi, vmsi := range mappedScriptItem {
		switch kmsi {
		case CSS:
			for _, vvmsi := range vmsi {
				if mappedCSS, okcss := o.Map[CSS]; okcss {
					if rawScriptFile, ok := mappedCSS[vvmsi]; ok {
						cssScript += rawScriptFile.Body
					}
				}
			}

		case JS:
			for _, vvmsi := range vmsi {
				if rawScriptFile, ok := o.Map[JS][vvmsi]; ok {
					jsScript += rawScriptFile.Body
				}
			}

		default:
			continue

		}

	}
	cssScript = BuildCSSBundle(cssScript)
	result[CSS] = cssScript
	jsScript = BuildJavascriptBundle(jsScript)
	result[JS] = jsScript
	return result
}
