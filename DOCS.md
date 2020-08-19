## Customize template

Create a template file.

```go
// ./content/custom.gohtml
{{define "basic"}}
    <main class="main">
    <h1>{{.Meta.Title}}</h1>
    <p>This is a custom template</p>
    {{.Meta.Content}}
    </main>
{{end}}
```

Now update you page to use the Layout:

```markdown
+++
Layout = "basic"
Title = "Using a custom layout"
+++

Here is some content!
```
