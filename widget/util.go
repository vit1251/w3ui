package widget

import (
	. "github.com/vit1251/html_builder"
	"golang.org/x/exp/slices"
	"strings"
)

func set(attributes Attrs, name string, value *string) Attrs {
    if value == nil {
	delete(attributes, name)
    } else {
	attributes[name] = *value
    }
    return attributes
}

func update(attributes Attrs, name string, callback func (value *string) *string) Attrs {
        var value *string
	if prop, ok := attributes[name]; ok {
		value = callback(&prop)
                return set(attributes, name, value)
	} else {
		value = callback(nil)
		return set(attributes, name, value)
	}
}

func extend(attributes Attrs, name string, value string) Attrs {
	return update(attributes, name, func(prop *string) *string {
		if (prop == nil) {
			return &value
		}
		var props []string = strings.Split(*prop, " ")
		if !slices.Contains(props, value) {
			props = append(props, value)
		}
		var result string = strings.Join(props, " ")
		return &result
	})
}
