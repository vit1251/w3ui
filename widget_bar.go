package w3ui

import "fmt"
import "strings"
import "html/template"

type WidgetBarButton struct {
	Name string
	URL string
}

type WidgetBar struct {
	Buttons []WidgetBarButton
}

func NewWidgetBar() (*WidgetBar, error) {
	return &WidgetBar{}, nil
}

func (w *WidgetBar) AddButton(name string, URL string) {
	b := WidgetBarButton{
		Name: name,
		URL: URL,
	}
	w.Buttons = append(w.Buttons, b)
}

func (w *WidgetBar) Render() template.HTML {
	var items []string

	items = append(items, "<div class=\"w3-bar w3-dark-grey\" style=\"margin-top:15px\">")
	for _, b := range w.Buttons {
		r := fmt.Sprintf("<a href=\"%s\" class=\"w3-bar-item w3-button w3-yellow\">%s</a>", b.URL, b.Name)
		items = append(items, r)
	}
	items = append(items, "</div>")
	content := strings.Join(items, "")
	return template.HTML(content)
}
