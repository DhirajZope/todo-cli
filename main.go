package main

import (
	"fmt"
	"os"

	"github.com/dhirajzope/todo/cmd"
	"github.com/dhirajzope/todo/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn, err := db.Init("todo.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open DB: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	cmd.SetDB(conn)

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
