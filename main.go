package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var (
	Version string
)

func main() {
	if _, err := os.Stat("/run/drone/env"); err == nil {
		godotenv.Overload("/run/drone/env")
	}

	app := cli.NewApp()
	app.Name = "synochat plugin"
	app.Usage = "synochat plugin"
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "synochat token",
			EnvVar: "PLUGIN_TOKEN,SYNOCHAT_TOKEN,INPUT_TOKEN",
		},
		cli.StringSliceFlag{
			Name:   "url",
			Usage:  "synochat server url",
			EnvVar: "PLUGIN_URL,SYNOCHAT_URL,INPUT_URL",
		},
		cli.StringFlag{
			Name:   "message",
			Usage:  "send synochat message",
			EnvVar: "PLUGIN_MESSAGE,SYNOCHAT_MESSAGE,INPUT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "skipssl",
			Usage:  "skip ssl certificate check",
			EnvVar: "PLUGIN_SKIPSSL,SYNOCHAT_SKIPSSL,INPUT_SKIPSSL",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Url:     c.String("url"),
		Token:   c.String("token"),
		Message: c.String("message"),
		SkipSSL: c.Bool("skipssl"),
	}
	return plugin.Exec()
}
