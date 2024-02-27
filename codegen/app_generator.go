package codegen

import (
	"fmt"
	"github.com/vanilla-rtb/extensions/stubs"
	"os"
	"path"
	"reflect"
	"strings"
	"text/template"
	"time"
)

type AppGenerator struct {
	OutputDir      string
	MatcherDir     string
	AppName        string
	TargetingModel string
	BuildType      string
	Template       *template.Template
}

func (appGen *AppGenerator) Execute() error {
	outFileName := path.Join(appGen.OutputDir, strings.ToLower(appGen.AppName)+".cpp")
	f, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	var cachedEntities []reflect.Type
	var ok bool
	if cachedEntities, ok = stubs.Targetings[appGen.TargetingModel]; !ok {
		return fmt.Errorf("Failed to find TargetingModel %s", appGen.TargetingModel)
	}

	err = appGen.Template.Execute(f, struct {
		Timestamp      time.Time
		CachedEntities []reflect.Type
		BidderName     string
		BuildType      string
	}{
		Timestamp:      time.Now(),
		CachedEntities: cachedEntities,
		BidderName:     appGen.AppName,
		BuildType:      appGen.BuildType,
	})
	return err
}
