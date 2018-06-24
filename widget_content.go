package w3ui

import "io"
import "bytes"
import "bufio"

type WidgetContent struct {
	buffer  bytes.Buffer
	writer  bufio.Writer
}

func NewWidgetContent() (*WidgetContent) {
	widgetContent := &WidgetContent{
	}
	return widgetContent
}

func (w *WidgetContent) Write(content []byte) {
	w.writer.Write(content)
}

func (w *WidgetContent) Execute(wr io.Writer, data interface{}) error {

	/* Complete buffer write */
	if err := w.writer.Flush(); err != nil {
		return err
	}

	/* Create reader and copy content */
	b := w.buffer.Bytes()
	reader := bytes.NewReader( b )
	if _, err := io.Copy(wr, reader) ; err != nil {
		return err
	}

	return nil
}
