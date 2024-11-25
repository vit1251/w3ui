package widget

import (
	. "github.com/vit1251/html_builder"
	"testing"
)

func TestUpdateOnEmptyAttribute(t *testing.T) {

    /* Step 1. Checks */
    attrs := Attrs{}
    attrs = update(attrs, "class", func(value *string) *string {
            if value != nil {
                t.Errorf("No empty attribute.")
            }
            var result string = "w3-box"
            return &result
    })

    /* Step 2. Prepare actual and expected */
    actual := attrs["class"]
    expected := "w3-box"

    /* Step 3. Compare */
    if expected != actual {
	t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
    }

}

func TestUpdateOnZeroValueAttribute(t *testing.T) {

    /* Step 1. Checks */
    attrs := Attrs{"class": ""}
    attrs = update(attrs, "class", func(value *string) *string {
            if *value != "" {
                t.Errorf("No empty attribute.")
            }
            var result string = "w3-box"
            return &result
    })

    /* Step 2. Prepare actual and expected */
    actual := attrs["class"]
    expected := "w3-box"

    /* Step 3. Compare */
    if expected != actual {
	t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
    }

}

func TestUpdateOnSomeValueAttribute(t *testing.T) {

    /* Step 1. Checks */
    attrs := Attrs{"class": "w3-panel"}
    attrs = update(attrs, "class", func(value *string) *string {
            if *value != "w3-panel" {
                t.Errorf("No empty attribute.")
            }
            var result string = "w3-box"
            return &result
    })

    /* Step 2. Prepare actual and expected */
    actual := attrs["class"]
    expected := "w3-box"

    /* Step 3. Compare */
    if expected != actual {
	t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
    }

}

func TestExtend(t *testing.T) {

    /* Step 1. Checks */
    attrs := Attrs{"class": "w3-panel"}
    attrs = extend(attrs, "class", "w3-yellow")

    /* Step 2. Prepare actual and expected */
    actual := attrs["class"]
    expected := "w3-panel w3-yellow"

    /* Step 3. Compare */
    if expected != actual {
        t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
    }

}
