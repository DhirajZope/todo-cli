package cmd

import (
	"fmt"

	"github.com/dhirajzope/todo/internal/service"
	"github.com/spf13/cobra"
)

var listListCmd = &cobra.Command{
	Use:   "all",
	Short: "Show all lists",
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := service.NewListService(dbConn)
		lists, err := svc.All()
		if err != nil {
			return err
		}
		for _, l := range lists {
			fmt.Printf("%d: %s (created %s)\n", l.ID, l.Name, l.CreatedAt.Format("2006-01-02"))
		}
		return nil
	},
}

func init() {
	listCmd.AddCommand(listListCmd)
}
