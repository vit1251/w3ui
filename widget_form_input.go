package w3ui

import "fmt"
import "io"

type WidgetFormInput struct {
	Name        string
	Value       string
	Placeholder string
	Class       string
	Label       string
	Id          string
	required    bool
	Help        string

	validate func(value string) bool
}

func NewWidgetFormInput() (*WidgetFormInput, error) {
	w := &WidgetFormInput{}
	return w, nil
}

func (w *WidgetFormInput) SetName(name string) {
	w.Name = name
}

func (w *WidgetFormInput) SetHelp(help string) {
	w.Help = help
}

func (w *WidgetFormInput) SetId(id string) {
	w.Id = id
}

func (w *WidgetFormInput) SetLabel(label string) {
	w.Label = label
}

func (w *WidgetFormInput) SetRequired(required bool) {
	w.required = required
}

func (w *WidgetFormInput) SetPlaceholder(placeholder string) {
	w.Placeholder = placeholder
}

func (w *WidgetFormInput) SetValue(value string) {
	w.Value = value
}

func (widgetFormInput *WidgetFormInput) writeLabel(wr io.Writer) error {

//	if widgetFormInput.required {
//		Label += "<span class=\"required\"> *</span>"
//	}
//	content = fmt.Sprintf("<label for=\"%s\">%s</label>", i.Id, Label)

	return nil

}

// 
func (widgetFormInput *WidgetFormInput) writeHelp(wr io.Writer) error {

	if widgetFormInput.Help != "" {
		content := fmt.Sprintf("<em class=\"info\">%s</em>", widgetFormInput.Help)
		wr.Write( []byte( content ) )
	}

	return nil
}

func (widgetFormInput *WidgetFormInput) Execute(wr io.Writer, data interface{}) error {

	wr.Write( []byte( "<p>" ) )

	/* Write message */
	if err := widgetFormInput.writeLabel(wr) ; err != nil {
		return err
	}

	/* Write input */
	content := fmt.Sprintf("<input class=\"w3-input w3-border\" type=\"text\" value=\"%s\" name=\"%s\" id=\"%s\" placeholder=\"%s\" />",
			widgetFormInput.Value,
			widgetFormInput.Name,
			widgetFormInput.Id,
			widgetFormInput.Placeholder
		)
	wr.Write( []byte( content ) )

	/* Write help */
	if err := widgetFormInput.writeHelp(wr) ; err != nil {
		return err
	}

	wr.Write( []byte( "</p>" ) )

	return nil
}
