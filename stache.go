package main

import (
	"bufio"
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
		fmt.Println("Version:", Version)
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
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			templateString += scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("FAIL")
			os.Exit(1)
		}
	} else if flag.Arg(0) != "" {
		// We got a filename as the first argument
		fileData, _ := ioutil.ReadFile(flag.Arg(0))
		templateString = string(fileData)
	} else {
		// TODO: Print usage
		fmt.Println("Print usage here")
		os.Exit(1)
	}

	// Render the template if we haven't exited
	fmt.Println(renderTemplate(templateString))
}
