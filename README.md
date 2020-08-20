# lil

> A lil static site generator.

I wanted a dead simple generator that produced a pretty but readable site from a set of markdown files with a pleasant dev expierence. I didn't want any config files, taxonomy systems or javascript frameworks.

Basically you chuck some markdown files in a directory and run `lil`, not much more too it.

There are a few things baked in.

- Live reloads in dev mode.
- Automatic code styling.
- Basic overwriteable [stylesheet](./style.css).

## Install

You can download the [latest release](https://github.com/hobochild/lil/releases) or use `go get github.com/hobochild/lil`

## Usage:

Create a page

```bash
echo "# Hello World!" > content/index.md
```

If you want to overwrite the default styles put a css file in the content directory.

```bash
echo "body { background: red; }" > content/style.css
```

You can also [customize the templates](/DOCS.md#customize-template) if you like.

## Dev mode:

```
lil -d
```

## Build for prod

Check the `./build` for your static site.

```
lil
```

Built with [goldsmith](https://github.com/FooSoft/goldsmith)
