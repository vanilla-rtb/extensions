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

    "log"
    //"reflect"

    "github.com/jessevdk/go-flags"
    "github.com/vanilla-rtb/extensions/codegen"
    "reflect"
)

func die(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

type Options struct {
    InputTemplate flags.Filename `short:"i" long:"input-template" description:"InputTemplate file" default:"-"`
    OutputDir     flags.Filename `short:"o" long:"output-dir" description:"OutputDir file" default:"-"`
    GeneratorType func(string) `short:"g" long:"generator-type" description:"GeneratorType" default:"cache"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

var generatorTypes = map[string]reflect.Type {
    "cmake"  : reflect.TypeOf(codegen.CmakeGenerator{}),
    "app"    : reflect.TypeOf(codegen.AppGenerator{}),
    "cache"  : reflect.TypeOf(codegen.CacheGenerator{}),
}

var generatorType reflect.Type

func main() {

    options.GeneratorType = func(gen string) {
      generatorType = generatorTypes[gen]
    }
    _, err := parser.Parse()
    die(err)

    templateContent, err := ioutil.ReadFile(string(options.InputTemplate))
    die(err)

    var Template= template.Must(template.New("").Funcs(codegen.FuncMap).Parse(string(templateContent)))
    if generatorType == reflect.TypeOf(codegen.CacheGenerator{}) {
        err := codegen.NewCacheGenerator(options.OutputDir, Template).Execute(nil)
        die(err)
    } else if generatorType == reflect.TypeOf(codegen.AppGenerator{}) {
        outFileName := strings.Join([]string{string(options.OutputDir), "bidder", ".cpp"}, "/")
        f, err := os.Create(outFileName)
        die(err)
        err = codegen.NewAppGenerator(Template).Execute(f)
        die(err)
    } else if generatorType == reflect.TypeOf(codegen.CmakeGenerator{}) {
        outFileName := strings.Join([]string{string(options.OutputDir), "CMakeFile.txt"}, "/")
        f, err := os.Create(outFileName)
        die(err)
        err = codegen.NewCmakeGenerator("bidder",Template).Execute(f)
        die(err)
    }


}
