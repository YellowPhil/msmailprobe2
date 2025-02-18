package cmd

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/yellowphil/msmailprobe2/cmd/examples"
	"github.com/yellowphil/msmailprobe2/cmd/identify"
	"github.com/yellowphil/msmailprobe2/cmd/userenum"
)

var Verbose bool

func init() {
	rootCmd.AddCommand(examples.ExamplesCmd, identify.IdentifyCmd, userenum.UserenumCmd)
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "msmailprobe",
	Short: "msmailprobe",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if Verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	},
}
