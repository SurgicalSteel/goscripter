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

And oh, we do have dependency on [stretchr/testify/assert](https://github.com/stretchr/testify/assert) for unit testing purpose. So, make sure you get [stretchr/testify/assert](https://github.com/stretchr/testify/assert) first before running tests on goscripter.

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
    staticScripts, err := goscripter.Initialize("files/scripts", kinds)
    ...
}
```

After we've initialized as shown above, we can use staticScripts as our static script collection. We can get scripts as we need.
