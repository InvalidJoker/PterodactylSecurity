package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func newDiagnosticsCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "diagnostics",
		Short: "Collect diagnostics information for the security system.",
		Run:   diagnosticsCmdRun,
	}
	return command
}

func diagnosticsCmdRun(*cobra.Command, []string) {
	log.Info("Test...")
}
