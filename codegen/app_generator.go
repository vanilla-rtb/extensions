
package codegen

import (
    "os"
    "text/template"
    "stubs"
    "reflect"
    "strings"
    "github.com/jessevdk/go-flags"
    "time"
)


type AppGenerator struct {
    OutputDir  flags.Filename
    MatcherDir  string
    AppName string
    TargetingModel string
    BuildType string
    Template  *template.Template
}

func NewAppGenerator(dirName flags.Filename, appName string, model string, buildType string, tmpl *template.Template) *AppGenerator {
    return &AppGenerator{
        OutputDir:    dirName,
        AppName: appName,
        TargetingModel: model,
        BuildType: buildType,
        Template:     tmpl,
    }
}

func (appGen *AppGenerator) Execute(f *os.File) error {
    outFileName := strings.Join([]string{string(appGen.OutputDir), strings.Join([]string{strings.ToLower(appGen.AppName), ".cpp"}, "")}, "/")
    f, err := os.Create(outFileName)
    if err != nil {
        return err
    }
    defer f.Close()

    var cachedEntities []reflect.Type
    var ok bool
    if cachedEntities , ok = stubs.Targetings[appGen.TargetingModel]; !ok {
        return nil //TODO: return the actual error
    }

    err = appGen.Template.Execute(f, struct {
        Timestamp    time.Time
        CachedEntities []reflect.Type
        BidderName string
        BuildType string
    }{
        Timestamp:    time.Now(),
        CachedEntities: cachedEntities,
        BidderName: appGen.AppName,
        BuildType: appGen.BuildType,
    })
    return err
}
