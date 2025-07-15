package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhirajzope/todo/internal/service"
	"github.com/spf13/cobra"
)

var taskCompleteCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		svc := service.NewTaskService(dbConn)
		if err := svc.Complete(id); err != nil {
			return err
		}
		fmt.Println("Task completed:", id)
		return nil
	},
}

func init() {
	taskCmd.AddCommand(taskCompleteCmd)
}
