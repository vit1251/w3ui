package widget

import "strings"
import "fmt"
import "html/template"


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

func (wf *WidgetForm) Render() ( template.HTML ) {

	var items []string

	items = append(items, fmt.Sprintf("<form class=\"w3-container\" enctype=\"application/x-www-form-urlencoded\" action=\"%s\" accept-charset=\"utf-8\" method=\"post\">", wf.Action))
	for _, input := range wf.inputs {
		//log.Printf("input = %v", input)
		content := input.Render()
		str := string(content)
		items = append(items, str)
	}
	items = append(items, "</form>")

	content := strings.Join(items, "\n")

	return template.HTML(content)
}
