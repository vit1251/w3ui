package widget

import "strings"
import "html/template"
import "fmt"

type WidgetIssueStatus struct {
	Status string
	Active bool
}

func NewWidgetIssueStatus(s string) (*WidgetIssueStatus, error) {
	return &WidgetIssueStatus{Status: s}, nil
}

func (w *WidgetIssueStatus) SetActive(a bool) {
	w.Active = a
}

func detectClass(status string) []string {

	var result []string

	switch status {
		case "Open":
			result = []string{"w3-tag", "w3-issue-status", "w3-white", "w3-text-blue", "w3-border-blue"}

		case "In Progress":
			result = []string{"w3-tag", "w3-issue-status", "w3-white", "w3-text-yellow", "w3-border-yellow"}

		case "Complete":
			result = []string{"w3-tag", "w3-issue-status", "w3-white", "w3-text-green", "w3-border-green"}

		default:
			result = []string{"w3-tag", "w3-issue-status", "w3-white", "w3-text-gray", "w3-border-gray"}
	}

	return result
}

func (w *WidgetIssueStatus) Render() template.HTML {

	var result string

	if w.Active {
		result += "<div class=\"w3-tag w3-round w3-white w3-border w3-border-green\">"
		result += fmt.Sprintf("<b>%s</b>", w.Status)
		result += "</div>"
	} else {
		var class []string = detectClass(w.Status)
		classes := strings.Join(class, " ")
		result += fmt.Sprintf("<div class=\"%s\">", classes)
		result += fmt.Sprintf("<b>%s</b>", w.Status)
		result += "</div>"
	}

	return template.HTML(result)
}
