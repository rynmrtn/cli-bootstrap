package main

import (
	"github.com/rynmrtn/cli-bootstrap/cmd"
	"os"
)

func main() {
	app = cmd.App()
	app.Run(os.Args)
}
