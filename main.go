package main

import (
	"errors"
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

	err = gToolbox.CheckToolbox([]string{"adb"})
	if err != nil {
		return err
	}

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
		return errors.New("you need to provide the apps file using -f or select all apps using -a")
	}

	for _, app := range apps {
		if cliClear {
			err = PMClear(app)
			if err != nil {
				fmt.Println(err)
			}
		}
		if cliUninstall {
			err = PMUninstall(app)
			if err != nil {
				fmt.Println(err)
			}
		}
		if cliUninstallUser {
			err = PMUninstallUser(app, cliUID)
			if err != nil {
				fmt.Println(err)
			}
		}
		if cliReinstall {
			err = PMReinstall(app)
			if err != nil {
				fmt.Println(err)
			}
		}
		if cliDisable {
			err = PMDisable(app)
			if err != nil {
				fmt.Println(err)
			}
		}
		if cliDisableUser {
			err = PMDisableUser(app, cliUID)
			if err != nil {
				fmt.Println(err)
			}
		}
		if cliEnable {
			err = PMEnable(app)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
