package main

import "os"

func main() {
	run()
}

func run() error {
	app := newApp()

	return app.Run(os.Args)
}
