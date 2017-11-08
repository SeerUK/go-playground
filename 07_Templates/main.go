package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"reflect"
	"strings"
	"time"
)

var rawData = `{
	"type": "article",
	"data": {
		"foo": "This is a foo",
		"bar": {
			"code": 500,
			"message": "This is a bar"
		},
		"baz": 123,
		"qux": true
	}
}`

type ContentData map[string]interface{}

type ContentEntry struct {
	Type string      `json:"type"`
	Data ContentData `json:"data"`
}

func TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"IsEqual": func(a, b interface{}) bool {
			return reflect.DeepEqual(a, b)
		},
	}
}

func main() {
	start := time.Now()

	entries := []ContentEntry{}
	entry := ContentEntry{}

	err := json.Unmarshal([]byte(rawData), &entry)
	if err != nil {
		log.Fatal(err)
	}

	entries = append(entries, entry)
	entries = append(entries, entry)

	tmpl := strings.TrimSpace(`
{{if IsEqual .foo "This is a foo"}}Foo {{.foo}}{{end}}
{{if .bar}}Bar {{.bar.message}} ({{.bar.code}}){{end}}
	`)

	b := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("article").Funcs(TemplateFuncs()).Parse(tmpl))

	for _, e := range entries {
		for _, et := range t.Templates() {
			if et.Name() != e.Type {
				continue
			}

			err := et.Execute(b, entry.Data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println(b.String())
	fmt.Println(time.Since(start))
}
