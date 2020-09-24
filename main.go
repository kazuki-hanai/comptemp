package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	version            = "v.0.1"
	defaultConfigPaths = os.Getenv("HOME") + "/.config/comptemp"
	config             Config
)

func handleUsageError(c *cli.Context, err error, _ bool) error {
	fmt.Println("%s %s", "Incorrect Usage.", err.Error())
	cli.ShowAppHelp(c)
	return cli.Exit("", 1)
}

func main() {
	app := cli.NewApp()
	app.Name = "cpt"
	app.Usage = "useful tools for competition programming(e.g. generate templete, build and run)."
	app.Version = "v.0.1"
	app.OnUsageError = handleUsageError
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "hnkz",
			Email: "hanakazu8989@gmail.com",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "generate tenplate",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "lang",
					Aliases: []string{"l"},
					Value:   "cpp",
					Usage:   "specify language(e.g. cpp, rust, python)",
				},
			},
			Action: func(c *cli.Context) error {
				var lang = c.String("lang")
				var temppath = defaultConfigPaths + "/" + config[lang].TempPath
				var filename = "./" + c.Args().First()
				if len(filename) == 0 {
					return cli.Exit("specify filename", 1)
				}

				err := exec.Command("cp", temppath, filename).Run()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}
				fmt.Println("cp", temppath, filename)

				return nil
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "build & run",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(defaultConfigPaths)

	if err := viper.ReadInConfig(); err != nil {
		// TODO
		os.Exit(1)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		// TODO
		os.Exit(1)
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
