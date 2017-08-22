package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pdebicki/james-m/client"
	"io/ioutil"
	"log"
)

var includeAbstractClass bool
var includeNonAbstractClassDescendants bool

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Manage information points of James agent",
	Long:  "Manage information points of James agent",
}

var ipAddCmd = &cobra.Command{
	Use:   "add <method reference> <script file>",
	Short: "Add an information point",
	Long:  "Add an information point",
	Run:   ipAddCmdHandler,
	Args:  cobra.ExactArgs(2),
}

var ipRemoveCmd = &cobra.Command{
	Use:   "remove <method reference>",
	Short: "Remove an information point",
	Long:  "Remove an information point",
	Run:   ipRemoveCmdHandler,
	Args:  cobra.ExactArgs(1),
}

var ipShowCmd = &cobra.Command{
	Use:   "show [method reference]",
	Short: "Show information points",
	Long:  "Show information points",
	Run:   ipShowCmdHandler,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	RootCmd.AddCommand(ipCmd)
	ipCmd.AddCommand(ipAddCmd)
	ipCmd.AddCommand(ipRemoveCmd)
	ipCmd.AddCommand(ipShowCmd)

	ipAddCmd.Flags().BoolVar(&includeAbstractClass, "include-abstract-class",
		false, "include abstract class")
	ipAddCmd.Flags().BoolVar(&includeNonAbstractClassDescendants, "include-non-abstract-class-descendants",
		false, "include non-abstract class descendants")
}

func ipAddCmdHandler(cmd *cobra.Command, args []string) {
	methodReference := client.CreateMethodReference(args[0])
	scriptFileName := args[1]
	jamesClient := &client.JamesClient{
		Host: agentHost,
		Port: agentPort,
	}
	script, err := ioutil.ReadFile(scriptFileName)
	if err != nil {
		log.Fatal(err)
	}
	jamesClient.AddInformationPoint(methodReference, script, includeAbstractClass, includeNonAbstractClassDescendants)
}

func ipRemoveCmdHandler(cmd *cobra.Command, args []string) {
	methodReference := client.CreateMethodReference(args[0])
	jamesClient := &client.JamesClient{
		Host: agentHost,
		Port: agentPort,
	}
	jamesClient.RemoveInformationPoint(methodReference)
}

func ipShowCmdHandler(cmd *cobra.Command, args []string) {
	jamesClient := &client.JamesClient{
		Host: agentHost,
		Port: agentPort,
	}
	if len(args) == 0 {
		ips := jamesClient.GetInformationPoints()
		printAsJson(ips)
	} else {
		methodReference := client.CreateMethodReference(args[0])
		ip := jamesClient.GetInformationPoint(methodReference)
		printAsJson(ip)
	}
}
