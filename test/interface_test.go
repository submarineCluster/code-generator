package test

import (
	"encoding/json"
	"os"
	"testing"
	"text/template"
)

const (
	temp = `
{{range .List}}
{{.name}}
{{end}}`
)

type ListEntry struct {
	List []interface{} `json:"list"`
}

type Item struct {
	Name string `json:"name"`
}

func TestInterfaceTemplate(t *testing.T) {

	const tttt = `{"List":[{"name":"xx"}, {"name":"ss"}]}`
	var list = ListEntry{}
	err := json.Unmarshal([]byte(tttt), &list)
	if err != nil {
		t.Fatalf("json Unmarshal:%v", err)
	}

	tt := template.New("")
	tt, err = tt.Parse(temp)
	if err != nil {
		t.Fatalf("Parse fail")
	}

	err = tt.Execute(os.Stdout, list)
	if err != nil {
		t.Fatalf("Execute fail:%v", err)
	}
}
