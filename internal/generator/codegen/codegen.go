package codegen

import (
	"bytes"
	"go/format"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/fatih/structtag"
	"github.com/spf13/afero"
	"golang.org/x/tools/imports"
)

type RunMethodParam struct {
	Name         string
	ValueField   string
	UpdateMethod string
}

type RunMethod struct {
	ReceiverName string
	Method       string
	URLPath      string
	Params       []RunMethodParam
}

// StructField is one field in a StructTmplHelper
type StructField struct {
	Name string
	Type string
	Tags *structtag.Tags
}

type StructTmplHelper struct {
	Name   string
	Fields []StructField
}

type CmdStructAndMethod struct {
	CmdStruct StructTmplHelper
	RunMethod RunMethod
}

type SvcTmpl struct {
	SvcStruct           StructTmplHelper
	CmdStructAndMethods []CmdStructAndMethod
}

type FileTmpl struct {
	CmdHelps       map[string]map[string]string
	FlagHelps      map[string]map[string]map[string]string
	PrimaryStructs []StructTmplHelper
	SvcTmpls       []SvcTmpl
}

func WriteFiles(files map[string]FileTmpl, outputPath string, fs afero.Fs) error {
	for name, fileTmpl := range files {
		err := writeGoFile(name, fileTmpl, outputPath, fs)
		if err != nil {
			return err
		}
	}
	return nil
}

//writeGoFile executes the named template and does the equivalent of `go fmt` and `goimports` on the output
func writeGoFile(filename string, fileTmpl FileTmpl, path string, fs afero.Fs) error {
	for _, svcTmpl := range fileTmpl.SvcTmpls {
		tmplSorting(&svcTmpl)
	}

	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "main", fileTmpl)
	if err != nil {
		return err
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	out, err = imports.Process("", out, nil)
	if err != nil {
		return err
	}
	fl := filepath.Join(path, filename)
	return afero.WriteFile(fs, fl, out, 0644)
}

func sortCmdStructFields(fields []StructField) {
	if len(fields) == 0 {
		return
	}
	newFields := make([]StructField, 0, len(fields))
	holdouts := make([]StructField, 0, len(fields))
	for _, field := range fields {
		if field.Name == "" {
			holdouts = append(holdouts, field)
			continue
		}
		newFields = append(newFields, field)
	}
	sort.Slice(newFields, func(i, j int) bool {
		return newFields[i].Name < newFields[j].Name
	})
	sort.Slice(holdouts, func(i, j int) bool {
		return holdouts[i].Type < holdouts[j].Type
	})
	newFields = append(newFields, holdouts...)
	copy(fields, newFields)
}

func tmplSorting(svcTmpl *SvcTmpl) {
	sort.Slice(svcTmpl.SvcStruct.Fields, func(i, j int) bool {
		return svcTmpl.SvcStruct.Fields[i].Name < svcTmpl.SvcStruct.Fields[j].Name
	})
	for _, csm := range svcTmpl.CmdStructAndMethods {
		sortCmdStructFields(csm.CmdStruct.Fields)

		sort.Slice(csm.RunMethod.Params, func(i, j int) bool {
			return csm.RunMethod.Params[i].Name < csm.RunMethod.Params[j].Name
		})
	}
	sort.Slice(svcTmpl.CmdStructAndMethods, func(i, j int) bool {
		return svcTmpl.CmdStructAndMethods[i].CmdStruct.Name < svcTmpl.CmdStructAndMethods[j].CmdStruct.Name
	})
}

var tmpl = template.Must(template.New("").Parse(tmplt))

// language=GoTemplate
const tmplt = `
{{define "main"}}
	// Code generated by octo-cli/generator; DO NOT EDIT.
	
	package generated
	
	import "github.com/octo-cli/octo-cli/internal"
	
	{{range .PrimaryStructs}}
	{{template "structtype" .}}
	{{end}}
	{{range .SvcTmpls}}{{template "svctmpl" .}}{{end}}
	
	{{if .CmdHelps}}
	var CmdHelps = map[string]map[string]string{
	{{range $topCmd, $topCmdVals := .CmdHelps}}"{{$topCmd}}": {
	{{range $cmd, $help := $topCmdVals}}"{{$cmd}}": {{printf "%q" $help}},
	{{end}} 
	},
	{{end}} 
	}
	{{end}}
	
	{{if .FlagHelps}}
	var FlagHelps = map[string]map[string]map[string]string{
	{{range $topCmd, $topCmdVals := .FlagHelps}}"{{$topCmd}}": {
	{{range $cmd, $flagHelps := $topCmdVals}}"{{$cmd}}": {
	{{range $flag, $help := $flagHelps}}"{{$flag}}": {{printf "%q" $help}},
	{{end}}
	}, 
	{{end}}
	}, 
	{{end}}
	} 
	{{end}}

{{end}} 

{{define "svctmpl"}}{{if .SvcStruct}}{{template "structtype" .SvcStruct}}{{end}}
	{{range .CmdStructAndMethods}}
	{{if .CmdStruct}}{{template "structtype" .CmdStruct}}{{end}}
	{{if .RunMethod}}{{template "runmethod" .RunMethod}}{{end}}
	{{end}}
{{end}}

{{define "structtype"}}type {{.Name}} struct { {{range .Fields}}
	{{.Name}} {{.Type}} {{if .Tags}}{{printf "%#q" .Tags}} {{end}}{{end}}
}{{end}}

{{define "runmethod"}}
func (c *{{.ReceiverName}}) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("{{.URLPath}}"){{range .Params}}{{template "runmethodparam" .}}{{end}}
	return c.DoRequest("{{.Method}}")
}{{end}}

{{define "runmethodparam"}}
	c.{{.UpdateMethod}}("{{.Name}}", c.{{.ValueField}}){{end}}
`
