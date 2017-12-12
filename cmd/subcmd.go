package cmd

import (
	"log"

	"github.com/urfave/cli"
)

func RunSubcmd(ctx *cli.Context) error {
	log.Println("Running sub command")
	return nil
}
