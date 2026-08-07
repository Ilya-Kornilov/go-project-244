package main

import (
	"context"
	"fmt"
	"os"

	// "code/internal/parser"
	"code/internal/differ"

	"github.com/urfave/cli/v3"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func main() {
	// Prevent unused-import errors while keeping libs in go.mod
	_ = assert.Equal
	var n yaml.Node
	_ = n

	app := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
		    &cli.StringFlag{
				Name: "format",
				Aliases: []string{"f"},
				Value: "stylish",
				Usage: "output format",
		    },
		},
		ArgsUsage: "file1 file2",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args().Slice()

			res, err := differ.GenDiff(args[0], args[1])
			if err != nil {
				return err
			}

			fmt.Println(res)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
