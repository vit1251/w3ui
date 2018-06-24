package w3ui

import "fmt"
import "io"

type Style struct {
	URL	string
}

type JavaScript struct {
	URL	string
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

type Document struct {

	/* Docuemnt HEAD section */
	title string
	styles []Style
	javaScripts []JavaScript

	/* Document BODY section */
	widget  Widget

	/* Other options */
	writeDebugComment bool

}

// NewDocument provide new document ready for execute and write
func NewDocument() (*Document, error) {
	doc := &Document{
		writeDebugComment: false,
	}
	return doc, nil
}

// Debug enable or disable debug page output
func (doc *Document) Debug(writeDebugComment bool) {
	doc.writeDebugComment = writeDebugComment
}

func (doc *Document) Title(title string) {
	doc.title = title
}

func (doc *Document) AddStyle(s Style) {
	doc.styles = append(doc.styles, s)
}

func (doc *Document) AddJavaScript(js JavaScript) {
	doc.javaScripts = append(doc.javaScripts, js)
}

func (doc *Document) writeCharset(wr io.Writer) error {

	charset := fmt.Sprintf("    <meta charset=\"%s\">\n", "utf-8")

	wr.Write( []byte( charset ) )

	return nil

}

func (doc *Document) writeTitle(wr io.Writer) error {

	contentTitle := fmt.Sprintf("    <title>%s</title>\n", doc.title )

	wr.Write( []byte( contentTitle ) )

	return nil
}

func (doc *Document) writeBody(wr io.Writer) error {

	wr.Write( []byte( "<body>\n" ) )

	if doc.widget != nil {

		doc.widget.Execute(wr, nil)

	} else {

		wr.Write( []byte("    <h1>Ooops...</h1>\n") )
		wr.Write( []byte("    \n") )
		wr.Write( []byte("    <p>View does not contain main widget.</p>\n") )

	}

	wr.Write( []byte( "</body>\n" ) )

	return nil
}

// writeDebug write duration page render
func (doc *Document) writeDebug(wr io.Writer) error {

	if doc.writeDebugComment {

		duration := 0.0

		debugMessage := fmt.Sprintf("<!-- Page generation %f sec. -->\n", duration)
		wr.Write( []byte( debugMessage ) )

	}

	return nil
}

// writeMeta provide page meta information attributes
func (doc *Document) writeMeta(wr io.Writer) error {

	wr.Write( []byte( "    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=1.0\">\n" ) )

	return nil
}

func (doc *Document) writeHead(wr io.Writer) error {

	wr.Write( []byte( "<head>\n" ) )

	if err := doc.writeCharset( wr ) ; err != nil {
		return err
	}

	if err := doc.writeMeta( wr ) ; err != nil {
		return err
	}

	if err := doc.writeTitle( wr ) ; err != nil {
		return err
	}

	/* Write CSS styles */
	for _, s := range doc.styles {
		link := fmt.Sprintf("    <link rel=\"stylesheet\" href=\"%s\">\n", s.URL)
		wr.Write( []byte( link ) )
	}
	/* Write JavaScript section */
	for _, js := range doc.javaScripts {
		script := fmt.Sprintf("\t<script src=\"%s\"></script>\n", js.URL)
		wr.Write( []byte( script ) )
	}

	wr.Write( []byte( "</head>\n" ) )

	return nil
}

func (doc *Document) Execute(wr io.Writer, data interface{}) error {

	wr.Write( []byte( "<!DOCTYPE html>\n" ) )
	wr.Write( []byte( "<html>\n" ) )

	/* Write HEAD content section */
	if err := doc.writeHead( wr ) ; err != nil {
		return err
	}

	/* Write BODY content section */
	if err := doc.writeBody( wr ) ; err != nil {
		return err
	}

	wr.Write( []byte( "</html>\n" ) )

	/* Write debug metrics */
	if err := doc.writeDebug( wr ) ; err != nil {
		return err
	}

	return nil

}
