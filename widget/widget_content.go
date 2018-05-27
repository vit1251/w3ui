package widget

import "html/template"

type WidgetContent struct {
	Content string
}

func NewWidgetContent() (*WidgetContent, error) {
	return &WidgetContent{}, nil
}

func (w *WidgetContent) SetContent(content string) {
	w.Content = content
}

func (w *WidgetContent) Render() template.HTML {
	return template.HTML(w.Content)
}
