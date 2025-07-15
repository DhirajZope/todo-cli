package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhirajzope/todo/internal/service"
	"github.com/spf13/cobra"
)

var taskDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		svc := service.NewTaskService(dbConn)
		if err := svc.Delete(id); err != nil {
			return err
		}
		fmt.Println("Task deleted:", id)
		return nil
	},
}

func init() {
	taskCmd.AddCommand(taskDeleteCmd)
}
