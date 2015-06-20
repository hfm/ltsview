package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ltsview"
	app.Version = Version
	app.Usage = ""
	app.Author = "Takahiro OKUMURA"
	app.Email = "hfm.garden@gmail.com"
	app.Action = doMain
	app.Run(os.Args)
}

func doMain(c *cli.Context) {
}
