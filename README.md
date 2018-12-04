# goscripter
An utility to manage static scripts for your go-rendered pages.
We scan recursively through your static scripts folder, collect them by your reference. And you can access them everytime you need by calling provided functions in an easy way.

## Table of Contents

* [Installation](#installation)
* [Usage](#usage)
* [Documentation](#documentation)
* [Contribution](#contribution)
* [Issue](#issue)

## Installation

To be able to use goscripter, all you need is to run

    $ go get github.com/SurgicalSteel/goscripter

(optional) To run unit tests:

    $ cd $GOPATH/src/github.com/SurgicalSteel/goscripter
    $ go test -v -cover

And oh, we do have dependency on [stretchr/testify](https://github.com/stretchr/testify) for unit testing purpose. So, make sure you get [stretchr/testify](https://github.com/stretchr/testify) first before running tests on goscripter.

## Usage
First, we need to do initialization. In this process, what goscripter does is to recursively scan the given folder path, and collect static scripts which match your preferences (in this case, file type)
```go
package yourpackage
import(
    ...
    "github.com/SurgicalSteel/goscripter"
    ...
)
func main(){
    ...
    kinds := []goscripter.FileType{goscripter.CSS, goscripter.JS, goscripter.JSON}
    OurStaticScripts, err := goscripter.Initialize("files/scripts", kinds)
    ...
}
```

After we've initialized as shown above, we can use staticScripts as our static script collection. We can get scripts as we need.

#### Getting Bundled Script

To get the desired and bundled script, we need to create a map as a specification which defines what kind of file that we want, and what is the file name (without file type).

For example, you have two js files (base.js and action.js), and you want to get them bundled (wrapped with <script> tag) so that it is ready to use.
Then you need to specify it, and get the bundled script (as specified) like this :

```go
package yourpackage

import(
    "net/http"
    "github.com/SurgicalSteel/goscripter"
    "text/template"
)

func Handle404PageRender(w http.ResponseWriter, r *http.Request) {
    //data map which will be passed for rendering purpose
    data := make(map[string]interface{})

    // specify bundleMap (what kind of scripts we need, and their names)    
    bundleMap := make(map[goscripter.FileType][]string)

    // in this case, we need two js files : base.js and action.js
    bundleMap[goscripter.JS] = []string{"base","action"}

    // get the bundled scripts
    scripts := OurStaticScripts.FindBundledScripts(bundleMap)

    // check if the requested bundle of script is present
    if javascript, okjs := scripts[goscripter.JS]; okjs {
        // if exist, add it to the data map (for rendering)
		data["javascript"] = javascript
	}

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := template.Must(template.New("404").ParseFiles("files/template/404.html"))
	tmpl.ExecuteTemplate(w, "404", data)
}    
```

To Include the script on the template, just use it like this :

```html
{{ define "404" }}
<html>
    <head>
        <title>Not found</title>
        {{ .javascript }}
    </head>
    <body>
        <h1>Oops, the page you requested was not found!</h1>
    </body>
</html>
{{ end }}
```

#### Getting A Script
To get a single script, you need to define it using goscripter's ScriptItem as our default Type to define script specification. After the ScriptItem specification has been defined, you can use FindAScript(scriptItem) to find a script file you need.

example :
```go
    package yourpackage

    import(
        ...
        "github.com/SurgicalSteel/goscripter"
        ...
    )
    func yourFunction(){
        ...
        baseJSSpec := goscripter.ScriptItem{
            Name : "base",
            Kind : goscripter.JS,
        }
        baseJSScript := OurStaticScripts.FindAScript(baseJSSpec) // this returns a goscripter ScriptFile
        // to get the script body
        baseJSScriptBody := baseJSScript.Body
        ...
    }
```


#### Getting the JSON String
To get the JSON string provided in the collected scripts (after initialization), in this case OurStaticScripts, all we need to do is create specification, and then call FindJSON on it.

Take a look at this :
```go
...
    studentDataScriptItem := goscripter.ScriptItem{
        Name : "StudentData",
        Kind : goscripter.JSON,
    }
    studentDataJSON := OurStaticScripts.FindJSON(studentDataScriptItem)
...
```

We also provide a simple utility to validate your JSON string. To use it, just call ValidateJSON().
Example :
```go
...
    isValidStudentData := goscripter.ValidateJSON(studentDataJSON) //returns boolean (true or false)
...
```

## Documentation
We use standard godoc as our code documentation tool. To view it, please follow these steps :
1. Open your terminal, head to this cloned repo (SurgicalSteel/goscripter)
2. run `godoc -http=:6060` (this will trigger godoc at port 6060)
3. Open your browser, and hit `http://127.0.0.1:6060/pkg/github.com/SurgicalSteel/goscripter/`

## Contribution
This repository is open for contribution. To make a contribution, please do following steps :
1. Fork this repository
2. Create a new branch (from master branch) for your feature
3. Ensure your changes have test covered, and code documentation
4. Create a pull request (don't forget to attach a clear description, tags are optional)
5. Within a week, we'll review your changes

## Issue
If you found some issues, feel free to submit it [here](https://github.com/SurgicalSteel/goscripter/issues/new)
