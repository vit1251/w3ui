package w3ui

import "strings"
import "html/template"

type WidgetContainer struct {
	Widgets []Widget
}

func NewWidgetContainer() (*WidgetContainer, error) {
	return &WidgetContainer{}, nil
}

func (w *WidgetContainer) AddOne(item Widget) {
	w.Widgets = append(w.Widgets, item)
}

func (w *WidgetContainer) Render() template.HTML {

	var items []string

	items = append(items, "<div class=\"w3-container w3-padding-24\">")
	for _, row := range w.Widgets {
		items = append(items, "<div class=\"w3-row\">")
		content := row.Render()
		items = append(items, string(content))
		items = append(items, "</div>")
	}
	items = append(items, "</div>")

	body := strings.Join(items, "")
	return template.HTML(body)
}
