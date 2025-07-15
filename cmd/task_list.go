package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhirajzope/todo/internal/service"
	"github.com/spf13/cobra"
)

var taskListCmd = &cobra.Command{
	Use:   "list [listID]",
	Short: "List tasks for a list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		lid, _ := strconv.Atoi(args[0])
		svc := service.NewTaskService(dbConn)
		tasks, err := svc.ListBy(lid)
		if err != nil {
			return err
		}
		for _, t := range tasks {
			status := " "
			if t.Completed {
				status = "✔"
			}
			fmt.Printf("%d: [%s] %s\n", t.ID, status, t.Description)
		}
		return nil
	},
}

func init() {
	taskCmd.AddCommand(taskListCmd)
}
