package codegen

import (
    "os"
    "path"
    "text/template"
    "time"
)

type CmakeGenerator struct {
    Template   *template.Template
    Executable string
    OutpuDir   string
}

func (g *CmakeGenerator) Execute() error {
    f, err := os.Create(path.Join(g.OutpuDir, "CMakeLists.txt"))
    if err != nil {
        return err
    }
    defer f.Close()
    return g.Template.Execute(f, struct {
        Timestamp  time.Time
        Executable string
    }{
        Timestamp:  time.Now(),
        Executable: g.Executable,
    })
}
