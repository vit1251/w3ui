package widget

import (
	. "github.com/vit1251/html_builder"
)

func W3Badge(attributes Attrs, children ...Node) Node {
	attributes = extend(attributes, "class", "w3-badge")
	return Span(attributes, children...)
}
