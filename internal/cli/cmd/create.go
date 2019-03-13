package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	pb "github.com/mjlaufer/go-grpc-service/api/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user.",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter email: ")
		scanner.Scan()
		email := scanner.Text()
		req := pb.CreateRequest{
			User: &pb.User{
				Id:    uuid.New().String(),
				Email: email,
			},
		}
		r, err := client.Create(context.Background(), &req)
		if err != nil {
			log.Fatalf("Error: Could not create: %v", err)
		}
		log.Printf("Create result: <%+v>\n", r)
	},
}
