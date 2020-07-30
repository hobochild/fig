package layout

var defaultLayout = `
  {{define "__header__"}}
  <!DOCTYPE html>
  <html lang="en">
      <head>
          <meta charset="utf-8">
          <meta name="viewport" content="width=device-width, initial-scale=1.0">
          <meta name="description" content="{{.Meta.Description}}">
          <title>{{.Meta.Title}}</title>
          <link href="https://classic.yarnpkg.com/en/package/normalize.css" rel="stylesheet" type="text/css"></link>
          <link href="https://cdn.jsdelivr.net/gh/hobochild/fig/style.css" rel="stylesheet" type="text/css"></link>
          {{range .Meta.cssFiles }}
            <link href="/{{.Path}}" rel="stylesheet" type="text/css"></style>
          {{end}}
      </head>
      <body>
  {{end}}

  {{define "__footer__"}}
      </body>
  </html>
  {{end}}

  {{define "__basic__"}}
      {{template "__header__" .}}
      <main class="main">
      {{if and (eq .Name "index.html") (eq .Dir ".") }} {{ else }}<a href="/">Home</a>{{end}}
      <h1>{{.Meta.Title}}</h1>
      {{.Meta.Content}}
      </main>
      {{template "__footer__" .}}
  {{end}}
  `
