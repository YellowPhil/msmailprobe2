package userenum

import "github.com/spf13/cobra"

var (
	username     string
	usernameFile string
	OnpremCmd    = &cobra.Command{
		Use:   "onprem",
		Short: "Time-based onprem enumeration",
		Run:   runOnprem,
	}
)

func runOnprem(cmd *cobra.Command, args []string) {
}

func init() {
	OnpremCmd.Flags().StringVarP(&usernameFile, "user-file", "U", "", "Path to the file with usernames")
	OnpremCmd.Flags().StringVarP(&username, "user", "u", "", "Single username to check")
	OnpremCmd.MarkFlagsOneRequired("user", "user-file")
	OnpremCmd.MarkFlagFilename("user-file")
	OnpremCmd.MarkFlagsMutuallyExclusive("user", "user-file")
}
