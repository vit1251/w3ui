package w3ui

import "fmt"
import "io"

type WidgetFormSubmit struct {
	Value string
}

func NewWidgetFormSubmit() (*WidgetFormSubmit, error) {
	w := &WidgetFormSubmit{}
	return w, nil
}

func (w *WidgetFormSubmit) SetValue(value string) {
	w.Value = value
}

func (widgetFormSubmit *WidgetFormSubmit) Execute(wr io.Writer, data interface{}) error {

	wr.Write( []byte( "<p>" ) )

	content := fmt.Sprintf("<input type=\"submit\" value=\"%s\" class=\"w3-button w3-padding w3-orange\" />",
			widgetFormSubmit.Value,
		)
	wr.Write( []byte( content ) )

	wr.Write( []byte( "</p>" ) )

	return nil
}
