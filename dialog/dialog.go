package dialog

import (
	"errors"
	"github.com/guark/guark/app"
	"github.com/gen2brain/dlgs"
)

// Dialog plugin
// TODO: add all dlgs methods
type Plugin struct {

}

// Do nothing
func (p Plugin) Init(a app.App) {
	a.Log.Debug("Init: dialog plugin.")
}

// Get plugin name
func (p Plugin) GetName() string {
	return "dialog"
}

// Get plugin version
func (p Plugin) GetVersion() string {
	return "v0.0.0"
}

// Get plugin functions
func (p Plugin) GetFuncs() map[string]app.Func {

	return map[string]app.Func{

		"info": func(c app.Context) (interface{}, error) {

			if c.Params.Has("text") == false {
				return nil, errors.New("text param required")
			}


			return dlgs.Info(c.Params.GetOr("title", c.App.Name).(string), c.Params.Get("text").(string))
		},

		"warning": func(c app.Context) (interface{}, error) {

			if c.Params.Has("text") == false {
				return nil, errors.New("text param required")
			}


			return dlgs.Warning(c.Params.GetOr("title", c.App.Name).(string), c.Params.Get("text").(string))
		},

		"file": func(c app.Context) (path interface{}, err error) {

			path, _, err = dlgs.File(c.Params.GetOr("title", c.App.Name).(string), "", false)
			return
		},
	}
}
