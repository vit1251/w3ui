package widget

import (
	. "github.com/vit1251/html_builder"
)

func W3Tag(attributes Attrs, children ...Node) Node {
	return Div(attributes, children...)
}
