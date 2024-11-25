package w3ui

import (
	. "github.com/vit1251/html_builder"
)

func W3Error(reason string, err error) Node {
        if reason == "" {
            reason = "Ooops..."
        }
        var msg string = "Unknown error"
	return W3(Attrs{},
		H1(Attrs{}, Text(reason)),
		P(Attrs{}, Text(msg)),
	)
}

func W3(attributes Attrs, children ...Node) Node {
	return Html(Attrs{},
		Head(Attrs{},
			Link(Attrs{"rel": "stylesheet", "href": "https://www.w3schools.com/w3css/4/w3.css"}), // <link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css"> 
			Meta(Attrs{"name": "viewport", "content": "width=device-width, initial-scale=1.0, maximum-scale=1.0"}),// <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=1.0\">
		),
		Body(Attrs{}, children...),
	)
}
