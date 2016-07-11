package main

import (
	"github.com/BuiltLabs/maintenance"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	m := maintenance.NewMaintenance()

	app.Commands = []cli.Command{
		{
			Name:  "poll",
			Usage: "polls remote state and creates/deletes maintenance file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Value: "/tmp/maintenance.enable",
					Usage: "Location to create/delete `FILE` for maintenance events",
				},
				cli.StringFlag{
					Name:  "table, t",
					Value: "maintenanceFlags",
					Usage: "Table where maintenance state is stored",
				},
				cli.StringFlag{
					Name:  "key, k",
					Value: "ops",
					Usage: "Primary lookup key to query maintenance data",
				},
				cli.StringFlag{
					Name:  "keyName, n",
					Value: "environment",
					Usage: "Primary lookup key column name",
				},
			},
			Action: func(c *cli.Context) error {
				m.FileTarget = c.String("file")
				m.TableName = c.String("table")
				m.Key = c.String("key")
				m.KeyName = c.String("keyName")

				m.PollStatus()

				return nil
			},
		},
		{
			Name:  "enable",
			Usage: "flags maintenance mode as enabled in dynamodb",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "table, t",
					Value: "maintenanceFlags",
					Usage: "Table where maintenance state is stored",
				},
				cli.StringFlag{
					Name:  "keyName, n",
					Value: "environment",
					Usage: "Primary lookup key column name",
				},
				cli.StringFlag{
					Name:  "key, k",
					Value: "ops",
					Usage: "Primary lookup key to query maintenance data",
				},
				cli.StringFlag{
					Name:  "metaFile, m",
					Usage: "`FILE` location containing metadata that will be injected into maintenance file (optional)",
				},
			},
			Action: func(c *cli.Context) error {
				m.TableName = c.String("table")
				m.KeyName = c.String("keyName")
				m.Key = c.String("key")

				if metaFile := c.String("metaFile"); metaFile != "" {
					m.ImportMetaData(metaFile)
				}

				m.EnableMaintenance()
				return nil
			},
		},
		{
			Name:  "disable",
			Usage: "flags maintenance mode as disabled in dynamodb",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "table, t",
					Value: "maintenanceFlags",
					Usage: "Table where maintenance state is stored",
				},
				cli.StringFlag{
					Name:  "keyname, n",
					Value: "environment",
					Usage: "Primary lookup key column name",
				},
				cli.StringFlag{
					Name:  "key, k",
					Value: "ops",
					Usage: "Primary lookup key to query maintenance data",
				},
			},
			Action: func(c *cli.Context) error {
				m.TableName = c.String("table")
				m.KeyName = c.String("keyname")
				m.Key = c.String("key")

				m.DisableMaintenance()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
