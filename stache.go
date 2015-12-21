package main

import (
	"flag"
	"fmt"
	"github.com/hoisie/mustache"
	"io/ioutil"
	"os"
	"strings"
)

func init() {
	version := false

	flag.BoolVar(&version, "version", false, "Print the version")
	flag.BoolVar(&version, "v", false, "Print the version")
	flag.Parse()

	if version {
		fmt.Println("Welcome to stache version", Version)
		os.Exit(0)
	}
}

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
