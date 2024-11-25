package widget

import (
	. "github.com/vit1251/html_builder"
)

func W3Button(attributes Attrs, children ...Node) Node {
	attributes = extend(attributes, "class", "w3-button")
	return Div(attributes, children...)
}
