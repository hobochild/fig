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
          <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/hobochild/style.min.css" />
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
      {{if and (eq .Name "index.html") (eq .Dir ".") }} {{ else }}<a href="/">
      <svg id="i-arrow-left" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 50 28" width="32" height="14"  zIndex="1" fill="none" stroke="currentcolor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2">
        <path d="M10 6 L2 16 10 26 M2 16 L30 16" />
    </svg>Home</a>{{end}}
      <h1>{{.Meta.Title}}</h1>
      {{.Meta.Content}}
      </main>
      {{template "__footer__" .}}
  {{end}}
  `
