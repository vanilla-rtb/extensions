// The following directive is necessary to make the package coherent:

// build ignore

// This program generates matchers It can be invoked by running
// go generate
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    _ "path/filepath"
    "reflect"
    "sort"
    "strings"
    "text/template"
    "time"

    "github.com/jessevdk/go-flags"
)

//TODO: use reflection to create a list of fields in the template, How do we pass this struct to generator ???
type Domain struct {
    name      string `cpp:"std::string"  ipc:"char_string" is_key:"yes"`
    domain_id uint32 `cpp:"uint32_t" ipc:"uint32_t"`
}

type CacheEntity struct {
    FieldName    string
    FieldType    string
    FieldTypeIPC string
    IsKey        bool
}

type CacheEntities []CacheEntity

func (s CacheEntities) Len() int {
	return len(s)
}

func (s CacheEntities) Less(i, j int) bool {
	return bool2int8(s[i].IsKey) < bool2int8(s[j].IsKey)
}

func (s CacheEntities) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func bool2int8 ( b bool ) int8 {
	var bits int8 = 0
	if b {
		bits = 1
	}
	return bits
}

func die(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

var funcMap = template.FuncMap{
    "ToUpper": strings.ToUpper,
    "ToLower": strings.ToLower,
    "Title":   strings.Title,
    "NotLast": func(x int, a interface{}) bool {
        return x != reflect.ValueOf(a).Len()-1
    },
	"IsKey": func(storage CacheEntity) bool {
		return storage.IsKey
	},
	"IsIPC" : func(storage CacheEntity) bool {
	    return storage.FieldType != storage.FieldTypeIPC
	},
	"GetCacheKeys" : func(storage CacheEntities) CacheEntities {
		keys := []CacheEntity{}
		for i := range storage {
			if(storage[i].IsKey) {
				keys = append(keys, storage[i])
			}
		}
		return keys
	},
	"Sort" : func(storage CacheEntities) CacheEntities {
		sort.Reverse(storage)
		return storage
	},
}

type Options struct {
    InputTemplate flags.Filename `short:"i" long:"input-template" description:"InputTemplate file" default:"-"`
    OutputDir     flags.Filename `short:"o" long:"output-dir" description:"OutputDir file" default:"-"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {

    _, err := parser.Parse()
    die(err)

    templateContent, err := ioutil.ReadFile(string(options.InputTemplate))
    die(err)
    var matcherTemplate = template.Must(template.New("").Funcs(funcMap).Parse(string(templateContent)))

    domain := Domain{"", 0}
    v := reflect.ValueOf(domain)

    cachedValues := make([]CacheEntity, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        field := v.Type().Field(i)
        cachedValues[i].FieldType = field.Tag.Get("cpp")
        cachedValues[i].FieldTypeIPC = field.Tag.Get("ipc")
        cachedValues[i].FieldName = field.Name
        is_key_str := field.Tag.Get("is_key")
		cachedValues[i].IsKey = len(is_key_str) > 0 && is_key_str=="yes"
        fmt.Println(cachedValues[i].FieldName, "=>", cachedValues[i].FieldType, "=>", cachedValues[i].FieldTypeIPC, "=>", cachedValues[i].IsKey)
    }

	generatedMatcherName := strings.ToLower(v.Type().Name())
    outFileName := strings.Join([]string{string(options.OutputDir) , strings.Join([]string{generatedMatcherName,".hpp"},"")} , "/")
    //generatedMatcherName := strings.TrimSuffix(outFileName, filepath.Ext(outFileName))

    f, err := os.Create(outFileName)
    die(err)
    defer f.Close()

    err = matcherTemplate.Execute(f, struct {
        Timestamp    time.Time
        Matchername  string
        CachedValues []CacheEntity
    }{
        Timestamp:    time.Now(),
        Matchername:  generatedMatcherName,
        CachedValues: cachedValues,
    })
    die(err)
}
