package main

import (
	"fmt"
	do "gopkg.in/godo.v2"
)

func ensureDir(c *do.Context, dir string) {
	c.Bash(fmt.Sprintf("mkdir -p %s", dir))
}
