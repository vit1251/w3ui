package w3ui

import "fmt"
import "html/template"

type Style struct {
	URL	string
}

type JavaScript struct {
	URL	string
}

type Node struct {
	cc	*ComponentContent
}

func NewStyle(URL string) (*Style) {
	s := &Style{
		URL: URL,
	}
	return s
}

func NewJavaScript(URL string) (*JavaScript) {
	js := &JavaScript{
		URL: URL,
	}
	return js
}

type DocumentHead struct {
	Title string
	Styles []Style
	JavaScripts []JavaScript
}

type DocumentBody struct {
	Nodes	[]Node
}

type Document struct {
    Head DocumentHead
    Body DocumentBody
}

func NewDocument() (*Document, error) {
	doc := &Document{}
	return doc, nil
}

func (h *DocumentHead) AddStyle(s Style) {
	h.Styles = append(h.Styles, s)
}

func (h *DocumentHead) AddJavaScript(js JavaScript) {
	h.JavaScripts = append(h.JavaScripts, js)
}

func (b *DocumentBody) AddNode(cc *ComponentContent) {
	node := Node{
		cc: cc,
	}
	b.Nodes = append(b.Nodes, node)
}

func (doc *Document) Content() (template.HTML, error) {

	var body string

	body += "<!DOCTYPE html>\n"
	body += "<html>\n"
	body += "<head>\n"
	body += fmt.Sprintf("\t<meta charset=\"%s\">\n", "utf-8")
	body += "\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=1.0\">\n"
	body += fmt.Sprintf("\t<title>%s</title>\n", doc.Head.Title)

	// Step 1. Pack CSS styles
	for _, s := range doc.Head.Styles {
		link := fmt.Sprintf("\t<link rel=\"stylesheet\" href=\"%s\">\n", s.URL)
		body += link
	}
	// Step 2. Pack JavaScript
	for _, js := range doc.Head.JavaScripts {
		script := fmt.Sprintf("\t<script src=\"%s\"></script>\n", js.URL)
		body += script
	}

	body += "</head>\n"
	body += "<body>\n"
	// Step 3. Pack component
	for _, node := range doc.Body.Nodes {
		body += fmt.Sprintf( "%s", node.cc.Content() )
	}
	body += "</body>\n"
	body += "</html>\n"

	return template.HTML(body), nil

}
