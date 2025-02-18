package userenum

import (
	"github.com/spf13/cobra"
)

var (
	target      string
	threads     int
	rpm         int
	rps         int
	output      string
	UserenumCmd = &cobra.Command{
		Use:   "userenum",
		Short: "Enumerate users",
	}
)

func init() {
	UserenumCmd.AddCommand(O365Cmd, OnpremCmd)
	UserenumCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "Host pointing to the exchange service")
	UserenumCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Path to the output file")
	UserenumCmd.PersistentFlags().IntVar(&threads, "threads", 5, "Path to the output file")
	UserenumCmd.PersistentFlags().IntVar(&rpm, "rpm", 0, "Max requests per minute (0 for unlimited)")
	UserenumCmd.PersistentFlags().IntVar(&rps, "rps", 0, "Max requests per second (0 for unlimited)")
	UserenumCmd.MarkFlagRequired("target")
}
