package templater

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/Masterminds/sprig"
)

//go:embed templates/command.template
var commandTemplate string

//go:embed templates/command_file.template
var commandFileTemplate string

//go:embed templates/header.template
var headerTemplate string

//go:embed templates/main.template
var mainTemplate string

//go:embed templates/root.template
var rootTemplate string

type additionalTemplate struct {
	templateName    string
	templateContent string
}

func GenerateContent(data RootData) ([]byte, error) {
	return GenerateRoot(data)
}

func GenerateMain(mainData MainData) ([]byte, error) {
	return generateContent("MainTemplate", mainTemplate, mainData, additionalTemplate{"HeaderTemplate", headerTemplate})
}

func GenerateRoot(rootData RootData) ([]byte, error) {
	return generateContent("RootTemplate", rootTemplate, rootData,
		additionalTemplate{"HeaderTemplate", headerTemplate},
		additionalTemplate{"CommandTemplate", commandTemplate})
}

func GenerateCommand(commandData Command) ([]byte, error) {
	return generateContent("CommandFileTemplate", commandFileTemplate, commandData,
		additionalTemplate{"HeaderTemplate", headerTemplate},
		additionalTemplate{"CommandTemplate", commandTemplate})
}

func generateContent(templateName, templateText string, data interface{}, optTemplates ...additionalTemplate) ([]byte, error) {
	tp := template.New(templateName)
	tp.Funcs(sprig.FuncMap()).Funcs(newIncludeFuncMap(tp))
	_, err := tp.Parse(templateText)
	if err != nil {
		return nil, err
	}
	for _, optTemplate := range optTemplates {
		_, err := tp.New(optTemplate.templateName).Parse(optTemplate.templateContent)
		if err != nil {
			return nil, err
		}
	}

	buf := &bytes.Buffer{}
	err = tp.Execute(buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func newIncludeFuncMap(tp *template.Template) template.FuncMap {
	return template.FuncMap{
		"include": newIncludeFn(tp),
	}
}

func newIncludeFn(tp *template.Template) func(name string, data interface{}) (string, error) {
	return func(name string, data interface{}) (string, error) {
		buf := bytes.NewBuffer(nil)
		if err := tp.ExecuteTemplate(buf, name, data); err != nil {
			return "", err
		}
		return buf.String(), nil
	}
}
