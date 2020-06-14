package notify

import (
	"errors"
	"github.com/guark/guark/app"
	"github.com/0xAX/notificator"
)

// Clipboard plugin
type Plugin struct {
	notify *notificator.Notificator
}

// Init the plugin
func (p *Plugin) Init(a app.App) {

	a.Log.Debug("Init: notify plugin.")

	p.notify = notificator.New(notificator.Options{
		DefaultIcon: a.Path("icons/icon.png"),
		AppName:     a.Name,
	})
}

// Get plugin name
func (p Plugin) GetName() string {
	return "notify"
}

// Get plugin version
func (p Plugin) GetVersion() string {
	return "v0.0.0"
}

// Get plugin functions
func (p *Plugin) GetFuncs() map[string]app.Func {

	return map[string]app.Func{

		"send": func(c app.Context) (interface{}, error) {

			if c.Params.Has("text") == false {
				return nil, errors.New("text param required!")
			}

			return nil, p.notify.Push(
				c.Params.GetOr("title", c.App.Name).(string),
				c.Params.Get("text").(string),
				c.App.Path("icons/icon.png"),
				notificator.UR_CRITICAL,
			)
		},
	}
}
