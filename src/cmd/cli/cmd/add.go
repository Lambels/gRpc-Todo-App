package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	pb "github.com/Lambels/gRpc-Todo-App/service"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add will add a task with the given name",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
		defer cancel()

		if len(args) == 0 {
			return fmt.Errorf("there werent provided enough args to run this command")
		}

		out := &pb.Message{Message: strings.Join(args, " ")}
		_, err := client.Add(ctx, out)

		return err
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
