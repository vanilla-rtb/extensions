
package codegen

import (
    "os"
    "text/template"
    "time"
)


type CmakeGenerator struct {
    Template   *template.Template
    Executable string
}

func NewCmakeGenerator(exe string, tmpl *template.Template) *CmakeGenerator {
    return &CmakeGenerator{
        Template:  tmpl,
        Executable: exe,
    }
}

func (g *CmakeGenerator) Execute(f *os.File) error {

    err := g.Template.Execute(f, struct {
        Timestamp    time.Time
        Executable  string
    }{
        Timestamp:    time.Now(),
        Executable:   g.Executable,
    })

    return err
}
