package test

import (
	"encoding/json"
	"os"
	"testing"
	"text/template"
)

const (
	// 模板
	temp = `
{{range .List}}
{{.name}}
{{end}}`
)

// ListEntry ...
type ListEntry struct {
	List []interface{} `json:"list"`
}

// Item ...
type Item struct {
	Name string `json:"name"`
}

// TestInterfaceTemplate ...
func TestInterfaceTemplate(t *testing.T) {

	// test
	const tttt = `{"List":[{"name":"xx"}, {"name":"ss"}]}`
	var list = ListEntry{}
	err := json.Unmarshal([]byte(tttt), &list)
	if err != nil {
		t.Fatalf("json Unmarshal:%v", err)
	}

	// template
	tt := template.New("")
	tt, err = tt.Parse(temp)
	if err != nil {
		t.Fatalf("Parse fail")
	}

	// execute
	err = tt.Execute(os.Stdout, list)
	if err != nil {
		t.Fatalf("Execute fail:%v", err)
	}
}
