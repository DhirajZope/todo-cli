package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhirajzope/todo/internal/service"
	"github.com/spf13/cobra"
)

var taskAddCmd = &cobra.Command{
	Use:   "add [listID] [description]",
	Short: "Add a new task to a list",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		lid, _ := strconv.Atoi(args[0])
		svc := service.NewTaskService(dbConn)
		if err := svc.Add(lid, args[1]); err != nil {
			return err
		}
		fmt.Println("Task added to list", lid)
		return nil
	},
}

func init() {
	taskCmd.AddCommand(taskAddCmd)
}
