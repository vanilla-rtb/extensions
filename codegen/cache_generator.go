package codegen

import (
	"reflect"
	"os"
	"text/template"
	"stubs"
	"strings"
	"github.com/jessevdk/go-flags"
)

type CacheGenerator struct {
	OutputDir          flags.Filename
	Template           *template.Template
}

func NewCacheGenerator(dirName flags.Filename, tmpl *template.Template) *CacheGenerator {
	return &CacheGenerator{
		OutputDir:    dirName,
		Template:     tmpl,
	}
}

func (cacheGen *CacheGenerator) Execute(_ *os.File) error {
	for _, value := range stubs.TypeRegistry {
		entityGen := NewEntityGenerator(reflect.New(value).Elem().Interface(), cacheGen.Template)

		outFileName := strings.Join([]string{string(cacheGen.OutputDir), strings.Join([]string{strings.ToLower(entityGen.GeneratedBasicName), ".hpp"}, "")}, "/")

		f, err := os.Create(outFileName)
		if err != nil {
			return err
		}
		defer f.Close()

		//Generate code for each stub
		err = entityGen.Execute(f)
		if err != nil {
			return err
		}
	}
    return nil
}

