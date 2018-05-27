package widget

import "fmt"
import "html/template"

type WidgetFormSubmit struct {
	Value string
}

func NewWidgetFormSubmit() (*WidgetFormSubmit, error) {
	w := &WidgetFormSubmit{}
	return w, nil
}

func (w *WidgetFormSubmit) SetValue(value string) {
	w.Value = value
}

func (i *WidgetFormSubmit) Render() template.HTML {

	var content string

	content += "<p>"
	content += fmt.Sprintf("<input type=\"submit\" value=\"%s\" class=\"w3-button w3-padding w3-orange\" />", i.Value)
	content += "</p>"

	return template.HTML(content)
}
