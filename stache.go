package main

import (
	"flag"
	"fmt"
	"github.com/hoisie/mustache"
	"io/ioutil"
	"os"
	"strings"
)

var file string
var outFile string

func init() {
	version := false

	flag.BoolVar(&version, "version", false, "Print the version")
	flag.BoolVar(&version, "v", false, "Print the version")
	flag.StringVar(&file, "f", "", "A template file")
	flag.StringVar(&outFile, "o", "", "Save output to a file")
	flag.StringVar(&file, "file", "", "A template file")
	flag.Parse()

	if version {
		fmt.Println("Version:", VERSION)
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

func renderTemplate(data string) string {
	tmplS, _ := mustache.ParseString(data)
	return tmplS.Render(getContext(), &tmplS)
}

func main() {
	info, _ := os.Stdin.Stat()
	var templateString string

	if info.Size() > 0 {
		// We have a unix pipe
		stdinBytes, _ := ioutil.ReadAll(os.Stdin)
		templateString = string(stdinBytes)
	} else if file != "" {
		// We got a filename as the first argument
		fileData, _ := ioutil.ReadFile(file)
		templateString = string(fileData)
	} else {
		flag.Usage()
		os.Exit(1)
	}

	output := renderTemplate(templateString)

	if outFile == "" {
		fmt.Print(output)
	} else {
		ioutil.WriteFile(outFile, []byte(output), 0644)
	}
}
