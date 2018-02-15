/*
    err := CodeGenerator{
        entity:             Domain{"", 0},
        templ:              matcherTemplate,
        generatedBasicName: generatedMatcherName,
        fileHandler:        f,
    }.generate()

    die(err)
 */

package codegen

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
	"text/template"
	"time"
	"os"
)

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

type CodeGenerator struct {
	entity             interface{}
	templ              *template.Template
	generatedBasicName string
	fileHandler *os.File
}

func (g *CodeGenerator) generate() error {

	v := reflect.ValueOf(g.entity)

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

	err := g.templ.Execute(g.fileHandler, struct {
		Timestamp    time.Time
		Matchername  string
		CachedValues []CacheEntity
	}{
		Timestamp:    time.Now(),
		Matchername:  g.generatedBasicName,
		CachedValues: cachedValues,
	})
	return err
}