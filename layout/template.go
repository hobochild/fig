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
          <style src="https://classic.yarnpkg.com/en/package/normalize.css"></style>
          <style src="https://cdn.jsdelivr.net/gh/hobochild/fig/style.css">
          {{range .Meta.cssFiles }}    
            <style src="/{{.Path}}"></style>
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
      {{if and (eq .Name "index.html") (eq .Dir ".") }} {{ else }}<a href="/">Home</a>{{end}}
      {{template "__footer__" .}}
  {{end}}
  `
