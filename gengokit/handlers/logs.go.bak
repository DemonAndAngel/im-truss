package handlers

import (
	"io"
	"github.com/pkg/errors"


	"github.com/metaverse/truss/gengokit"
	"github.com/metaverse/truss/gengokit/handlers/templates"
)

const LogPath = "handlers/logs.gotemplate"

// NewHook returns a new HookRender
func NewLog() *Log {
	var l Log
	return &l
}

type Log struct {
	prev io.Reader
}
// Load loads the previous version of the middleware file.
func (l *Log) Load(prev io.Reader) {
	l.prev = prev
}

func (m *Log) Render(path string, data *gengokit.Data) (io.Reader, error) {
	if path != LogPath {
		return nil, errors.Errorf("cannot render unknown file: %q", path)
	}
	if m.prev != nil {
		return m.prev, nil
	}
	return data.ApplyTemplate(templates.Log, "Log")
}
