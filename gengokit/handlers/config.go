package handlers

import (
	"io"

	"github.com/pkg/errors"

	"github.com/metaverse/truss/gengokit"
	"github.com/metaverse/truss/gengokit/handlers/templates"


	"strings"
)

// MiddlewaresPath is the path to the middleware gotemplate file.
const ConfigPath = "config/config.gotemplate"
const ConfigAppJsonExamplePath = "config/app.json.example"
const ConfigAppJsonPath = "config/app.json"

// NewMiddlewares returns a Renderable that renders the middlewares.go file.
func NewConfig() *Config {
	var m Config

	return &m
}

// Middlewares satisfies the gengokit.Renderable interface to render
// middlewares.
type Config struct {
	prev io.Reader
}

// Load loads the previous version of the middleware file.
func (m *Config) Load(prev io.Reader) {
	m.prev = prev
}

// Render creates the middlewares.go file. With no previous version it renders
// the templates, if there was a previous version loaded in, it passes that
// through.
func (m *Config) Render(path string, data *gengokit.Data) (io.Reader, error) {
	if path != ConfigPath && path != ConfigAppJsonExamplePath && path != ConfigAppJsonPath {
		return nil, errors.Errorf("cannot render unknown file: %q", path)
	}
	if m.prev != nil {
		return m.prev, nil
	}
	if (strings.Contains(path, "example")) {
		return data.ApplyTemplate(templates.ConfigAppJsonExample, "ConfigAppJsonExample")
	}else if (strings.Contains(path, "json")){
		return data.ApplyTemplate(templates.ConfigAppJson, "ConfigAppJson")
	}else{
		return data.ApplyTemplate(templates.Config, "Config")
	}
}
