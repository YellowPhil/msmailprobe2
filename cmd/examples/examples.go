package examples

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ExamplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "Example usage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(` ~~ Example usage ~~
msmailprobe identify -h mail.target.com
msmailprobe userenum onprem -t mail.target.com -U users.txt -o validusers.txt --threads 20
msmailprobe userenum onprem -t mail.target.com -U users.txt -o validusers.txt --threads 20 --rps 10
msmailprobe userenum o365 -t mail.target.com -E emailList.txt -o validemails.txt --threads 30
msmailprobe userenum o365 -t mail.target.com -E emailList.txt -o validemails.txt --threads 30 --rpm 100
msmailprobe userenum onprem -t mail.target.com -u admin
msmailprobe userenum o365 -e admin@target.com`)
	},
}
