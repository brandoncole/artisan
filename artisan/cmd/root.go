package cmd

import (
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "artisan",
    Short: "Artisan generates frontends and backends from Protocol Buffers",
    Long: "Artisan generates frontends and backends from Protocol Buffers",
    Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
    },
}