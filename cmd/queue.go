package cmd

import (
	"../client"
	"github.com/spf13/cobra"
)

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Show asynchronous queues of James agent",
	Long:  "Show asynchronous queues of James agent",
	Run:   queueCmdHandler,
}

var queueScriptEngineCmd = &cobra.Command{
	Use:   "script-engine",
	Short: "Show asynchronous queues of James agent's script engine",
	Long:  "Show asynchronous queues of James agent's script engine",
	Run:   queueScriptEngineCmdHandler,
	Args:  cobra.NoArgs,
}

var queueEventPublisherCmd = &cobra.Command{
	Use:   "event-publisher",
	Short: "Show asynchronous queues of James agent's event publisher",
	Long:  "Show asynchronous queues of James agent's event publisher",
	Run:   queueEventPublisherCmdHandler,
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(queueCmd)
	queueCmd.AddCommand(queueScriptEngineCmd)
	queueCmd.AddCommand(queueEventPublisherCmd)
}

func queueCmdHandler(cmd *cobra.Command, args []string) {
	jamesClient := &client.JamesClient{
		Host: agentHost,
		Port: agentPort,
	}
	queues := jamesClient.GetAllQueues()
	printAsJson(queues)
}

func queueScriptEngineCmdHandler(cmd *cobra.Command, args []string) {
	jamesClient := &client.JamesClient{
		Host: agentHost,
		Port: agentPort,
	}
	queue := jamesClient.GetScriptEngineQueue()
	printAsJson(queue)
}

func queueEventPublisherCmdHandler(cmd *cobra.Command, args []string) {
	jamesClient := &client.JamesClient{
		Host: agentHost,
		Port: agentPort,
	}
	queue := jamesClient.GetEventPublisherQueue()
	printAsJson(queue)
}
