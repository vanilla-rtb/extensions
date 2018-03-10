package codegen

import (
    "os"
    "path"
    "reflect"
    "strings"
    "stubs"
    "text/template"
)

type CacheGenerator struct {
    OutputDir string
    Template  *template.Template
}

func (cacheGen *CacheGenerator) Execute() error {
    for _, value := range stubs.TypeRegistry {
        entityGen := NewEntityGenerator(reflect.New(value).Elem().Interface(), cacheGen.Template)

        outFileName := path.Join(cacheGen.OutputDir, strings.ToLower(entityGen.GeneratedBasicName)+".hpp")

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
