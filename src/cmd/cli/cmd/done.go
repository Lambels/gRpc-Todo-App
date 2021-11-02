package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"

	pb "github.com/Lambels/gRpc-Todo-App/service"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Done will complete a command",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
		defer cancel()

		if len(args) != 1 {
			return fmt.Errorf("there were provided too many or not enough args to run this command")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		_, err = client.Done(ctx, &pb.TodoTask{Id: int64(id)})
		return err
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
