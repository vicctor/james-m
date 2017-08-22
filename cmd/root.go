package cmd

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/spf13/cobra"
)

var agentHost string
var agentPort uint16

var RootCmd = &cobra.Command{
	Use:   "james-m",
	Short: "Management CLI tool for James agent.",
	Long:  "James-M is a management CLI tool for James agent.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&agentHost, "host", "localhost", "host of James agent")
	RootCmd.PersistentFlags().Uint16Var(&agentPort, "port", 7007, "port of James agent")
}

func printAsJson(value interface{}) {
	jsonBytes, _ := json.MarshalIndent(value, "", "  ")
	fmt.Println(string(jsonBytes))
}
