package widget

import "fmt"
import "strings"
import "html/template"

type WidgetPaginator struct {
	Paginator *Paginator
}

func NewWidgetPaginator() (*WidgetPaginator, error) {
	return &WidgetPaginator{}, nil
}

func (w *WidgetPaginator) SetPaginator(p *Paginator) {
	w.Paginator = p
}

func (w *WidgetPaginator) Render() ( template.HTML ) {

	var items []string

	items = append(items, "<div class=\"w3-container w3-padding-24\">")
	items = append(items, "	<div class=\"w3-bar\">")

	pages := w.Paginator.Pages()
	for _, page := range pages {
		var pageURL string
		pageURL = page.URL
		if pageURL == "" {
			pageURL = fmt.Sprintf("?page=%d", page.Page)
		}
		var classNames []string
		classNames = append(classNames, "w3-bar-item")
		classNames = append(classNames, "w3-button")
		if page.Active {
			classNames = append(classNames, "w3-red")
		} else {
			classNames = append(classNames, "w3-black")
		}
		class := strings.Join(classNames, " ")
		items = append(items, fmt.Sprintf("<a href=\"%s\" class=\"%s\">%d</a>", pageURL, class, page.Page))
	}
	items = append(items, "	</div>")
	items = append(items, "</div>")

	content := strings.Join(items, "\n")

	return template.HTML(content)
}
