package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Short: "Migrates an artisan backend",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Migrate")
    },
}

func init() {
    RootCmd.AddCommand(migrateCmd)
}
