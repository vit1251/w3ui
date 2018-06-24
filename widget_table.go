package w3ui

import "log"
import "fmt"
import "html/template"
import "strings"

type WidgetTableHeader struct {
	Title string
	Name string
	Userdata interface{}
	RenderValue func ( st interface{}, wt WidgetTableHeader )( string )
	NoValue string
}

type WidgetTable struct {
	Headers []WidgetTableHeader
	Rows []interface{}
}

func NewWidgetTable() ( *WidgetTable ) {
	wt := &WidgetTable{}
	//wt.Rows = make(interface{}, 0)
	return wt
}

func (wt *WidgetTable) SetHeaders( Headers []WidgetTableHeader ) {
	wt.Headers = Headers
}

func (wt *WidgetTable) SetRows( Rows []interface{} ) {
	wt.Rows = Rows
}

func (wt *WidgetTable) Render() (template.HTML) {
	var items []string
	items = append(items, "<table class=\"w3-table\">")

	// Step 2. Render header
	items = append(items, "	<thead>")
	items = append(items, "		<tr class=\"w3-red\">")
	for _, header := range wt.Headers {
		items = append(items, fmt.Sprintf("			<th>%s</th>", header.Title))
	}
	items = append(items, "		</tr>")
	items = append(items, "	</thead>")

	// Step 3. Render content
	items = append(items, "	<tbody>")
	for _, row := range wt.Rows {

		log.Printf("row = %v", row)

		items = append(items, "	<tr>")

		for _, header := range wt.Headers {
			log.Printf("header = %v", header)
			if header.RenderValue != nil {
				content := header.RenderValue( row, header )
				items = append(items, content)
			} else {
				items = append(items, fmt.Sprintf("		<td>%s</td>", header.NoValue))
			}

		}

		items = append(items, "	</tr>")
	}
	items = append(items, "	</tbody>")

	items = append(items, "</table>")
	content := strings.Join(items, "\n")
	return template.HTML(content)
}
