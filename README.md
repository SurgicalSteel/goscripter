# goscripter
An utility to manage static scripts for your go-rendered pages
We scan iteratively through your static scripts folder, collect them by your reference. And you can access them everytime you need by calling provided functions.

## Table of Contents

* [Installation](#installation)
* [Usage](#Usage)

## Installation

To be able to use goscripter, all you need is to run

    $ go get github.com/SurgicalSteel/goscripter

(optional) To run unit tests:

    $ cd $GOPATH/src/github.com/SurgicalSteel/goscripter
    $ go test -v -cover

And oh, we do have dependency on [stretchr/testify](https://github.com/stretchr/testify) for unit testing purpose. So, make sure you get [stretchr/testify](https://github.com/stretchr/testify) first before running tests on goscripter.

## Usage
First, we need to do initialization. In this process, what goscripter does is to recursively scan the given folder path, and collect static scripts which match your preferences (in this case, file type)
```
package yourpackage
import(
    ...
    "github.com/SurgicalSteel/goscripter"
    ...
)
func main(){
    ...
    kinds := []goscripter.FileType{goscripter.CSS, goscripter.JS}
    OurStaticScripts, err := goscripter.Initialize("files/scripts", kinds)
    ...
}
```

After we've initialized as shown above, we can use staticScripts as our static script collection. We can get scripts as we need.

To get the desired and bundled script, we need to create a map as a specification which defines what kind of file that we want, and what is the file name (without file type).

For example, you have two js files (base.js and action.js), and you want to get them bundled (wrapped with <script> tag) so that it is ready to use.
Then you need to specify it, and get the bundled script (as specified) like this :

```
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

```
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
