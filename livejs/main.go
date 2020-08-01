// Package livejs injects code to reload the current page when it (or its
// dependencies) are modified. This plugin is helpful for authoring web content
// locally, but should be disabled for site deployment. This can be achieved by
// conditionally including it using the "condition" filter.
package livejs

import (
	"fmt"
	"github.com/FooSoft/goldsmith"
	"github.com/FooSoft/goldsmith-components/filters/wildcard"
	"github.com/PuerkitoBio/goquery"
)

// LiveJs chainable context.
type LiveJs struct {
	html string
}

// New creates a new instance of the LiveJs plugin.
func New() *LiveJs {
	return new(LiveJs)
}

func (*LiveJs) Name() string {
	return "livejs"
}

func (plugin *LiveJs) Initialize(context *goldsmith.Context) (goldsmith.Filter, error) {
	plugin.html = fmt.Sprintf("\n<!-- begin livejs code -->\n<script>\n%s\n</script>\n<!-- end livejs code -->\n", js)
	return wildcard.New("**/*.html", "**/*.htm"), nil
}

func (plugin *LiveJs) Process(context *goldsmith.Context, inputFile *goldsmith.File) error {
	if outputFile := context.RetrieveCachedFile(inputFile.Path(), inputFile); outputFile != nil {
		outputFile.Meta = inputFile.Meta
		context.DispatchFile(outputFile)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(inputFile)
	if err != nil {
		return err
	}

	doc.Find("body").AppendHtml(plugin.html)

	html, err := doc.Html()
	if err != nil {
		return err
	}

	outputFile := context.CreateFileFromData(inputFile.Path(), []byte(html))
	outputFile.Meta = inputFile.Meta
	context.DispatchAndCacheFile(outputFile)
	return nil
}
