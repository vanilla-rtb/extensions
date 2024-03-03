// The following directive is necessary to make the package coherent:

// build ignore

// This program generates matchers It can be invoked by running
// go generate
package main

import (
	"log"
	"os"
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
	Executable    string         `short:"e" long:"executable" description:"Executable name" default:"bidder"`
	TargetingName string         `short:"T" long:"targeting-name" description:"Directory for header files" default:"-"`
	BuildType     string         `short:"B" long:"build_type" description:"Directory for header files" default:"APP"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

type GenerateFunc func(*template.Template) error

var generatorTypes = map[string]GenerateFunc{
	"cmake": func(t *template.Template) error {
		return (&codegen.CmakeGenerator{
			Template:   t,
			Executable: options.Executable,
			OutpuDir:   string(options.OutputDir),
		}).Execute()
	},
	"app": func(t *template.Template) error {
		return (&codegen.AppGenerator{
			OutputDir:      string(options.OutputDir),
			AppName:        options.Executable,
			TargetingModel: options.TargetingName,
			BuildType:      options.BuildType,
			Template:       t,
		}).Execute()
	},
	"matchers": func(t *template.Template) error {
		return (&codegen.CacheGenerator{
			OutputDir: string(options.OutputDir),
			Template:  t,
		}).Execute()
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
	templateContent, err := os.ReadFile(string(options.InputTemplate))
	die(err)

	var tmpl = template.Must(template.New("").Funcs(codegen.FuncMap).Parse(string(templateContent)))

	err = generateFunc(tmpl)
	die(err)
}
