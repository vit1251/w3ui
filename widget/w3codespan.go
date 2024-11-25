package widget

import (
	. "github.com/vit1251/html_builder"
)

func W3Cardspan(attributes Attrs, code string) Node {
	return Div(attributes, Text(code))
}
