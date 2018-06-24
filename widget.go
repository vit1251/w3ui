package w3ui

import "io"

type Widget interface {
	Execute(wr io.Writer, data interface{}) error
}
