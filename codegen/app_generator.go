
package codegen

import (
    "os"
    "text/template"
)


type AppGenerator struct {
    Template  *template.Template
}

func NewAppGenerator(tmpl *template.Template) *AppGenerator {
    return &AppGenerator{
        Template:  tmpl,
    }
}

func (g *AppGenerator) Execute(f *os.File) error {
    var err error
    return err
}
