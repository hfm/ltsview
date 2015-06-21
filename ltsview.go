package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "OKUMURA Takahiro"
	app.Email = "hfm.garden@gmail.com"
	app.Usage = "Labeled Tab Separated Value manipulator"

	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "key, k",
			Usage: "Print specified keys (separated by ',')",
		},
		cli.StringFlag {
			Name: "nocolor, n",
			Usage: "Print without color",
		},
	}

	app.Action = parseLine
	app.Run(os.Args)
}

func parseLine(c *cli.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		// fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin:", err)
	}
}

// func parseLtsv(line string) map[string]string {
// 	line := make(map[string]string)
// }
