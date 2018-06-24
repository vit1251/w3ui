package w3ui

import "fmt"
import "html/template"

type WidgetFormInput struct {
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

func NewWidgetFormInput() (*WidgetFormInput, error) {
	w := &WidgetFormInput{}
	return w, nil
}

func (w *WidgetFormInput) SetName(name string) {
	w.Name = name
}

func (w *WidgetFormInput) SetHelp(help string) {
	w.Help = help
}

func (w *WidgetFormInput) SetId(id string) {
	w.Id = id
}

func (w *WidgetFormInput) SetLabel(label string) {
	w.Label = label
}

func (w *WidgetFormInput) SetRequired(required bool) {
	w.Required = required
}

func (w *WidgetFormInput) SetPlaceholder(placeholder string) {
	w.Placeholder = placeholder
}

func (w *WidgetFormInput) SetValue(value string) {
	w.Value = value
}

func (i *WidgetFormInput) Render() template.HTML {

	var content string

	content += "<p>"
	var Label string = i.Label
	if i.Required {
		Label += "<span class=\"required\"> *</span>"
	}
	content += fmt.Sprintf("<label for=\"%s\">%s</label>", i.Id, Label)
	content += fmt.Sprintf("<input class=\"w3-input w3-border\" type=\"text\" value=\"%s\" name=\"%s\" id=\"%s\" placeholder=\"%s\" />", i.Value, i.Name, i.Id, i.Placeholder)
	if i.Help != "" {
		content += fmt.Sprintf("<em class=\"info\">%s</em>", i.Help)
	}
	content += "</p>"

	return template.HTML(content)
}
