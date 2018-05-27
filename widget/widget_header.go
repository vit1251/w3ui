package widget

import "fmt"
import "html/template"

type WidgetHeader struct {

	Label string
	Level uint8

}

func NewWidgetHeader() (*WidgetHeader, error) {
	w := &WidgetHeader{
		Level: 1,
	}
	return w, nil
}

func (w *WidgetHeader) SetLevel(level uint8) {
	w.Level = level
}

func (w *WidgetHeader) SetLabel(label string) {
	w.Label = label
}

func (w *WidgetHeader) Render() template.HTML {

	var content string

	content += fmt.Sprintf("<h%d>%s</h%d>", w.Level, w.Label, w.Level)

	return template.HTML(content)
}
