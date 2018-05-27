package widget

import "html/template"

type Widget interface {
	Render() template.HTML
}
