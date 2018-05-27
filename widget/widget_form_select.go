package widget

//          <label for="issue_status_id">Status<span class="required"> *</span></label>
//          <select class="w3-select w3-border" name="issue.status_id" id="issue_status_id">
//            <option selected="selected" value="1">New</option>
//          </select>


import "fmt"
import "html/template"

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

func (ws *WidgetFormSelect) Render() template.HTML {

	var content string

	content += fmt.Sprintf("<label for></label>")
	content += fmt.Sprintf("<select name=\"%s\">", ws.Name)
	for name, value := range ws.Options {
		content += fmt.Sprintf("<option value=\"%s\">%s</option>", value, name)
	}
	content += "</select>"

	return template.HTML(content)
}
