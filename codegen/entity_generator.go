
package codegen

import (
    "fmt"
    "os"
    "reflect"
    "sort"
    "strings"
    "text/template"
    "time"
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



var FuncMap = template.FuncMap{
    "ToUpper": strings.ToUpper,
    "ToLower": strings.ToLower,
    "Title":   strings.Title,
    "NotLast": func(x int, a interface{}) bool {
        return x != reflect.ValueOf(a).Len()-1
    },
    "IsKey": func(storage CacheEntity) bool {
        return storage.IsKey
    },
    "IsIPC": func(storage CacheEntity) bool {
        return storage.FieldType != storage.FieldTypeIPC
    },
    "GetCacheKeys": func(storage CacheEntities) CacheEntities {
        keys := []CacheEntity{}
        for i := range storage {
            if storage[i].IsKey {
                keys = append(keys, storage[i])
            }
        }
        return keys
    },
    "Sort": func(storage CacheEntities) CacheEntities {
        sort.Reverse(storage)
        return storage
    },
}

type EntityGenerator struct {
    Entity             interface{}
    Template           *template.Template
    GeneratedBasicName string
}

func NewEntityGenerator(entity interface{}, tmpl *template.Template) *EntityGenerator {
    return &EntityGenerator{
        Entity:             entity,
        Template:           tmpl,
        GeneratedBasicName: reflect.ValueOf(entity).Type().Name(),
    }
}

func (g *EntityGenerator) Execute(f *os.File) error {

    v := reflect.ValueOf(g.Entity)

    cachedValues := make([]CacheEntity, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        field := v.Type().Field(i)
        is_key_str := field.Tag.Get("is_key")
        cachedValues[i] = CacheEntity{
            FieldType:    field.Tag.Get("cpp"),
            FieldTypeIPC: field.Tag.Get("ipc"),
            FieldName:    field.Name,
            IsKey:        len(is_key_str) > 0 && is_key_str == "yes",
        }
        fmt.Println(cachedValues[i].FieldName, "=>", cachedValues[i].FieldType, "=>", cachedValues[i].FieldTypeIPC, "=>", cachedValues[i].IsKey)
    }

    err := g.Template.Execute(f, struct {
        Timestamp    time.Time
        Matchername  string
        CachedValues []CacheEntity
    }{
        Timestamp:    time.Now(),
        Matchername:  g.GeneratedBasicName,
        CachedValues: cachedValues,
    })
    return err
}
