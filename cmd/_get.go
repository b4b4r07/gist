package cmd

import (
	"fmt"

	"github.com/b4b4r07/gist/cli"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Manipulate gist with the command passed in the argument",
	Long:  "Manipulate gist with the command passed in the argument",
	RunE:  get,
}

func get(cmd *cobra.Command, args []string) error {
	screen, err := cli.NewScreen()
	if err != nil {
		return err
	}

	rows, err := screen.Select()
	if err != nil {
		return err
	}

	for _, row := range rows {
		if len(args) == 0 {
			fmt.Println(row.Path)
			continue
		}
		if err := cli.Run(args[0], row.Path); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	RootCmd.AddCommand(getCmd)
}
