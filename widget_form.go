package w3ui

import "io"
import "fmt"

type WidgetForm struct {

	Action string
	Method string

	inputs []Widget
}

func NewWidgetForm() ( *WidgetForm ) {
	wf := &WidgetForm{
		Method: "post",
	}
	return wf
}

func (wf *WidgetForm) AppendInput( input Widget ) {
	wf.inputs = append(wf.inputs, input)
}

func (wf *WidgetForm) Validate() {
}

func (wf *WidgetForm) Execute(wr io.Writer, data interface{}) error {

	startForm := fmt.Sprintf("<form class=\"w3-container\" enctype=\"application/x-www-form-urlencoded\" action=\"%s\" accept-charset=\"utf-8\" method=\"post\">", wf.Action)
	wr.Write( []byte( startForm ) )
	for _, input := range wf.inputs {
		if err := input.Execute(wr, nil) ; err != nil {
			return err
		}
	}
	stopForm := "</form>"
	wr.Write( []byte( stopForm ) )

	return nil
}
