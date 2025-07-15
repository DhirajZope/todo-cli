package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhirajzope/todo/internal/service"
	"github.com/spf13/cobra"
)

var listDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete List",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		svc := service.NewListService(dbConn)
		if err := svc.Delete(id); err != nil {
			return err
		}
		fmt.Println("List deleted:", id)
		return nil
	},
}

func init() {
	listCmd.AddCommand(listDeleteCmd)
}
