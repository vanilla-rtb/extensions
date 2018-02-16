// The following directive is necessary to make the package coherent:

// build ignore

// This program generates matchers It can be invoked by running
// go generate
package main

import (
    "io/ioutil"
    "os"
    "strings"
    "text/template"

    "github.com/jessevdk/go-flags"
    "github.com/vanilla-rtb/extensions/codegen"
    "log"
    "reflect"
    "stubs"
)


func die(err error) {
    if err != nil {
        log.Fatal(err)
    }
}


type Options struct {
    InputTemplate flags.Filename `short:"i" long:"input-template" description:"InputTemplate file" default:"-"`
    OutputDir     flags.Filename `short:"o" long:"output-dir" description:"OutputDir file" default:"-"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {

    _, err := parser.Parse()
    die(err)

    templateContent, err := ioutil.ReadFile(string(options.InputTemplate))
    die(err)

    var matcherTemplate = template.Must(template.New("").Funcs(codegen.FuncMap).Parse(string(templateContent)))
    for _,value  := range stubs.TypeRegistry {
        gen := codegen.NewCodeGenerator(reflect.New(value).Elem().Interface(), matcherTemplate, )

        outFileName := strings.Join([]string{string(options.OutputDir), strings.Join([]string{strings.ToLower(gen.GeneratedBasicName), ".hpp"}, "")}, "/")

        f, err := os.Create(outFileName)
        die(err)
        defer f.Close()

        //Generate code for each stub
        err = gen.Execute(f)
        die(err)
    }
}
