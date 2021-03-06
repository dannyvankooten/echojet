echojet
=========

Echojet wires the [cloudykit/jet](https://github.com/CloudyKit/jet) templating package with [Echo](https://echo.labstack.com/)'s `c.HTML(..)` rendering function.

Additionally, it helps you load your templates from a bindata generated source.

# Usage

```go
import "github.com/dannyvankooten/echojet"
```

```go
e := echo.New()
e.Renderer = echojet.New(echojet.Options{
   Directory: "templates/", // Path from current working dir
   DevelopmentMode: true,
})
```

You can then use `c.HTML(header, "template-name.jet", nil)` in your Echo request handlers to render your Jet templates.

## Bindata

To use echojet with [go-bindata](https://github.com/jteeuwen/go-bindata), supply the included `BinLoader` as the `Loader` option when instantiating your renderer.

```go
e := echo.New()
e.Renderer = echojet.New(echojet.Options{
   Loader: &echojet.Loader{
     Root:      "templates/", 
     AssetFunc: bindata.Asset,
   },
   DevelopmentMode: true,
})
```

# License
MIT licensed.
