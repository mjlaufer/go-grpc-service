package cmd

import (
	"fmt"
	"os"

	pb "github.com/mjlaufer/go-grpc-service/api/proto"
	c "github.com/mjlaufer/go-grpc-service/internal/cli/grpc-client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	client pb.UserServiceClient
	conn   *grpc.ClientConn
)

var rootCmd = &cobra.Command{
	Use: "UserService",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UserService CLI")
	},
}

// Execute runs the CLI
func Execute() {
	// Run the gRPC client.
	client, conn = c.RunClient()
	defer conn.Close()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
