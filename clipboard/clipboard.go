package clipboard

import (
	"errors"
	"github.com/guark/guark/app"
	"github.com/atotto/clipboard"
)

// Clipboard plugin
type Plugin struct {}

// Do nothing
func (p Plugin) Init(a app.App) {
	a.Log.Debug("Init: clipboard plugin.")
}

// Get plugin name
func (p Plugin) GetName() string {
	return "clipboard"
}

// Get plugin version
func (p Plugin) GetVersion() string {
	return "v0.0.0"
}

// Get plugin functions
func (p Plugin) GetFuncs() map[string]app.Func {

	return map[string]app.Func{

		"read": func(c app.Context) (interface{}, error) {
			return clipboard.ReadAll()
		},

		"write": func(c app.Context) (interface{}, error) {

			if c.Has("text") == false {
				return nil, errors.New("text param is required!")
			}

			return nil, clipboard.WriteAll(c.Get("text").(string))
		},
	}
}
