package layout

import (
	. "github.com/vit1251/html_builder"
)

func W3Container(attributes Attrs, children ...Node) Node {
    return Div(attributes, children...)
}
