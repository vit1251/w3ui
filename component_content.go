package w3ui

type ComponentContent struct {
	content	string
}

func NewComponentContent() (*ComponentContent) {
	cc := &ComponentContent{}
	return cc
}

func (cc *ComponentContent) Content() (string) {
	return cc.content
}

func (cc *ComponentContent) SetContent(content string) {
	cc.content = content
}
