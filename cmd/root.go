package cmd

import (
	"github.com/InvalidJokerDE/PterodactylSecurity/config"
	"github.com/InvalidJokerDE/PterodactylSecurity/http"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "psecurity",
	Short: "Runs the API server for the Pterodactyl Security System.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := config.Create(); err != nil {
			log.Fatalf("failed to create config: %s", err)
		}
	},
	Run: rootCmdRun,
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatalf("failed to execute root command: %s", err)
	}
}

func init() {
	rootCommand.AddCommand(newDiagnosticsCommand())
}

func rootCmdRun(cmd *cobra.Command, _ []string) {
	log.Info("testing panel connection")

	cfg := config.Get()

	client := http.New(cfg.Panel.URL, cfg.Panel.Key, cfg.Panel.ID)
	if st, err := client.TestConnection(); err != nil {
		log.Error("%s (status: %d)", err, st)
		log.Error("make sure that your credentials are correct")
		return
	}
}
