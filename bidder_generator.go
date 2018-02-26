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
)

func die(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

type Options struct {
    InputTemplate flags.Filename `short:"i" long:"input-template" description:"InputTemplate file" default:"-"`
    OutputDir     flags.Filename `short:"o" long:"output-dir" description:"OutputDir file" default:"-"`
    GeneratorType flags.Filename `short:"g" long:"generator-type" description:"GeneratorType" default:"cache"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

//var generatorTypes = map[string]reflect.Type {
//    "cmake"  : reflect.ValueOf(codegen.CmakeGenerator{}).Type(),
//    "app"    : reflect.ValueOf(codegen.AppGenerator{}).Type(),
//    "cache"  : reflect.ValueOf(codegen.CacheGenerator{}).Type(),
//}

func main() {

    _, err := parser.Parse()
    die(err)

    templateContent, err := ioutil.ReadFile(string(options.InputTemplate))
    die(err)

    var Template= template.Must(template.New("").Funcs(codegen.FuncMap).Parse(string(templateContent)))
    if options.GeneratorType == "cache" {
        err := codegen.NewCacheGenerator(options.OutputDir, Template).Execute(nil)
        die(err)
    } else if options.GeneratorType == "app" {
        outFileName := strings.Join([]string{string(options.OutputDir), "bidder", ".cpp"}, "/")
        f, err := os.Create(outFileName)
        die(err)
        err = codegen.NewAppGenerator(Template).Execute(f)
        die(err)
    } else if options.GeneratorType == "cmake" {
        outFileName := strings.Join([]string{string(options.OutputDir), "CMakeFile.txt"}, "/")
        f, err := os.Create(outFileName)
        die(err)
        err = codegen.NewCmakeGenerator("bidder",Template).Execute(f)
        die(err)
    }

    //for _, value := range stubs.TypeRegistry {
    //    gen := codegen.NewEntityGenerator(reflect.New(value).Elem().Interface(), matcherTemplate)
	//
	//   outFileName := strings.Join([]string{string(options.OutputDir), strings.Join([]string{strings.ToLower(gen.GeneratedBasicName), ".hpp"}, "")}, "/")
	//
    //    f, err := os.Create(outFileName)
    //    die(err)
    //    defer f.Close()
	//
    //    //Generate code for each stub
    //    err = gen.Execute(f)
    //    die(err)
    //}
}
