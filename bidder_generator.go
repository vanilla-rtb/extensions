// The following directive is necessary to make the package coherent:

// build ignore

// This program generates matchers It can be invoked by running
// go generate
package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

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
	GeneratorType func(string)   `short:"g" long:"generator-type" description:"GeneratorType" default:"matchers"`
	Executable    flags.Filename `short:"e" long:"executable" description:"Executable name" default:"bidder"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

type GenerateFunc func(*Options, *template.Template) error

var generatorTypes = map[string]GenerateFunc{
	"cmake": func(o *Options, t *template.Template) error {
		outFileName := strings.Join([]string{string(options.OutputDir), "CMakeLists.txt"}, "/")
		f, err := os.Create(outFileName)
		if err != nil {
			return err
		}
		return codegen.NewCmakeGenerator("bidder", t).Execute(f)
	},
	"app": func(o *Options, t *template.Template) error {
		outFileName := strings.Join([]string{string(options.OutputDir), strings.Join([]string{string(options.Executable), ".cpp"},"")}, "/")
		f, err := os.Create(outFileName)
		if err != nil {
			return err
		}
		return codegen.NewAppGenerator(t).Execute(f)
	},
	"matchers": func(o *Options, t *template.Template) error {
		return codegen.NewCacheGenerator(options.OutputDir, t).Execute(nil)
	},
}

func main() {
	var generateFunc GenerateFunc
	options.GeneratorType = func(gen string) {
		var ok bool
		if generateFunc, ok = generatorTypes[gen]; !ok {
			log.Fatalf("Unsupported generator-type %s\n", gen)
		}
	}
	_, err := parser.Parse()
	die(err)
	templateContent, err := ioutil.ReadFile(string(options.InputTemplate))
	die(err)

	var tmpl = template.Must(template.New("").Funcs(codegen.FuncMap).Parse(string(templateContent)))
	err = generateFunc(&options, tmpl)
	die(err)
}
