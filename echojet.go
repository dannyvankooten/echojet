package echojet

import (
	"github.com/CloudyKit/jet"
	"io"

	"github.com/labstack/echo"
)

type Map map[string]interface{}

type Renderer struct {
	Templates *jet.Set
}

type Options struct {
	Loader          jet.Loader
	Directory       string
	DevelopmentMode bool
}

func New(o Options) *Renderer {
	r := &Renderer{}

	if o.Loader != nil {
		r.Templates = jet.NewHTMLSetLoader(o.Loader)
	} else {
		r.Templates = jet.NewHTMLSet(o.Directory)
	}

	r.Templates.SetDevelopmentMode(o.DevelopmentMode)
	return r
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, err := r.Templates.GetTemplate(name)
	if err != nil {
		return err
	}

	// convert Map into jet.VarMap
	vars := make(jet.VarMap)

	// Add global methods if data is a map
	if datamap, ok := data.(map[string]interface{}); ok {
		for k := range datamap {
			vars.Set(k, datamap[k])
		}
	}
	if datamap, ok := data.(Map); ok {
		for k := range datamap {
			vars.Set(k, datamap[k])
		}
	}

	// render template
	if err = t.Execute(w, vars, nil); err != nil {
		return err
	}

	return nil
}
