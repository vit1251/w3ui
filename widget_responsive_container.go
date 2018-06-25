package w3ui

import "io"

type WidgetResponsiveContainer struct {
	rows []Widget
}

func NewWidgetResponsiveContainer() (*WidgetContainer) {
	newWidget := &WidgetResponsiveContainer{}
	return newWidget
}

func (w *WidgetResponsiveContainer) AddRow(w Widget) {
	w.rows = append(w.rows, w)
}

func (w WidgetResponsiveContainer) Execute(wr io.Writer, data interface{}) error {
	wr.Write( []byte("<div class=\"w3-container\">\n") )
	for _, row := range w.rows {
		wr.Write( []byte("    <div class=\"w3-row\">\n" ) )
		if err := row.Execute(wr, nil) ; err != nil {
			return err
		}
		wr.Write( []byte("    </div>\n") )
	}
	wr.Write( []byte("</div>\n") )
	return nil
}
