package dialog

import (
	"strings"

	"github.com/guark/guark/app"
	"github.com/pkg/errors"
	"github.com/sqweek/dialog"
)

const (
	Title   = "title"
	Message = "message"

	SaveOp    = "save"
	LoadOp    = "load"
	BrowseOp  = "browse"
	Operation = "operation"

	FileTypeFilterDesc = "file_type_filter_desc"
	FileTypeFilter     = "file_type_filter"

	StartDir = "start_dir"
)

var ErrInvalidOperation = errors.New("invalid operation")

// Dialog plugin
type Plugin struct{}

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

			if c.Has(Message) == false {
				return nil, errors.New("message param required")
			}

			dialog.Message("%s", c.Get(Message).(string)).Title(c.GetOr(Title, c.App.Name).(string)).Info()
			return nil, nil
		},

		"error": func(c app.Context) (interface{}, error) {

			if c.Has(Message) == false {
				return nil, errors.New("message param required")
			}

			dialog.Message("%s", c.Get(Message).(string)).Title(c.GetOr(Title, c.App.Name).(string)).Error()
			return nil, nil
		},

		"file": func(c app.Context) (path interface{}, err error) {
			fd := dialog.File().Title(c.GetOr(Title, c.App.Name).(string))
			if ftf, ok := c.Get(FileTypeFilter).(string); ok {
				fils := strings.Split(ftf, ",")
				if len(fils) > 0 {
					fd = fd.Filter(c.GetOr(FileTypeFilterDesc, ftf).(string), fils...)

				}
			}
			if sd, ok := c.Get(fd.StartDir).(string); ok {
				fd = fd.SetStartDir(sd)
			}
			op := c.GetOr(Operation, LoadOp)
			switch op {
			case SaveOp:
				return fd.Save()
			case LoadOp:
				return fd.Load()
			}
			return nil, errors.Wrapf(err, "got (%q): expected one of (%q,%q)", op, LoadOp, SaveOp)
		},

		"dir": func(c app.Context) (path interface{}, err error) {
			return dialog.Directory().Title(c.GetOr(Title, c.App.Name).(string)).Browse()
		},
	}
}
