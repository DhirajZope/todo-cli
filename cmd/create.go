package cmd

import (
	"fmt"

	"github.com/dhirajzope/todo/internal/service"
	"github.com/spf13/cobra"
)

var listCreatedcmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := service.NewListService(dbConn)
		if err := svc.Create(args[0]); err != nil {
			return err
		}
		fmt.Println("List Created:", args[0])
		return nil
	},
}

func init() {
	listCmd.AddCommand(listCreatedcmd)
}
