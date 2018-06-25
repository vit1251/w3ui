package w3ui

import "fmt"
import "io"

type WidgetHeader struct {

	label string
	level uint8

}

func NewWidgetHeader() (*WidgetHeader, error) {
	w := &WidgetHeader{
		level: 1,
	}
	return w, nil
}

func (w *WidgetHeader) SetLevel(level uint8) {
	w.level = level
}

func (w *WidgetHeader) SetLabel(label string) {
	w.label = label
}

func (w *WidgetHeader) Execute(wr io.Writer, data interface{}) error {

	content := fmt.Sprintf("<h%d>%s</h%d>", w.level, w.label, w.level)
	wr.Write( []byte(content) )

	return nil

}
