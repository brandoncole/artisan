package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Shows version information",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Version")
    },
}

func init() {
    RootCmd.AddCommand(versionCmd)
}
