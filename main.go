package main

import (
	"github.com/yellowphil/msmailprobe2/cmd"
	"github.com/yellowphil/msmailprobe2/internal/logging"
)

func main() {
	logging.SetupLogger()
	cmd.Execute()
}
