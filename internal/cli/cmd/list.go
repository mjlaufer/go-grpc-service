package cmd

import (
	"context"
	"log"

	pb "github.com/mjlaufer/go-grpc-service/api/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List users.",
	Run: func(cmd *cobra.Command, args []string) {
		req := pb.ListRequest{}
		r, err := client.List(context.Background(), &req)
		if err != nil {
			log.Fatalf("Error: Could not list users: %v", err)
		}
		log.Printf("List result: <%+v>\n", r)
	},
}
