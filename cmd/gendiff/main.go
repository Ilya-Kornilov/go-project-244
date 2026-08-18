package main

import (
	"context"
	"fmt"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
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

			formatFlag := cmd.String("format")
			res, err := code.GenDiff(args[0], args[1], formatFlag)
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
