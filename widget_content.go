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
	if err := w.writer.Flush(); err != nil {
		return err
	}
	reader := NewReader(w.buffer)
	io.Copy(wr, reader)
	return nil
}
