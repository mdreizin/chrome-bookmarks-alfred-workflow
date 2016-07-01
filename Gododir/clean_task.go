package main

import (
	"fmt"
	do "gopkg.in/godo.v2"
)

func cleanTask(dir string) func(*do.Context) {
	return func(c *do.Context) {
		c.Bash(fmt.Sprintf("rm -rf %s", dir))
	}
}
