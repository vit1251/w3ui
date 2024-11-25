package widget

import (
	. "github.com/vit1251/html_builder"
)

func W3Card(attributes Attrs, children ...Node) Node {
	return Div(attributes, children...)
}
