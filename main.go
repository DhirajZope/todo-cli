package main

import (
	"fmt"
	"os"

	"github.com/dhirajzope/todo/cmd"
	"github.com/dhirajzope/todo/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database for commands that need it
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "todo.db"
	}

	// Only initialize DB if not a help/version command
	if !isHelpOrVersionCommand() {
		conn, err := db.Init(dbPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open DB: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close()
		cmd.SetDB(conn)
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func isHelpOrVersionCommand() bool {
	if len(os.Args) < 2 {
		return false
	}

	arg := os.Args[1]
	return arg == "version" || arg == "help" || arg == "--help" || arg == "-h" || arg == "--version"
}
