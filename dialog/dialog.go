package dialog

import (
	"errors"
	"github.com/guark/guark/app"
	"github.com/sqweek/dialog"
)

// Dialog plugin
type Plugin struct {}

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

			if c.Has("message") == false {
				return nil, errors.New("message param required")
			}

			dialog.Message("%s", c.Get("message").(string)).Title(c.GetOr("title", c.App.Name).(string)).Info()
			return nil, nil
		},

		"error": func(c app.Context) (interface{}, error) {

			if c.Has("message") == false {
				return nil, errors.New("message param required")
			}

			dialog.Message("%s", c.Get("message").(string)).Title(c.GetOr("title", c.App.Name).(string)).Error()
			return nil, nil
		},

		"file": func(c app.Context) (path interface{}, err error) {
			return dialog.File().Title(c.GetOr("title", c.App.Name).(string)).Load()
		},

		"dir": func(c app.Context) (path interface{}, err error) {
			return dialog.Directory().Title(c.GetOr("title", c.App.Name).(string)).Browse()
		},
	}
}
