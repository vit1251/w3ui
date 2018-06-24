package w3ui

import "io"

type WidgetContainer struct {
	widget Widget
	data interface{}
}

func NewWidgetContainer(widget Widget, data interface{}) (*WidgetContainer) {
	newWidget := &WidgetContainer{
		widget: widget,
		data: data,
	}
	return newWidget
}

func (w WidgetContainer) Execute(wr io.Writer, data interface{}) error {
	wr.Write( []byte("<div class=\"w3-container\">") )
	if err := w.widget.Execute(wr, w.data) ; err != nil {
		return err
	}
	wr.Write( []byte("</div>") )
	return nil
}
