package output

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/owenrumney/gtail/pkg/logger"
)

type output struct {
	templateString       string
	interestedSeverities []string
	defaultFormatFunc    func(value interface{}) error
}

type Output interface {
	Write(value interface{}) error
}

func New(templateString string, interestedSeverities []string, defaultFormatFunc func(value interface{}) error) Output {
	return &output{
		templateString:       templateString,
		interestedSeverities: interestedSeverities,
		defaultFormatFunc:    defaultFormatFunc,
	}
}

func (o *output) Write(value interface{}) error {
	if value == nil {
		logger.Warn("No value to output")
		return nil
	}

	logger.Debug("Output: %v", value)
	if o.templateString == "" {
		if o.defaultFormatFunc != nil {
			return o.defaultFormatFunc(value)
		}
		fmt.Printf("%v\n", value)
		return nil
	} else {
		switch strings.ToLower(o.templateString) {
		case "json":
			return o.writeJson(value)
		default:
			if o.templateString != "" && !strings.HasSuffix(o.templateString, "\n") {
				o.templateString = fmt.Sprintf("%s\n", o.templateString)
			}
			return o.writeTemplate(value)
		}
	}
}

func (o *output) writeTemplate(value interface{}) error {

	t, err := template.New("output").Parse(o.templateString)
	if err != nil {
		logger.Error("Error parsing template: %v", err)
		fmt.Printf("%v\n", err)
	}
	if err := t.Execute(os.Stdout, value); err != nil {
		logger.Error("Error executing template: %v", err)
		fmt.Printf("%v\n", err)
	}

	return nil
}
