package cmd

import (
	"fmt"
	"os"
	"os/signal"

	pb "github.com/Lambels/gRpc-Todo-App/service"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List will list all or some of your tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		var d, _ = cmd.Flags().GetBool("done")
		var m, _ = cmd.Flags().GetString("message")

		ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
		defer cancel()

		out := &pb.ListRequest{Message: m, Done: d}

		resp, err := client.List(ctx, out)
		if err != nil {
			return err
		}

		for _, t := range resp.GetTasks() {
			fmt.Fprintf(os.Stdout, "Id: %v, Message: %v, Done: %v \n", t.Id, t.Message, t.Done)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("done", "d", false, "See only the done tasks (default: false)")
	listCmd.Flags().String("message", "", "Sets a specific message to filter the results")
}
