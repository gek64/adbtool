package main

import (
	"fmt"
	"github.com/gek64/gek/gToolbox"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	cliClear         bool
	cliUninstall     bool
	cliUninstallUser bool
	cliReinstall     bool
	cliDisable       bool
	cliDisableUser   bool
	cliEnable        bool
	cliAll           bool
	cliUID           int
	cliFile          string
)

func init() {
	err := gToolbox.CheckToolbox([]string{"adb"})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	cmds := []*cli.Command{
		{
			Name:  "clear",
			Usage: "clear apps data",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "all",
					Aliases:     []string{"a"},
					Usage:       "select all apps in device",
					Destination: &cliAll,
				},
				&cli.StringFlag{
					Name:        "file",
					Aliases:     []string{"f"},
					Usage:       "use all apps in a file",
					Destination: &cliFile,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				cliClear = true
				return run()
			},
		},
		{
			Name:  "uninstall",
			Usage: "uninstall apps",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "all",
					Aliases:     []string{"a"},
					Usage:       "select all apps in device",
					Destination: &cliAll,
				},
				&cli.StringFlag{
					Name:        "file",
					Aliases:     []string{"f"},
					Usage:       "use all apps in a file",
					Destination: &cliFile,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				cliUninstall = true
				return run()
			},
		},
		{
			Name:  "uninstall-user",
			Usage: "uninstall apps for user",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:        "uid",
					Aliases:     []string{"u"},
					Usage:       "set user id",
					Required:    true,
					Destination: &cliUID,
				},
				&cli.BoolFlag{
					Name:        "all",
					Aliases:     []string{"a"},
					Usage:       "select all apps in device",
					Destination: &cliAll,
				},
				&cli.StringFlag{
					Name:        "file",
					Aliases:     []string{"f"},
					Usage:       "use all apps in a file",
					Destination: &cliFile,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				cliUninstallUser = true
				return run()
			},
		},
		{
			Name:  "reinstall",
			Usage: "reinstall apps",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "all",
					Aliases:     []string{"a"},
					Usage:       "select all apps in device",
					Destination: &cliAll,
				},
				&cli.StringFlag{
					Name:        "file",
					Aliases:     []string{"f"},
					Usage:       "use all apps in a file",
					Destination: &cliFile,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				cliReinstall = true
				return run()
			},
		},
		{
			Name:  "disable",
			Usage: "disable apps",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "all",
					Aliases:     []string{"a"},
					Usage:       "select all apps in device",
					Destination: &cliAll,
				},
				&cli.StringFlag{
					Name:        "file",
					Aliases:     []string{"f"},
					Usage:       "use all apps in a file",
					Destination: &cliFile,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				cliDisable = true
				return run()
			},
		},
		{
			Name:  "disable-user",
			Usage: "disable apps for user",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:        "uid",
					Aliases:     []string{"u"},
					Usage:       "set user id",
					Required:    true,
					Destination: &cliUID,
				},
				&cli.BoolFlag{
					Name:        "all",
					Aliases:     []string{"a"},
					Usage:       "select all apps in device",
					Destination: &cliAll,
				},
				&cli.StringFlag{
					Name:        "file",
					Aliases:     []string{"f"},
					Usage:       "use all apps in a file",
					Destination: &cliFile,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				cliDisableUser = true
				return run()
			},
		},
		{
			Name:  "enable",
			Usage: "enable apps",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "all",
					Aliases:     []string{"a"},
					Usage:       "select all apps in device",
					Destination: &cliAll,
				},
				&cli.StringFlag{
					Name:        "file",
					Aliases:     []string{"f"},
					Usage:       "use all apps in a file",
					Destination: &cliFile,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				cliEnable = true
				return run()
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s", cCtx.App.Version)
	}

	app := &cli.App{
		Usage:    "ADB Batch Tool",
		Version:  "v1.00",
		Commands: cmds,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	var apps []string

	if cliAll {
		apps, err = GetAppListFromADB()
		if err != nil {
			log.Panicln(err)
		}
	} else if cliFile != "" {
		apps, err = GetAppsFromFile(cliFile)
		if err != nil {
			log.Panicln(err)
		}
	} else {
		log.Fatalln("you need to provide the apps file using -f or select all apps using -a")
	}

	for _, app := range apps {
		if cliClear {
			return PMClear(app)
		}
		if cliUninstall {
			return PMUninstall(app)
		}
		if cliUninstallUser {
			return PMUninstallUser(app, cliUID)
		}
		if cliReinstall {
			return PMReinstall(app)
		}
		if cliDisable {
			return PMDisable(app)
		}
		if cliDisableUser {
			return PMDisableUser(app, cliUID)
		}
		if cliEnable {
			return PMEnable(app)
		}
	}
	return nil
}
