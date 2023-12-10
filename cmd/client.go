package cmd

import (
	"fmt"
	"log"

	"github.com/Jagadishwaran-KA/sih-ipfs/pkg"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Starts the IPFS client",
	Long:  "Starts the IPFS client",
	Run: func(cmd *cobra.Command, args []string) {
		// Code to start the client
		// call the client logic here

		// get the server address and port
		server, err := cmd.Flags().GetString("server")
		if err != nil {
			log.Fatalf("Error while getting server address: %v", err)
		}

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			log.Fatalf("Error while getting server port: %v", err)
		}

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("Error while getting file: %v", err)
		}

		fmt.Println("Starting the client")

		fmt.Println("Server: ", server)
		fmt.Println("Port: ", port)
		fmt.Println("File: ", file)

		// call the client logic - I have called a dummy logic here
		pkg.CallDummyLogic()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringP("server", "s", "", "Server address")
	clientCmd.Flags().StringP("port", "p", "", "Server port")
	clientCmd.Flags().StringP("file", "f", "", "File to be sent")

	err := clientCmd.MarkFlagRequired("server")
	if err != nil {
		log.Fatalf("Error while marking server flag required: %v", err)
	}

	err = clientCmd.MarkFlagRequired("port")
	if err != nil {
		log.Fatalf("Error while marking port  flag required: %v", err)
	}

	err = clientCmd.MarkFlagRequired("file")
	if err != nil {
		log.Fatalf("Error while marking file flag required: %v", err)
	}
}
