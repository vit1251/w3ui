package w3ui

import "fmt"
import "io"

type WidgetFormSelect struct {
	Name string
	Options map[string]string
}

func NewWidgetFormSelect() (*WidgetFormSelect, error) {

	ws := &WidgetFormSelect{}
	ws.Options = make(map[string]string, 0)

	return ws, nil
}

func (ws *WidgetFormSelect) SetName(name string) {
	ws.Name = name
}

func (ws *WidgetFormSelect) AddOption(name string, value string) {
	ws.Options[value] = name
}

//
func (widgetFormSelect *WidgetFormSelect) writeLabel(wr io.Writer) error {

//	content := fmt.Sprintf("<label for=\"\"></label>")
	wr.Write( []byte("<label for=\"\"></label>") )

	return nil
}

func (widgetFormSelect *WidgetFormSelect) writeStartSelect(wr io.Writer) error {

	content := fmt.Sprintf("<select name=\"%s\">", widgetFormSelect.Name)
	wr.Write( []byte(content) )

	return nil
}

func (widgetFormSelect *WidgetFormSelect) writeStopSelect(wr io.Writer) error {

	content := "</select>"
	wr.Write( []byte(content) )

	return nil
}


func (widgetFormSelect *WidgetFormSelect) Execute(wr io.Writer, data interface{}) error {

	/* Write label */
	if err := widgetFormSelect.writeLabel(wr) ; err != nil {
		return err
	}

	/* Write start select */
	if err := widgetFormSelect.writeStartSelect(wr) ; err != nil {
		return err
	}

	/* Write option */
	for name, value := range widgetFormSelect.Options {
		content := fmt.Sprintf("<option value=\"%s\">%s</option>", value, name)
		wr.Write( []byte(content) )
	}

	/* Write stop select */
	if err := widgetFormSelect.writeStopSelect(wr) ; err != nil {
		return err
	}

	return nil
}
