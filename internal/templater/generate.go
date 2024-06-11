package templater

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/Masterminds/sprig"
)

//go:embed templates/exit.template
var exitTemplate string

//go:embed templates/flags/flag.template
var flagTemplate string

//go:embed templates/flags/bool.template
var boolFlagTemplate string

//go:embed templates/flags/int64.template
var int64FlagTemplate string

//go:embed templates/flags/int64slice.template
var int64sliceFlagTemplate string

//go:embed templates/flags/string.template
var stringFlagTemplate string

//go:embed templates/command.template
var commandTemplate string

//go:embed templates/new_app.template
var newAppTemplate string

//go:embed templates/root.template
var rootTemplate string

type additionalTemplate struct {
	templateName    string
	templateContent string
}

// GenerateContext executes templates over data to generate cli code
func GenerateContent(data Root) ([]byte, error) {
	return generateContent("RootTemplate", rootTemplate, data,
		additionalTemplate{"NewAppTemplate", newAppTemplate},
		additionalTemplate{"CommandTemplate", commandTemplate},
		additionalTemplate{"FlagTemplate", flagTemplate},
		additionalTemplate{"StringFlagTemplate", stringFlagTemplate},
		additionalTemplate{"BoolFlagTemplate", boolFlagTemplate},
		additionalTemplate{"Int64FlagTemplate", int64FlagTemplate},
		additionalTemplate{"Int64SliceFlagTemplate", int64sliceFlagTemplate},
		additionalTemplate{"ExitTemplate", exitTemplate},
	)
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
