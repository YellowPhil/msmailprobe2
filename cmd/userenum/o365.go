package userenum

import "github.com/spf13/cobra"

var (
	email     string
	emailFile string
	O365Cmd   = &cobra.Command{
		Use:   "o365",
		Short: "Enumerate using i365 resources",
		Run:   runO365,
	}
)

func runO365(cmd *cobra.Command, args []string) {
}

func init() {
	O365Cmd.PersistentFlags().StringVarP(&email, "email-file", "E", "", "Path to the file with emails")
	O365Cmd.PersistentFlags().StringVarP(&emailFile, "email", "e", "", "Single email to check")
	O365Cmd.MarkFlagsOneRequired("email", "email-file")
	O365Cmd.MarkFlagFilename("email-file")
	O365Cmd.MarkFlagsMutuallyExclusive("email", "email-file")
}
