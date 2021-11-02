package cmd

import (
	"fmt"
	"os"

	pb "github.com/Lambels/gRpc-Todo-App/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var client pb.TasksServiceClient

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo will keep track of your tasks",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cc, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stdout, "error connecting to grpc server: %s", err)
		os.Exit(1)
	}
	client = pb.NewTasksServiceClient(cc)
}
