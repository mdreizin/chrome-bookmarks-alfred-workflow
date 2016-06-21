package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("default", do.Series{"assets", "generate", "compress"}, nil)

	p.Task("assets", nil, func(c *do.Context) {})

	p.Task("generate", nil, func(c *do.Context) {})

	p.Task("compress", nil, func(c *do.Context) {})
}

func main() {
	do.Godo(tasks)
}
