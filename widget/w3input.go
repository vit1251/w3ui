package widget

import (
	. "github.com/vit1251/html_builder"
)

func W3Input(attributes Attrs, children ...Node) Node {
	return Input(attributes, children...)
}
