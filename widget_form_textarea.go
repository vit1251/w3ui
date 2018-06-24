package w3ui

import "html/template"
import "fmt"

type WidgetFormTextarea struct {
	Name        string
	Value       string
	Placeholder string
	Class       string
	Label       string
	Id          string
	Required    bool
	Help        string

	validate func(value string) bool
}

func NewWidgetFormTextarea() (*WidgetFormTextarea, error) {
	w := &WidgetFormTextarea{}
	return w, nil
}

func (w *WidgetFormTextarea) SetName(name string) {
	w.Name = name
}

func (w *WidgetFormTextarea) SetId(id string) {
	w.Id = id
}

func (w *WidgetFormTextarea) SetLabel(label string) {
	w.Label = label
}

func (w *WidgetFormTextarea) SetRequired(required bool) {
	w.Required = required
}

func (w *WidgetFormTextarea) SetPlaceholder(placeholder string) {
	w.Placeholder = placeholder
}

func (w *WidgetFormTextarea) SetValue(value string) {
	w.Value = value
}

func (i *WidgetFormTextarea) Render() template.HTML {

	var content string

	content += "<p>"

	content += fmt.Sprintf("<label for=\"%s\">%s</label>", i.Id, i.Label)
	content += fmt.Sprintf("<textarea class=\"w3-input w3-border\" rows=\"8\" name=\"%s\" id=\"%s\">%s</textarea>", i.Name, i.Id, i.Value)

	if i.Help != "" {
		content += fmt.Sprintf("<em class=\"info\">%s</em>", i.Help)
	}
	content += "</p>"

	return template.HTML(content)

}
