package cmd

import (
	"database/sql"

	"github.com/spf13/cobra"
)

var (
	dbConn  *sql.DB
	RootCmd = &cobra.Command{
		Use:   "todo",
		Short: "Simple todo CLI",
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Command for managing lists",
	}
	taskCmd = &cobra.Command{
		Use:   "task",
		Short: "Command for managing tasks",
	}
)

func SetDB(conn *sql.DB) {
	dbConn = conn
}

func GetDB() *sql.DB {
	return dbConn
}

func init() {
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(taskCmd)
}
