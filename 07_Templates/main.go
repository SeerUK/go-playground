package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"reflect"
	"strings"
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

type ContentEntry struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func AssertJSONObject(o interface{}) (map[string]interface{}, bool) {
	r, ok := o.(map[string]interface{})

	return r, ok
}

func main() {
	entry := ContentEntry{}

	err := json.Unmarshal([]byte(rawData), &entry)
	if err != nil {
		log.Fatal(err)
	}

	funcs := template.FuncMap{
		"IsEqual": func(a, b interface{}) bool {
			return reflect.DeepEqual(a, b)
		},
	}

	if a, ok := AssertJSONObject(entry.Data); ok {
		if b, ok := AssertJSONObject(a["bar"]); ok {
			fmt.Println(b["message"])
		}
	}

	tmpl := strings.TrimSpace(`
{{if IsEqual .foo "This is a foo"}}Foo {{.foo}}{{end}}
{{if .bar}}Bar {{.bar.message}} ({{.bar.code}}){{end}}
	`)

	b := bytes.NewBuffer([]byte{})

	t := template.Must(template.New("test").Funcs(funcs).Parse(tmpl))
	t.Execute(b, entry)

	fmt.Println(b.String())
}
