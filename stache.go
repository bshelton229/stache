package main

import (
	"flag"
	"fmt"
	"github.com/hoisie/mustache"
	"io/ioutil"
	"os"
	"strings"
)

func getContext() map[string]string {
	context := make(map[string]string)

	for _, el := range os.Environ() {
		a := strings.Split(el, "=")
		context[a[0]] = a[1]
	}
	return context
}

func main() {
	// TODO: Add some kind, any kind, of error handling and alerting
	tmplData, _ := ioutil.ReadFile(flag.Arg(0))
	tmplS, _ := mustache.ParseString(string(tmplData))
	fmt.Println(tmplS.Render(getContext(), &tmplS))
}
