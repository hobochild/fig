// Package layout transforms content by applying Go templates to the content
// and metadata of HTML files. This plugin can be easily used with the
// "frontmatter" and "markdown" plugins to generate easy to maintain
// content-driven websites that are completely decoupled from layout details.
package layout

import (
	"bytes"
	"html/template"
	"sync"

	"github.com/FooSoft/goldsmith"
	"github.com/FooSoft/goldsmith-components/filters/wildcard"
)

// Layout chainable context.
type Layout struct {
	layoutKey     string
	contentKey    string
	defaultLayout *string
	helpers       template.FuncMap

	inputFiles    []*goldsmith.File
	templateFiles []*goldsmith.File
  cssFiles      []*goldsmith.File

	mutex         sync.Mutex

	template *template.Template
}

// New creates a new instance of the Layout plugin.
func New() *Layout {
	return &Layout{
		layoutKey:  "Layout",
		contentKey: "Content",
		helpers:    nil,
	}
}

// LayoutKey sets the metadata key used to access the layout identifier (default: "Layout").
func (plugin *Layout) LayoutKey(key string) *Layout {
	plugin.layoutKey = key
	return plugin
}

// DefaultLayout sets the name of the layout to use if none is specified.
func (plugin *Layout) DefaultLayout(name string) *Layout {
	plugin.defaultLayout = &name
	return plugin
}

// ContentKey sets the metadata key used to access the source content (default: "Content").
func (layout *Layout) ContentKey(key string) *Layout {
	layout.contentKey = key
	return layout
}

// Helpers sets the function map used to lookup template helper functions.
func (plugin *Layout) Helpers(helpers template.FuncMap) *Layout {
	plugin.helpers = helpers
	return plugin
}

func (*Layout) Name() string {
	return "layout"
}

func (plugin *Layout) Initialize(context *goldsmith.Context) (goldsmith.Filter, error) {
	plugin.template = template.New("").Funcs(plugin.helpers)
  plugin.template.Parse(defaultLayout)

	return wildcard.New("**/*.html", "**/*.htm", "**/*.tmpl", "**/*.gohtml", "**/*.css"), nil
}

func (plugin *Layout) Process(context *goldsmith.Context, inputFile *goldsmith.File) error {
	plugin.mutex.Lock()
	defer plugin.mutex.Unlock()

	switch inputFile.Ext() {
	case ".html", ".htm":
		if _, ok := plugin.getFileLayout(inputFile); ok {
			var buff bytes.Buffer
			if _, err := inputFile.WriteTo(&buff); err != nil {
				return err
			}

			inputFile.Meta[plugin.contentKey] = template.HTML(buff.Bytes())
			plugin.inputFiles = append(plugin.inputFiles, inputFile)
		} else {
			context.DispatchFile(inputFile)
		}
	case ".tmpl", ".gohtml":
		plugin.templateFiles = append(plugin.templateFiles, inputFile)
  case ".css":
    plugin.cssFiles = append(plugin.cssFiles, inputFile)
		context.DispatchFile(inputFile)
	}

	return nil
}

func (plugin *Layout) Finalize(context *goldsmith.Context) error {
	for _, templateFile := range plugin.templateFiles {
		var buff bytes.Buffer
		if _, err := templateFile.WriteTo(&buff); err != nil {
			return err
		}

		if _, err := plugin.template.Parse(string(buff.Bytes())); err != nil {
			return err
		}
	}

	for _, inputFile := range plugin.inputFiles {
		if name, ok := plugin.getFileLayout(inputFile); ok {
			var buff bytes.Buffer
      inputFile.Meta["cssFiles"] = plugin.cssFiles
			if err := plugin.template.ExecuteTemplate(&buff, name, inputFile); err != nil {
				return err
			}

			outputFile := context.CreateFileFromData(inputFile.Path(), buff.Bytes())
			outputFile.Meta = inputFile.Meta
			context.DispatchFile(outputFile)
		} else {
			context.DispatchFile(inputFile)
		}
	}

	return nil
}

func (plugin *Layout) getFileLayout(file *goldsmith.File) (string, bool) {
	if name, ok := file.Meta[plugin.layoutKey].(string); ok {
		return name, true
	}

	if plugin.defaultLayout != nil {
		return *plugin.defaultLayout, true
	}

	return "", false
}
