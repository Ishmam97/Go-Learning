package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "gifm",
		Short:        "Hello Ishmam",
		SilenceUsage: true,
	}
	cmd.AddCommand(printTimeCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
