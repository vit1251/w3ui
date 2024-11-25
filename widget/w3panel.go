package widget

import (
	. "github.com/vit1251/html_builder"
)

func W3Panel(attributes Attrs, children ...Node) Node {
	attributes = extend(attributes, "class", "w3-panel")
	return Div(attributes, children...)
}
