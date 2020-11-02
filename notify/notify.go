package notify

import (
	"errors"
	"github.com/guark/guark/app"
	"github.com/gen2brain/beeep"
)

// Clipboard plugin.
type Plugin struct {}

// Init the plugin.
func (p *Plugin) Init(a app.App) {

	a.Log.Debug("Init: notify plugin.")
}

// GetName returns plugin name.
func (p Plugin) GetName() string {
	return "notify"
}

// GetVersion returns plugin version.
func (p Plugin) GetVersion() string {
	return "v0.0.0"
}

// GetFuncs returns plugin functions.
func (p *Plugin) GetFuncs() map[string]app.Func {

	return map[string]app.Func{

		"send": func(c app.Context) (interface{}, error) {

			if c.Has("message") == false {
				return nil, errors.New("message param required!")
			}

			return nil, beeep.Notify(
				c.GetOr("title", c.App.Name).(string),
				c.Get("message").(string),
				c.App.Path("icon.png"),
			)
		},
	}
}
