/*
MIT License

Copyright (c) 2018 Danny van Kooten

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
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
