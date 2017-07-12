package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Gallery struct {
	item.Item

	Image []string `json:"image"`
}

// MarshalEditor writes a buffer of html to edit a Gallery within the CMS
// and implements editor.Editable
func (g *Gallery) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(g,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Gallery field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.FileRepeater("Image", g, map[string]string{
				"label":       "Image",
				"placeholder": "Upload the Image here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Gallery editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Gallery"] = func() interface{} { return new(Gallery) }
}
