# lil

> A lil static site generator.

I wanted a dead simple generator that produced a pretty but readable site from a set of markdown files and has a pleasant dev expierence. I didn't want any config files, taxonomy systems or javascript frameworks.

# Install

You download the [latest release](https://github.com/hobochild/lil/releases) or use `go get github.com/hobochild/lil`

# Usage:

Create a page

```bash
echo "# Hello World!" > content/index.md
```

If you want to overwrite the default styles put a css file in the content directory.

```bash
echo "body { background: red; }" > content/style.css
```

You can also [customize the templates](/docs.md#customize-template) if you like.

# Dev mode:

```
lil -d
```

# Build for prod

```
lil
```

Built on [goldsmith](https://github.com/FooSoft/goldsmith)
