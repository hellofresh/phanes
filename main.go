package main

import (
	"os"
	"sort"

	"github.com/fatih/color"
	"github.com/hellofresh/phanes/provider"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "phanes"
	app.Usage = "Creates a new oauth client for a provide"
	app.Author = "Italo Lelis de Vietro"
	app.Copyright = "HelloFresh SE"
	app.Email = "il@hellofresh.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "provider, p",
			Value: "hellofresh",
			Usage: "The authentication provider to create the oauth client",
		},
		cli.StringFlag{
			Name:   "provider-url, u",
			Usage:  "The identity provider client endpoint",
			EnvVar: "PROVIDER_URL",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Creates a new oauth client",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "name, n",
					Usage:  "The client name",
					EnvVar: "CLIENT_NAME",
				},
				cli.StringFlag{
					Name:  "redirect-url, r",
					Value: "http://localhost",
					Usage: "The callback url",
				},
			},
			Action: func(c *cli.Context) error {
				provider := provider.Create(c.GlobalString("provider-url"), c.GlobalString("provider"))
				if provider == nil {
					color.Red("Invalid provider selected, please choose a valid provider")
					return nil
				}

				client, err := provider.Create(c.String("name"), c.String("redirect-url"))
				if err != nil {
					return err
				}

				color.Green("Credentials for %s created!", c.String("name"))
				color.Cyan("Client ID: %s", client.GetID())
				color.Cyan("Client Secret: %s", client.GetSecret())

				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Deletes an oauth client by ID",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "provider, p",
					Value: "hellofresh",
					Usage: "The authentication provider to create the oauth client",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "The client's ID",
				},
			},
			Action: func(c *cli.Context) error {
				provider := provider.Create(c.GlobalString("provider-url"), c.GlobalString("provider"))
				if provider == nil {
					color.Red("Invalid provider selected, please choose a valid provider")
					return nil
				}

				err := provider.Delete(c.String("id"))
				if err != nil {
					return err
				}

				color.Green("Client %s was deleted!", c.String("id"))

				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	app.Run(os.Args)
}
