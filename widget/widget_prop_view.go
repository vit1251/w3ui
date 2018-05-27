package widget

import "fmt"
import "strings"
import "html/template"

type WidgetPropItem struct {
	Name string
	Widget Widget
}

type WidgetPropView struct {
	Props []WidgetPropItem
}

func NewWidgetPropView() (*WidgetPropView, error) {
	return &WidgetPropView{}, nil
}

func (w *WidgetPropView) AddProp(name string, value Widget) {
	p := WidgetPropItem{
		Name: name,
		Widget: value,
	}
	w.Props = append(w.Props, p)
}

func (w *WidgetPropView) Render() template.HTML {

	var items []string

	for _, p := range w.Props {
		items = append(items, fmt.Sprintf("<span class=\"name\">%s</span>", p.Name))
		v := p.Widget.Render()
		items = append(items, fmt.Sprintf("<span class=\"value\">%s</span>", v))
	}

	content := strings.Join(items, "")
	return template.HTML(content)
}
