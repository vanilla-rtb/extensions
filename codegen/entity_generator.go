
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

type FieldEntity struct {
    FieldName    string
    FieldType    string
    FieldTypeIPC string
    IsKey        bool
    IsValue      bool
}

type FieldEntities []FieldEntity

func (s FieldEntities) Len() int {
    return len(s)
}

func (s FieldEntities) Less(i, j int) bool {
    return bool2int8(s[i].IsKey) < bool2int8(s[j].IsKey)
}

func (s FieldEntities) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}



var FuncMap = template.FuncMap{
    "ToUpper": strings.ToUpper,
    "ToLower": strings.ToLower,
    "Title":   strings.Title,
    "NotLast": func(x int, a interface{}) bool {
        return x != reflect.ValueOf(a).Len()-1
    },
    "IsKey": func(storage FieldEntity) bool {
        return storage.IsKey
    },
    "IsIPC": func(storage FieldEntity) bool {
        return storage.FieldType != storage.FieldTypeIPC
    },
    "GetFieldKeys": func(storage FieldEntities) FieldEntities {
        keys := []FieldEntity{}
        for i := range storage {
            if storage[i].IsKey {
                keys = append(keys, storage[i])
            }
        }
        return keys
    },
    "GetFieldValues": func(storage FieldEntities) FieldEntities {
        values := []FieldEntity{}
        for i := range storage {
            if storage[i].IsValue {
                values = append(values, storage[i])
            }
        }
        return values
    },
    "Sort": func(storage FieldEntities) FieldEntities {
        sort.Reverse(storage)
        return storage
    },
    "GetReflectionTypeName" : func(t reflect.Type) string {
         return t.Name()
    },
    "GetReflectionValueName" : func(v reflect.Value) string {
        return v.Type().Name()
    },
    "GetCacheFields" : func(t reflect.Type) FieldEntities {

        v := reflect.ValueOf(reflect.New(t).Elem().Interface())

        cachedFields := make([]FieldEntity, v.NumField())

        for i := 0; i < v.NumField(); i++ {
            field := v.Type().Field(i)
            is_key_str := field.Tag.Get("is_key")
            is_value_str := field.Tag.Get("is_value")
            cachedFields[i] = FieldEntity{
                FieldType:    field.Tag.Get("cpp"),
                FieldTypeIPC: field.Tag.Get("ipc"),
                FieldName:    field.Name,
                IsKey:        len(is_key_str) > 0 && is_key_str == "yes",
                IsValue:      len(is_value_str) > 0 && is_value_str == "yes",
            }
        }
        return cachedFields
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

    cachedFields := make([]FieldEntity, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        field := v.Type().Field(i)
        is_key_str := field.Tag.Get("is_key")
        is_value_str := field.Tag.Get("is_value")
        cachedFields[i] = FieldEntity{
            FieldType:    field.Tag.Get("cpp"),
            FieldTypeIPC: field.Tag.Get("ipc"),
            FieldName:    field.Name,
            IsKey:        len(is_key_str) > 0 && is_key_str == "yes",
            IsValue:      len(is_value_str) > 0 && is_value_str == "yes",
        }
        fmt.Println(cachedFields[i].FieldName, "=>", cachedFields[i].FieldType, "=>", cachedFields[i].FieldTypeIPC, "=>", cachedFields[i].IsKey, "=>", cachedFields[i].IsValue)
    }

    err := g.Template.Execute(f, struct {
        Timestamp    time.Time
        Matchername  string
        CachedFields []FieldEntity
    }{
        Timestamp:    time.Now(),
        Matchername:  g.GeneratedBasicName,
        CachedFields: cachedFields,
    })
    return err
}
