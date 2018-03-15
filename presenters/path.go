package presenters

import (
	"bytes"
	"text/template"
)

type PathPresenter struct {
    Template string
}

func (p PathPresenter) Present(data interface{}) (string, error) {
	if p.Template == "" {
		p.Template = "{{.Metric}}"
	}

	var metric bytes.Buffer
	tmpl, err := template.New("metric").Parse(p.Template)

	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&metric, data)
	if err != nil {
		return "", err
	}

	return metric.String(), nil
}
