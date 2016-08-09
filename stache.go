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

func parseContextStrings(cs []string) map[string]string {
	context := make(map[string]string)

	for _, el := range cs {
		a := strings.SplitN(el, "=", 2)
		context[a[0]] = a[1]
	}
	return context
}

func renderTemplate(data string, context map[string]string) string {
	return mustache.Render(data, context)
}

func main() {
	stat, _ := os.Stdin.Stat()
	var templateString string

	if file != "" {
		// We got a filename as the first argument
		fileData, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		templateString = string(fileData)

	} else if (stat.Mode() & os.ModeCharDevice) == 0 {
		// We have a unix pipe
		stdinBytes, _ := ioutil.ReadAll(os.Stdin)
		templateString = string(stdinBytes)
	} else {
		flag.Usage()
		os.Exit(1)
	}

	output := renderTemplate(templateString, parseContextStrings(os.Environ()))

	if outFile == "" {
		fmt.Print(output)
	} else {
		ioutil.WriteFile(outFile, []byte(output), 0644)
	}
}
