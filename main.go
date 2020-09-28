package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	version           = "v.0.1"
	defaultConfigPath = os.Getenv("HOME") + "/.config/comptemp/"
	configFilename    = "config"
	tmpconfigFilename = "tmp"
	configViper       = viper.New()
	tmpConfigViper    = viper.New()
	config            Config
	tmpConfig         TmpConfig
)

func handleUsageError(c *cli.Context, err error, _ bool) error {
	fmt.Println("Incorrect Usage : ", err.Error())
	cli.ShowAppHelp(c)
	return cli.Exit("", 1)
}

func readTempConfig() error {
	tmpConfigViper.SetConfigName(tmpconfigFilename)
	tmpConfigViper.AddConfigPath(defaultConfigPath)

	if err := tmpConfigViper.ReadInConfig(); err != nil {
		return err
	}
	err := tmpConfigViper.UnmarshalKey("config", &tmpConfig)
	if err != nil {
		return err
	}
	return nil
}

func build() error {
	cmd := exec.Command("g++", tmpConfig.Filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println("g++", tmpConfig.Filename)
	if err != nil {
		return err
	}
	return nil
}

func run() error {
	err := exec.Command("./a.out").Run()
	if err != nil {
		return err
	}
	return nil
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
				var temppath = defaultConfigPath + config[lang].TemplatePath
				var filename = "./" + c.Args().First()
				if len(filename) == 2 {
					return cli.Exit("specify filename", 1)
				}

				err := exec.Command("cp", temppath, filename).Run()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}
				fmt.Println("cp", temppath, filename)

				var tmpConfig = TmpConfig{Language: lang, Filename: filename}
				tmpConfigViper.SetDefault("config", tmpConfig)
				err = tmpConfigViper.WriteConfigAs(defaultConfigPath + tmpconfigFilename)
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}
				fmt.Println("write tmpConfig to", defaultConfigPath+tmpconfigFilename)

				return nil
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build",
			Action: func(c *cli.Context) error {
				err := readTempConfig()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}

				err = build()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}

				return nil
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "build & run",
			Action: func(c *cli.Context) error {
				err := readTempConfig()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}

				err = build()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}

				err = run()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}

				return nil
			},
		},
	}

	configViper.SetConfigType("yml")
	tmpConfigViper.SetConfigType("yml")
	configViper.SetConfigName(configFilename)
	configViper.AddConfigPath(defaultConfigPath)

	if err := configViper.ReadInConfig(); err != nil {
		// TODO
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	err := configViper.Unmarshal(&config)
	if err != nil {
		// TODO
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}
