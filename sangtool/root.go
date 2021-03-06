package main

import (
	"com.samderlust/sangtoolbox/sangtool/command"
	"github.com/spf13/cobra"
)

func rootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sangtool",
		Short:   "Tools to help you push up progress",
		Version: "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
			}

			return nil
		},
	}

	cmd.SetVersionTemplate("sangtool CLI v{{.Version}}\n")
	cmd.AddCommand(command.FlutterCmd())

	return cmd
}
