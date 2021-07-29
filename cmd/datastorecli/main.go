package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/akm/datastorecli"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "project-id",
			},
			&cli.StringFlag{
				Name: "namespace",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "query",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "offset",
						Value: 0,
					},
					&cli.IntFlag{
						Name:  "limit",
						Value: 100,
					},
					&cli.BoolFlag{
						Name: "keys-only",
					},
				},
				ArgsUsage: "KIND",
				Action: func(c *cli.Context) error {
					client, err := newClient(c)
					if err != nil {
						return err
					}
					if c.Bool("keys-only") {
						if d, err := client.QueryKeys(context.Background(), c.Int("offset"), c.Int("limit")); err != nil {
							return err
						} else {
							return formatStrings(d)
						}
					} else {
						if d, err := client.QueryData(context.Background(), c.Int("offset"), c.Int("limit")); err != nil {
							return err
						} else {
							return formatArray(d)
						}
					}

				},
			},
			{
				Name:      "get",
				ArgsUsage: "KIND",
				Action: func(c *cli.Context) error {
					client, err := newClient(c)
					if err != nil {
						return err
					}
					if d, err := client.Get(context.Background(), c.Args().Get(0)); err != nil {
						return err
					} else {
						return formatData(d)
					}
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func newClient(c *cli.Context) (*datastorecli.Client, error) {
	if c.Args().Len() < 1 {
		return nil, errors.Errorf("kind is required")
	}
	kind := c.Args().First()
	return datastorecli.NewClient(c.String("project-id"), c.String("namespace"), kind), nil
}

func formatData(d interface{}) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(b); err != nil {
		return err
	}
	return nil
}

func formatArray(d *[]interface{}) error {
	for _, i := range *d {
		if err := formatData(i); err != nil {
			return err
		}
	}
	return nil
}

func formatStrings(d *[]string) error {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(*d, "\n"))
	return nil
}
