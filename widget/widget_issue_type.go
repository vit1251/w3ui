package widget

import "html/template"

type WidgetIssueType struct {
	Type string
}

func NewWidgetIssueType(t string) (*WidgetIssueType, error) {
	return &WidgetIssueType{Type: t}, nil
}

func (w *WidgetIssueType) Render() template.HTML {
	return template.HTML(w.Type)
}
