// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates matchers It can be invoked by running
// go generate
package main

import (
    "log"
    "os"
    "strings"
    "path/filepath"
    "text/template"
    "time"
    "reflect"
    "fmt"
    "io/ioutil"
    "github.com/jessevdk/go-flags"
)

//TODO: use reflection to create a list of fields in the template, How do we pass this struct to generator ???
type  Domain struct {
    name string   `cpp:"std::string"  ipc:"char_string""`
    dom_id uint32 `cpp:"uint32_t" ipc:"uint32_t"`
}

type CacheStorage struct {
    FieldName    string
    FieldType    string
	FieldTypeIPC string
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"Title" : strings.Title,
}

type Options struct {
	Input  flags.Filename `short:"i" long:"input" description:"Input file" default:"-"`
	Output flags.Filename `short:"o" long:"output" description:"Output file" default:"-"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)


func main() {

    _, err := parser.Parse()
    die(err)

	templateContent, err := ioutil.ReadFile(string(options.Input))
    die(err)
	var matcherTemplate = template.Must(template.New("").Funcs(funcMap).Parse(string(templateContent)))

	domain := Domain{"",0}
	v := reflect.ValueOf(domain)

	cachedValues := make([]CacheStorage, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		cachedValues[i].FieldType = field.Tag.Get("cpp")
		cachedValues[i].FieldTypeIPC = field.Tag.Get("ipc")
		cachedValues[i].FieldName = field.Name
		fmt.Println(cachedValues[i].FieldName , "=>" , cachedValues[i].FieldType, "=>" , cachedValues[i].FieldTypeIPC)
	}

	outFileName := string(options.Output)
	generatedMatcherName  := strings.TrimSuffix(outFileName, filepath.Ext(outFileName))


	f, err := os.Create(outFileName)
	die(err)
	defer f.Close()

	matcherTemplate.Execute(f, struct {
		Timestamp   time.Time
		Matchername string
		CachedValues []CacheStorage
	}{
		Timestamp:   time.Now(),
		Matchername: generatedMatcherName,
		CachedValues: cachedValues,
	})

}
